package main

import (
	"WorkSpace/GoDevWork/MicroServiceInAction/JaegerDemo/chat_grpc/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"time"
)

const (
	addr = "127.0.0.1:1234"
	path = "/Users/yostar/workSpace/gowork/src/GoDevWork/MicroServiceInAction/JaegerDemo/chat_grpc/pem/"
	pem  = path + "server.pem"
	tls  = true
)

type custome struct{}

func (custome) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"token": "token token !!!"}, nil
}

func (custome) RequireTransportSecurity() bool {
	return tls
}

func main() {
	var err error
	var opts []grpc.DialOption

	if tls {
		if creds, err := credentials.NewClientTLSFromFile(pem, "abc.yzhdomain.com"); err != nil {
			panic(err)
		} else {
			opts = append(opts, grpc.WithTransportCredentials(creds))
		}
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	opts = append(opts, grpc.WithPerRPCCredentials(new(custome)), grpc.WithUnaryInterceptor(interceptor))

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewHelloClient(conn)
	if resp, err := client.SayHello(context.Background(), &pb.HelloReq{Name: "Rpc"}); err != nil {
		panic(err)
	} else {
		fmt.Println("resp => ", resp.Message)
	}

}

func interceptor(ctx context.Context, method string, req, replay interface{}, cc *grpc.ClientConn, invoke grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	start := time.Now()
	err := invoke(ctx, method, req, replay, cc, opts...)
	fmt.Printf("method=%s req=%v replay=%v duration=%s error=%v \n", method, req, replay, time.Since(start), err)
	return err
}
