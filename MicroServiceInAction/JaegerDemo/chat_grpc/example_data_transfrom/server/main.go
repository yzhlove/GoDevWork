package main

import (
	"WorkSpace/GoDevWork/MicroServiceInAction/JaegerDemo/chat_grpc/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net"
)

type base struct{}

// SayHello hello test program...
func (base) SayHello(ctx context.Context, in *pb.HelloReq) (*pb.HelloResp, error) {

	//获取client metadata  数据
	if meta, ok := metadata.FromOutgoingContext(ctx); ok {
		showMeta("FromOutgoingContext", meta)
	}
	//if meta, ok := metadata.FromIncomingContext(ctx); ok {
	//	showMeta("FromIncomingContext", meta)
	//}
	return &pb.HelloResp{Message: "succeed."}, nil
}

func showMeta(typ string, data metadata.MD) {
	fmt.Print("typ:", typ)
	for k, vars := range data {
		fmt.Print(k, "->")
		for _, v := range vars {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}

func main() {

	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &base{})
	if err := s.Serve(listen); err != nil {
		panic(err)
	}
}
