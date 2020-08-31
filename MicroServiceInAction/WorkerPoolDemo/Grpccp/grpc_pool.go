package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	status "google.golang.org/grpc/connectivity"
	"sync"
	"time"
)

var errNoReady = fmt.Errorf("no ready")
var defaultCheckReadyFunc = func(ctx context.Context, conn *grpc.ClientConn) status.State {
	for {
		s := conn.GetState()
		if s == status.Ready || s == status.Shutdown {
			return s
		}
		if !conn.WaitForStateChange(ctx, s) {
			return status.Idle
		}
	}
}
var defaultDialFunc = func(addr string) (*grpc.ClientConn, error) {
	return grpc.Dial(addr, grpc.WithInsecure())
}

const (
	defaultTimeout    = 10 * time.Second
	checkReadyTimeout = 5 * time.Second
	heartbeatInterval = 20 * time.Second
)

type DialFunc func(addr string) (*grpc.ClientConn, error)
type CheckReadyFunc func(ctx context.Context, conn *grpc.ClientConn) status.State

type ConnTracker struct {
	sync.Mutex
	dialFunc          DialFunc
	checkReadyFunc    CheckReadyFunc
	conns             map[string]*C
	actives           map[string]*C
	timeout           time.Duration
	checkReadyTimeout time.Duration
	heartbeatInterval time.Duration

	ctx    context.Context
	cancel context.CancelFunc
}

type C struct {
	sync.Mutex
	addr   string
	conn   *grpc.ClientConn
	parent *ConnTracker
	state  status.State
	expire time.Time
	retry  int
	cancel context.CancelFunc
}

type TrackerOption func(c *ConnTracker)

func WithTimeout(t time.Duration) TrackerOption {
	return func(c *ConnTracker) {
		c.timeout = t
	}
}

func WithCheckReadyTimeout(t time.Duration) TrackerOption {
	return func(c *ConnTracker) {
		c.checkReadyTimeout = t
	}
}

func WithHeartbeatInterval(t time.Duration) TrackerOption {
	return func(c *ConnTracker) {
		c.heartbeatInterval = t
	}
}

func WithCheckReadyFunc(f CheckReadyFunc) TrackerOption {
	return func(c *ConnTracker) {
		c.checkReadyFunc = f
	}
}

func (t *ConnTracker) loadOption(opts ...TrackerOption) *ConnTracker {
	for _, f := range opts {
		f(t)
	}
	return t
}

func New(dialFunc DialFunc, opts ...TrackerOption) *ConnTracker {
	ctx, cancel := context.WithCancel(context.Background())
	trace := &ConnTracker{
		dialFunc:          dialFunc,
		checkReadyFunc:    defaultCheckReadyFunc,
		conns:             make(map[string]*C),
		actives:           make(map[string]*C),
		timeout:           defaultTimeout,
		checkReadyTimeout: checkReadyTimeout,
		heartbeatInterval: heartbeatInterval,
		ctx:               ctx,
		cancel:            cancel,
	}
	return trace.loadOption(opts...)
}

func (t *ConnTracker) Actives() []string {
	t.Lock()
	defer t.Unlock()
	s := make([]string, 0, len(t.actives))
	for addr := range t.actives {
		s = append(s, addr)
	}
	return s
}

func (t *ConnTracker) GetConn(addr string) (*grpc.ClientConn, error) {
	return t.conn(addr, false)
}

func (t *ConnTracker) DialConn(addr string) (*grpc.ClientConn, error) {
	return t.conn(addr, true)
}

func (t *ConnTracker) conn(addr string, force bool) (*grpc.ClientConn, error) {
	t.Lock()
	c, ok := t.conns[addr]
	if !ok {
		c = &C{addr: addr, parent: t}
		t.conns[addr] = c
	}
	t.Unlock()

	if err := c.tryConn(t.ctx, force); err != nil {
		return nil, err
	}
	return c.conn, nil
}

func (c *C) tryConn(ctx context.Context, force bool) error {
	c.Lock()
	defer c.Unlock()

	//是否强制更新连接信息
	if !force && c.conn != nil {
		if c.state == status.Ready {
			return nil
		}
		if c.state == status.Idle {
			return errNoReady
		}
	}

	if c.conn != nil {
		c.conn.Close()
	}

	if _conn, err := c.parent.DialConn(c.addr); err != nil {
		return err
	} else {
		c.conn = _conn
	}

	readyCtx, cancel := context.WithTimeout(ctx, c.parent.checkReadyTimeout)
	defer cancel()

	if s := c.parent.checkReadyFunc(readyCtx, c.conn); s != status.Ready {
		return errNoReady
	}

	newCtx, newCancel := context.WithCancel(ctx)
	c.cancel = newCancel

	go c.heartbeat(newCtx)

	c.stateToReady()
	return nil
}

func (c *C) heartbeat(ctx context.Context) {
	t := time.NewTicker(c.parent.heartbeatInterval)
	for c.State() != status.Shutdown {
		select {
		case <-ctx.Done():
			c.stateToShutdown()
		case <-t.C:
			c.check(ctx)
		}
	}
}

func (c *C) check(ctx context.Context) {
	c.Lock()
	defer c.Unlock()

	ctx, cancel := context.WithTimeout(ctx, c.parent.checkReadyTimeout)
	defer cancel()

	switch c.parent.checkReadyFunc(ctx, c.conn) {
	case status.Ready:
		c.stateToReady()
	case status.Shutdown:
		c.stateToShutdown()
	case status.Idle:
		if c.expire.Before(time.Now()) {
			c.stateToShutdown()
		} else {
			c.stateToIdle()
		}
	}
}

func (c *C) State() status.State {
	c.Lock()
	defer c.Unlock()
	return c.state
}

func (t *ConnTracker) readyConn(c *C) {
	t.Lock()
	defer t.Unlock()
	t.actives[c.addr] = c
}

func (t *ConnTracker) unreadyConn(c *C) {
	t.Lock()
	defer t.Unlock()
	delete(t.actives, c.addr)
}

func (c *C) stateToReady() {
	c.state = status.Ready
	c.expire = time.Now().Add(c.parent.timeout)
	c.retry = 0
	c.parent.readyConn(c)
}

func (c *C) stateToIdle() {
	c.state = status.Idle
	c.retry++
	c.parent.unreadyConn(c)
}

func (c *C) stateToShutdown() {
	c.state = status.Shutdown
	c.conn.Close()
	c.cancel()
	c.parent.unreadyConn(c)
}
