///////////////////////////////////////////////////////////////////////
// grpc 拦截器
///////////////////////////////////////////////////////////////////////

package main

import (
	"WorkSpace/GoDevWork/MicroServiceInAction/JaegerDemo/chat_grpc/pb"
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"net"
)

const (
	addr = "127.0.0.1:1234"
	path = "/Users/yostar/workSpace/gowork/src/GoDevWork/MicroServiceInAction/JaegerDemo/chat_grpc/pem/"
	key  = path + "server.key"
	pem  = path + "server.pem"
)

type base struct{}

func (base) SayHello(context.Context, *pb.HelloReq) (*pb.HelloResp, error) {
	return &pb.HelloResp{Message: "succeed."}, nil
}

func main() {

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	var opts []grpc.ServerOption
	creds, err := credentials.NewServerTLSFromFile(pem, key)
	if err != nil {
		panic(err)
	}

	opts = append(opts, grpc.Creds(creds), grpc.UnaryInterceptor(authInterceptor))
	s := grpc.NewServer(opts...)

	pb.RegisterHelloServer(s, &base{})
	if err := s.Serve(listen); err != nil {
		panic(err)
	}
}

func checkAuth(ctx context.Context) error {
	meta, ok := metadata.FromIncomingContext(ctx)
	if ok {
		if vars, ok := meta["token"]; ok {
			fmt.Println("token => ", vars[0])
			return nil
		}
	}
	return errors.New("auth err : not found token ")
}

func authInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, h grpc.UnaryHandler) (interface{}, error) {
	if err := checkAuth(ctx); err != nil {
		return nil, err
	}
	fmt.Println("info => ", info.FullMethod, info.Server)
	return h(ctx, req)
}
