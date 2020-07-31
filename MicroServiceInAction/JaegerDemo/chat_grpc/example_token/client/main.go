package main

import (
	"WorkSpace/GoDevWork/MicroServiceInAction/JaegerDemo/chat_grpc/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	addr      = "127.0.0.1:1234"
	openTLS   = true
	path      = "/Users/yostar/workSpace/gowork/src/GoDevWork/MicroServiceInAction/JaegerDemo/chat_grpc/pem/"
	serverPem = path + "server.pem"
)

type custom struct{}

func (custom) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{"appid": "1001", "appkey": "12138"}, nil
}

func (custom) RequireTransportSecurity() bool {
	return openTLS
}

func main() {

	var err error
	var opts []grpc.DialOption

	if openTLS {
		creds, err := credentials.NewClientTLSFromFile(serverPem, "*.yzhdomain.com")
		if err != nil {
			panic(err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	opts = append(opts, grpc.WithPerRPCCredentials(new(custom)))
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewHelloClient(conn)

	//ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("yzhlove", "12345"))

	if resp, err := c.SayHello(context.Background(), &pb.HelloReq{
		Name: "Rain",
	}); err != nil {
		panic(err)
	} else {
		fmt.Print("result => ", resp.Message)
	}
}
