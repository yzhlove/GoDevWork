package main

import (
	"WorkSpace/GoDevWork/HttpServerTo/com.yzh.grpc/chat15/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial(":1234", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := proto.NewPubSubServiceInterfaceClient(conn)
	stream, err := client.Sub(context.Background(), &proto.String{Var: "docker:"})
	if err != nil {
		panic(err)
	}

	for {
		result, err := stream.Recv()
		if err != nil {
			panic(err)
		}
		fmt.Println("docker result:", result)
	}

}
