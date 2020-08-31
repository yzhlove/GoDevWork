package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	ac "google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/reflection"
	"net"
	"sort"
	"testing"
	"time"
)

func customReadyCheck(ctx context.Context, conn *grpc.ClientConn) ac.State {
	name := "test"
	client := NewGreeterClient(conn)
	if _, err := client.SayHello(context.Background(), &HelloRequest{Name: name}); err != nil {
		return ac.Idle
	} else {
		return ac.Ready
	}
}

func TestNewWithOption(t *testing.T) {
	type want struct {
		timeout           time.Duration
		checkReadyTimout  time.Duration
		heartbeatInterval time.Duration
		readyCheckFunc    ReadyCheckFunc
	}

	td := 123 * time.Second

	tables := []struct {
		name string
		args []TrackerOption
		want want
	}{

		{
			name: "default",
			args: []TrackerOption{},
			want: want{
				timeout:           defaultTimeout,
				checkReadyTimout:  checkReadyTimeout,
				heartbeatInterval: heartbeatInterval,
				readyCheckFunc:    defaultReadyCheck,
			},
		},

		{
			name: "with timeout",
			args: []TrackerOption{WithTimeout(td)},
			want: want{
				timeout:           td,
				checkReadyTimout:  checkReadyTimeout,
				heartbeatInterval: heartbeatInterval,
				readyCheckFunc:    defaultReadyCheck,
			},
		},

		{
			name: "with heartbeat interval",
			args: []TrackerOption{WithHeartbeatInterval(td)},
			want: want{
				timeout:           defaultTimeout,
				checkReadyTimout:  checkReadyTimeout,
				heartbeatInterval: td,
				readyCheckFunc:    defaultReadyCheck,
			},
		},

		{
			name: "with custom ready check",
			args: []TrackerOption{WithCustomReadyCheckFunc(customReadyCheck)},
			want: want{
				timeout:           defaultTimeout,
				checkReadyTimout:  checkReadyTimeout,
				heartbeatInterval: heartbeatInterval,
				readyCheckFunc:    customReadyCheck,
			},
		},
	}

	for _, table := range tables {

		t.Run(table.name, func(t *testing.T) {

			tc := New(dialFunc, table.args...)

			if tc.timeout != table.want.timeout {
				t.Errorf("tacker timeout is %d,excepted %d.", tc.timeout, table.want.timeout)
			}

			if tc.checkReadyTimeout != table.want.checkReadyTimout {
				t.Errorf("tacker ready timeout is %d,excepted %d.", tc.checkReadyTimeout, table.want.checkReadyTimout)
			}

			if tc.heartbeatInterval != table.want.heartbeatInterval {
				t.Errorf("tacker heartbeat interval is %d,excepted %d.", tc.heartbeatInterval, table.want.heartbeatInterval)
			}

			if fmt.Sprintf("%v", tc.readyCheck) != fmt.Sprintf("%v", table.want.readyCheckFunc) {
				t.Errorf("tacker check func is %d,excepted %d.", tc.timeout, table.want.heartbeatInterval)
			}

		})
	}

}

type server struct{}

func (s *server) SayHello(ctx context.Context, in *HelloRequest) (*HelloReplay, error) {
	return &HelloReplay{Message: "Hello:" + in.Name}, nil
}

func startGrpcServer(t *testing.T) (*grpc.Server, string) {
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("failed to listen:%v ", err)
	}
	serviceServer := grpc.NewServer()
	RegisterGreeterServer(serviceServer, &server{})
	reflection.Register(serviceServer)

	go func() {
		if err := serviceServer.Serve(lis); err != nil {
			t.Fatalf("failed to server: %v ", err)
		}
	}()

	return serviceServer, lis.Addr().String()
}

func testHelloWorld(t *testing.T, conn *grpc.ClientConn) bool {
	name := "test"
	client := NewGreeterClient(conn)
	result, err := client.SayHello(context.Background(), &HelloRequest{Name: name})
	if err != nil {
		t.Fatal(err)
		return false
	}

	expected := fmt.Sprintf("Hello:%s", name)
	if result.Message != expected {
		t.Fatalf("server replay is %s excepted %s.", result.Message, expected)
		return false
	}
	return true
}

func TestServer(t *testing.T) {
	sev, addr := startGrpcServer(t)
	defer sev.GracefulStop()

	conn, err := dialFunc(addr)
	if err != nil {
		t.Fatal(err)
	}
	testHelloWorld(t, conn)
}

func testActives(t *testing.T, actives, excepted []string) {
	if len(actives) != len(excepted) {
		t.Fatalf("activies conn no consistent")
	}
	sort.Strings(actives)
	sort.Strings(excepted)
	for i, addr := range excepted {
		if actives[i] != addr {
			t.Fatalf("addr no equal")
		}
	}
}

func TestServerWithGrpccp(t *testing.T) {
	s, addr := startGrpcServer(t)
	defer s.GracefulStop()

	conn, err := GetConn(addr)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	testHelloWorld(t, conn)

	actives := Actives()
	testActives(t, actives, []string{addr})
}

func TestServerWithDial(t *testing.T) {
	s, addr := startGrpcServer(t)
	defer s.GracefulStop()

	//defaultPool := New(dialFunc)

	conn, err := Dial(addr)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	testHelloWorld(t, conn)
	actives := Actives()
	testActives(t, actives, []string{addr})
}

func TestServerWithCustomReadyCheck(t *testing.T) {
	s, addr := startGrpcServer(t)
	defer s.GracefulStop()

	pool := New(dialFunc, WithCustomReadyCheckFunc(customReadyCheck))
	conn, err := pool.GetConn(addr)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	testHelloWorld(t, conn)
	actives := pool.GetActives()
	testActives(t, actives, []string{addr})
}

func TestConnErrAddr(t *testing.T) {
	s, addr := startGrpcServer(t)
	defer s.GracefulStop()

	dialF := func(addr string) (*grpc.ClientConn, error) {
		return grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second))
	}

	pool := New(dialF, WithCheckReadyTimeout(time.Second))

	conn, err := pool.GetConn(addr)
	if err != nil {
		t.Fatal(err)
	}
	testHelloWorld(t, conn)

	errAddr := "xxxx"
	_, err = pool.GetConn(errAddr)
	if err == nil {
		t.Fatal("test conn failed")
	}
	t.Log("Err:", err)

	actives := pool.GetActives()
	testActives(t, actives, []string{addr})

}

func TestStopServer(t *testing.T) {
	s, addr := startGrpcServer(t)

	pool := New(dialFunc, WithTimeout(time.Second*2),
		WithCheckReadyTimeout(time.Second),
		WithHeartbeatInterval(time.Second))

	conn, err := pool.GetConn(addr)
	if err != nil {
		t.Fatal(err)
	}

	testHelloWorld(t, conn)
	actives := pool.GetActives()
	testActives(t, actives, []string{addr})
	time.Sleep(5 * time.Second)
	conn.Close()
	s.Stop()
	time.Sleep(10 * time.Second)

	as := pool.GetActives()
	testActives(t, as, []string{})

	_, err = pool.GetConn(addr)
	if err == nil {
		t.Fatal("conn err addr no raise error")
	}
	t.Log("GetConn err:", err)
}

func TestConnIdle(t *testing.T) {
	s, addr := startGrpcServer(t)
	defer s.GracefulStop()

	pool := New(dialFunc)

	_, err := pool.GetConn(addr)
	if err != nil {
		t.Fatal(err)
	}
	tc := pool.connections[addr]
	tc.cancel()
	tc.idle()

	_, err = pool.GetConn(addr)
	if err != errNoReady {
		t.Fatal(err)
	}

}

func TestConnExpired(t *testing.T) {
	s, addr := startGrpcServer(t)
	defer s.GracefulStop()

	first := true
	mockCheckFunc := func(ctx context.Context, conn *grpc.ClientConn) ac.State {
		if first == true {
			first = false
			return ac.Ready
		}
		return ac.Idle
	}

	pool := New(dialFunc, WithTimeout(2*time.Second),
		WithCheckReadyTimeout(time.Second),
		WithHeartbeatInterval(time.Second), WithCustomReadyCheckFunc(mockCheckFunc))

	conn, err := pool.GetConn(addr)
	if err != nil {
		t.Fatal(err)
	}
	testHelloWorld(t, conn)

	actives := pool.GetActives()
	testActives(t, actives, []string{addr})

	time.Sleep(time.Second * 5)

	as := pool.GetActives()
	testActives(t, as, []string{})
}
