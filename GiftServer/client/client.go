package main

import (
	"WorkSpace/GoDevWork/GiftServer/config"
	"WorkSpace/GoDevWork/GiftServer/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
)

func main() {

	conn, err := grpc.Dial(config.Listen, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := proto.NewGiftServiceClient(conn)
	stream, err := client.Sync(context.Background(), &proto.SyncReq{Zone: 1})
	go func() {
		for {
			msg, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
			fmt.Println("SyncResp => ", msg.String())
		}
	}()

	resp, err := client.CodeVerify(context.Background(), &proto.VerifyReq{Zone: 1, Code: "12345", UserId: 12345})
	if err != nil {
		panic(err)
	}
	fmt.Println("CodeVerifyResp => ", resp.String())
}
