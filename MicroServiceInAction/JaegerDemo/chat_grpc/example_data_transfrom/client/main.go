package main

import (
	"WorkSpace/GoDevWork/MicroServiceInAction/JaegerDemo/chat_grpc/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {

	conn, err := grpc.Dial(":1234", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewHelloClient(conn)

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("love", "wyq"))
	ctx = metadata.NewIncomingContext(ctx, metadata.Pairs("miss", "wyq"))

	if resp, err := client.SayHello(ctx, &pb.HelloReq{Name: "yzh"}); err != nil {
		panic(err)
	} else {
		fmt.Println("resp -> ", resp.Message)
	}
}
