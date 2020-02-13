package main

import (
	"WorkSpace/GoDevWork/HttpServerTo/com.yzh.grpc/chat11/proto"
	"context"
	"google.golang.org/grpc"
	"net"
)

type HelloService struct{}

func (p *HelloService) Hello(ctx context.Context, req *proto.String) (resp *proto.String, err error) {
	replay := &proto.String{
		Value: "Hello:" + req.Value,
	}
	return replay, nil
}

func main() {
	grpcS := grpc.NewServer()
	proto.RegisterHelloServiceInterfaceServer(grpcS, new(HelloService))
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	grpcS.Serve(l)
}
