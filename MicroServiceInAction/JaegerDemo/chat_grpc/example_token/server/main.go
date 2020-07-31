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

const addr = "127.0.0.1:1234"

const (
	path      = "/Users/yostar/workSpace/gowork/src/GoDevWork/MicroServiceInAction/JaegerDemo/chat_grpc/pem/"
	serverKey = path + "server.key"
	serverPem = path + "server.pem"
)

type hello struct{}

func (hello) SayHello(ctx context.Context, in *pb.HelloReq) (*pb.HelloResp, error) {

	fmt.Println("in --> ", in.Name)

	meta, ok := metadata.FromIncomingContext(ctx)
	//if !ok {
	//	return nil, errors.New("metadata invalid")
	//}

	fmt.Println("metadata ok => ", ok)
	for k, vars := range meta {
		fmt.Print("k => ", k)
		for _, v := range vars {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}

	var id, key string
	if vars, ok := meta["appid"]; ok {
		id = vars[0]
	}
	if vars, ok := meta["appkey"]; ok {
		key = vars[0]
	}
	if id != "1001" && key != "12138" {
		return nil, errors.New("meta data invalid")
	}
	return &pb.HelloResp{Message: fmt.Sprintf("succeed.[%s-%s]", id, key)}, nil
}

func main() {

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	creds, err := credentials.NewServerTLSFromFile(serverPem, serverKey)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterHelloServer(s, &hello{})

	if err := s.Serve(listen); err != nil {
		panic(err)
	}
}
