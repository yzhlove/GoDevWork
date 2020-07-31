package main

import (
	"WorkSpace/GoDevWork/MicroServiceInAction/JaegerDemo/chat_grpc/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

const (
	grpcAddr  = "127.0.0.1:1234"
	addr      = "/Users/yostar/workSpace/gowork/src/GoDevWork/MicroServiceInAction/JaegerDemo/chat_grpc/pem/"
	serverKey = addr + "server.key"
	serverPem = addr + "server.pem"
)

func main() {

	listen, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		panic(err)
	}

	//TLS认证
	creds, err := credentials.NewServerTLSFromFile(serverPem, serverKey)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer(grpc.Creds(creds))

	pb.RegisterHelloServer(s, &hello{})
	log.Println("listen to <<<0.0.0.0:1234>>>")
	if err := s.Serve(listen); err != nil {
		panic(err)
	}
}

type hello struct{}

func (h hello) SayHello(ctx context.Context, in *pb.HelloReq) (*pb.HelloResp, error) {
	out := new(pb.HelloResp)
	out.Message = fmt.Sprintf("Hello %s.", in.Name)
	return out, nil
}
