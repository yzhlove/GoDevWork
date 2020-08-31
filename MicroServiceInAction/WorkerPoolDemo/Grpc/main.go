package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	ac "google.golang.org/grpc/connectivity"
	"sync"
	"time"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//	grpc 五种生命周期
//CONNECTING:
//	标识该 channel 正在尝试建立连接，并且正在等待在 名称解析、TCP 连接建立 或 TLS 握手 中涉及的其中一个步骤上取得进展。它可以被用作创建 channel 时的初始状态。
//READY:
//	标识该 channel 已经通过 TLS握手(或等效)和协议级(HTTP/2等)握手 成功建立了连接，并且所有后续的通信尝试都已成功(或者在没有任何已知故障的情况下等待)。
//TRANSIENT_FAILURE:
//	标识该 channel 出现了一些瞬时故障(例如，TCP 3次握手超时 或 套接字错误)。处于此状态的 channel 最终将切换到 CONNECTING 状态并尝试再次建立连接。
//	由于重试是通过指数退避(exponential backoff)完成的，因此，连接失败的 channel 在刚开始时在此状态下花费很少的时间，但是随着重复尝试并失败的次数增加，
//	channel 将在此状态下花费越来越多的时间。对于许多非致命故障(例如，由于服务器尚不可用而导致 TCP 连接尝试超时)，channel 可能在该状态下花费越来越多的大量时间。
//IDLE:
//	这个状态标识由于缺少新的(new)或未决(pending)的RPC，channel 甚至没有尝试创建连接。新的 RPC 可能会在这个状态被创建。任何在 channel 上启动 RPC
//	的尝试都会将通道从此状态推送到 CONNECTING 状态。
//	如果 channel 上已经在指定的 IDLE_TIMEOUT 时间内没有 RPC 活动，即在此期间没有新的(new)或挂起(pending)的(或活跃的) RPC，则 READY 或 CONNECTING
//	状态的 channel 将转换到 IDLE 状态。通常情况下，IDLE_TIMEOUT 的默认值是 300秒。
//	此外，已经接收到 GOAWAY 的 channel 在没有活跃(active)或挂起(pending)的 RPCs 时，也应当转换到 IDLE 状态，以避免尝试断开连接的服务器上的连接过载。
//SHUTDOWN:
//	标识该 channel 已经开始关闭。任何新的 RPCs 都应该立即失败。待处理(pending)的 RPCs 可能会继续运行，直到应用程序取消它们。channel 可能会因为应用程序显式请求关闭，
//	或者在尝试连接通信期间发生了不可恢复的错误而进入此状态。(截至2015年12月6日，没有已知的错误(连接或通信时)被归类为不可恢复的错误。)
//	一旦 channel 进入 SHUTDOWN 状态，就绝不会再离开。也就是说，SHUTDOWN 是状态机的结束。
//
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	defaultTimeout    = 100 * time.Second
	checkReadyTimeout = 5 * time.Second
	heartbeatInterval = 20 * time.Second
)

var (
	errNoReady = fmt.Errorf("no ready")
)

type DialFunc func(addr string) (*grpc.ClientConn, error)
type ReadyCheckFunc func(ctx context.Context, conn *grpc.ClientConn) ac.State

type ConnectionTracker struct {
	sync.Mutex
	dial              DialFunc
	readyCheck        ReadyCheckFunc
	connections       map[string]*trackedConn
	actives           map[string]*trackedConn
	timeout           time.Duration
	checkReadyTimeout time.Duration
	heartbeatInterval time.Duration

	ctx    context.Context
	cancel context.CancelFunc
}

type TrackerOption func(c *ConnectionTracker)

func WithTimeout(t time.Duration) TrackerOption {
	return func(c *ConnectionTracker) {
		c.timeout = t
	}
}

func WithCheckReadyTimeout(t time.Duration) TrackerOption {
	return func(c *ConnectionTracker) {
		c.checkReadyTimeout = t
	}
}

func WithHeartbeatInterval(t time.Duration) TrackerOption {
	return func(c *ConnectionTracker) {
		c.heartbeatInterval = t
	}
}

func WithCustomReadyCheckFunc(f ReadyCheckFunc) TrackerOption {
	return func(c *ConnectionTracker) {
		c.readyCheck = f
	}
}

func LoadPackage(trace *ConnectionTracker, opts ...TrackerOption) *ConnectionTracker {
	for _, f := range opts {
		f(trace)
	}
	return trace
}

func New(dial DialFunc, opts ...TrackerOption) *ConnectionTracker {
	ctx, cancel := context.WithCancel(context.Background())
	trace := &ConnectionTracker{
		dial:              dial,
		readyCheck:        defaultReadyCheck,
		connections:       make(map[string]*trackedConn),
		actives:           make(map[string]*trackedConn),
		timeout:           defaultTimeout,
		checkReadyTimeout: checkReadyTimeout,
		heartbeatInterval: heartbeatInterval,
		ctx:               ctx,
		cancel:            cancel,
	}

	return LoadPackage(trace, opts...)
}

func (trace *ConnectionTracker) GetConn(addr string) (*grpc.ClientConn, error) {
	return trace.getconn(addr, false)
}

func (trace *ConnectionTracker) Dial(addr string) (*grpc.ClientConn, error) {
	return trace.getconn(addr, true)
}

func (trace *ConnectionTracker) getconn(addr string, force bool) (*grpc.ClientConn, error) {
	trace.Lock()
	tc, ok := trace.connections[addr]
	if !ok {
		tc = &trackedConn{addr: addr, tracker: trace}
		trace.connections[addr] = tc
	}
	trace.Unlock()

	if err := tc.tryconn(trace.ctx, force); err != nil {
		return nil, err
	}
	return tc.conn, nil
}

func (trace *ConnectionTracker) connReady(tc *trackedConn) {
	trace.Lock()
	defer trace.Unlock()
	trace.actives[tc.addr] = tc
}

func (trace *ConnectionTracker) connUnReady(addr string) {
	trace.Lock()
	defer trace.Unlock()
	delete(trace.actives, addr)
}

func (trace *ConnectionTracker) GetActives() []string {
	trace.Lock()
	defer trace.Unlock()
	address := make([]string, 0, len(trace.actives))
	for addr := range trace.actives {
		address = append(address, addr)
	}
	return address
}

type trackedConn struct {
	sync.Mutex
	addr    string
	conn    *grpc.ClientConn
	tracker *ConnectionTracker
	state   ac.State
	expired time.Time
	retry   int
	cancel  context.CancelFunc
}

func (tc *trackedConn) tryconn(ctx context.Context, force bool) error {
	tc.Lock()
	defer tc.Unlock()

	if !force && tc.conn != nil {
		if tc.state == ac.Ready {
			return nil
		}
		if tc.state == ac.Idle {
			return errNoReady
		}
	}

	if tc.conn != nil {
		tc.conn.Close()
	}
	if conn, err := tc.tracker.dial(tc.addr); err != nil {
		return err
	} else {
		tc.conn = conn
	}

	readyCtx, cancel := context.WithTimeout(ctx, tc.tracker.checkReadyTimeout)
	defer cancel()

	if check := tc.tracker.readyCheck(readyCtx, tc.conn); check != ac.Ready {
		return errNoReady
	}
	_ctx, _cancel := context.WithCancel(ctx)
	tc.cancel = _cancel
	go tc.heartbeat(_ctx)

	tc.ready()
	return nil
}

func (tc *trackedConn) GetState() ac.State {
	tc.Lock()
	defer tc.Unlock()
	return tc.state
}

func (tc *trackedConn) healthCheck(ctx context.Context) {
	tc.Lock()
	defer tc.Unlock()
	ctx, cancel := context.WithTimeout(ctx, tc.tracker.checkReadyTimeout)
	defer cancel()

	switch tc.tracker.readyCheck(ctx, tc.conn) {
	case ac.Ready:
		tc.ready()
	case ac.Shutdown:
		tc.shutdown()
	case ac.Idle:
		if tc.expire() {
			tc.shutdown()
		} else {
			tc.idle()
		}
	}
}

func defaultReadyCheck(ctx context.Context, conn *grpc.ClientConn) ac.State {
	for {
		if state := conn.GetState(); state == ac.Ready || state == ac.Shutdown {
			return state
		} else {
			if !conn.WaitForStateChange(ctx, state) {
				return ac.Idle
			}
		}
	}
}

func (tc *trackedConn) ready() {
	tc.state = ac.Ready
	tc.expired = time.Now().Add(tc.tracker.timeout)
	tc.retry = 0
	tc.tracker.connReady(tc)
}

func (tc *trackedConn) idle() {
	tc.state = ac.Idle
	tc.retry++
	tc.tracker.connUnReady(tc.addr)
}

func (tc *trackedConn) shutdown() {
	tc.state = ac.Shutdown
	tc.conn.Close()
	tc.cancel()
	tc.tracker.connUnReady(tc.addr)
}

func (tc *trackedConn) expire() bool {
	return tc.expired.Before(time.Now())
}

func (tc *trackedConn) heartbeat(ctx context.Context) {
	tk := time.NewTicker(tc.tracker.heartbeatInterval)
	for tc.GetState() != ac.Shutdown {
		select {
		case <-ctx.Done():
			tc.shutdown()
		case <-tk.C:
			tc.healthCheck(ctx)
		}
	}
}

var (
	defaultPool *ConnectionTracker
	once        sync.Once
)

var dialFunc = func(addr string) (*grpc.ClientConn, error) {
	return grpc.Dial(addr, grpc.WithInsecure())
}

func pool() *ConnectionTracker {
	once.Do(func() {
		defaultPool = New(dialFunc)
	})
	return defaultPool
}

func GetConn(addr string) (*grpc.ClientConn, error) {
	return pool().GetConn(addr)
}

func Dial(addr string) (*grpc.ClientConn, error) {
	return pool().Dial(addr)
}

func Actives() []string {
	return pool().GetActives()
}
