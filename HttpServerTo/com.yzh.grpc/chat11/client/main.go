package main

import (
	"WorkSpace/GoDevWork/HttpServerTo/com.yzh.grpc/chat11/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial(":1234", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	cli := proto.NewHelloServiceInterfaceClient(conn)
	replay, err := cli.Hello(context.Background(), &proto.String{
		Value: "i am to client",
	})
	if err != nil {
		log.Println("Err => ", err)
	} else {
		fmt.Println("replay => ", replay)
	}
}
