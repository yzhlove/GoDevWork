package main

import (
	"WorkSpace/GoDevWork/MicroServiceInAction/JaegerDemo/chat_grpc/pb"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

const (
	grpcAddr  = "127.0.0.1:1234"
	addr      = "/Users/yostar/workSpace/gowork/src/GoDevWork/MicroServiceInAction/JaegerDemo/chat_grpc/pem/"
	serverPem = addr + "server.pem"
)

func main() {

	creds, err := credentials.NewClientTLSFromFile(serverPem, "abc.yzhdomain.com")
	if err != nil {
		panic(err)
	}
	conn, err := grpc.Dial(grpcAddr, grpc.WithTransportCredentials(creds))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewHelloClient(conn)
	if resp, err := c.SayHello(context.Background(),
		&pb.HelloReq{Name: "yzh"}); err != nil {
		panic(err)
	} else {
		log.Println("resp => ", resp.Message)
	}
}
