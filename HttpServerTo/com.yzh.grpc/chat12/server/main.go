package main

import (
	"WorkSpace/GoDevWork/HttpServerTo/com.yzh.grpc/chat12/proto"
	"context"
	"errors"
	"google.golang.org/grpc"
	"io"
	"net"
)

type HelloService struct{}

func (p *HelloService) Hello(ctx context.Context, request *proto.String) (*proto.String, error) {
	replay := &proto.String{Var: "server:" + request.Var}
	return replay, nil
}

func (p *HelloService) Channel(stream proto.HelloServiceInterface_ChannelServer) error {
	for {
		msg, err := stream.Recv()
		if err != nil {
			//io.EOF客户端断开连接
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}
		replay := &proto.String{Var: "channel msg:" + msg.Var}
		if err = stream.Send(replay); err != nil {
			return err
		}
	}
}

func main() {
	grpcServer := grpc.NewServer()
	proto.RegisterHelloServiceInterfaceServer(grpcServer, new(HelloService))
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	grpcServer.Serve(l)
}
