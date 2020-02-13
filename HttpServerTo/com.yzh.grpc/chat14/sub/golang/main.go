package main

import (
	"WorkSpace/GoDevWork/HttpServerTo/com.yzh.grpc/chat14/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
)

func main() {
	conn, err := grpc.Dial(":1234", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := proto.NewPubSubServiceInterfaceClient(conn)
	stream, err := client.Sub(context.Background(), &proto.String{Var: "golang:"})
	if err != nil {
		panic(err)
	}
	for {
		replay, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		fmt.Println("golang replay => ", replay)
	}
}
