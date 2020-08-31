package main

import (
	"WorkSpace/GoDevWork/MicroServiceInAction/WorkerPoolDemo/Grpccp/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"time"
)

func main() {

	startServer()

	pool := New(defaultDialFunc, WithTimeout(time.Second*5),
		WithCheckReadyTimeout(time.Second),
		WithHeartbeatInterval(time.Second))
	conn, err := pool.DialConn(":1234")
	if err != nil {
		panic(err)
	}
	client := pb.NewHelloClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := client.SayHello(ctx, &pb.Request{Name: "hello world"})
	if err != nil {
		panic(err)
	}
	fmt.Println("result => ", result)
}

type hello struct{}

func (hello) SayHello(ctx context.Context, in *pb.Request) (*pb.Replay, error) {
	return &pb.Replay{Message: "Hello:" + in.Name}, nil
}

func startServer() {
	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &hello{})

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	go func() {
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()
}
