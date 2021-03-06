package main

import (
	"WorkSpace/GoDevWork/GiftServer/config"
	"WorkSpace/GoDevWork/GiftServer/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial(config.Listen, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := proto.NewGiftServiceClient(conn)
	stream, err := client.Sync(context.Background(),
		&proto.SyncReq{Zone: 1})
	if err != nil {
		panic(err)
	}
	defer stream.CloseSend()
	for {
		msg, err := stream.Recv()
		if err != nil {
			panic(err)
		}
		fmt.Println("RecvMessage:", msg)
	}
}
