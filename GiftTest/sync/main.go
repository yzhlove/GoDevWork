package main

import (
	"WorkSpace/GoDevWork/GiftTest/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

///////////////////////////////////////
// Test Sync interface
///////////////////////////////////////

func main() {

	conn, err := grpc.Dial(":53000", grpc.WithInsecure())
	if err != nil {
		panic("grpc dial err:" + err.Error())
	}
	client := pb.NewGiftServiceClient(conn)

	stream, err := client.Sync(context.Background(), &pb.SyncReq{Zone: 1})
	if err != nil {
		panic("sync request err:" + err.Error())
	}
	status := make(chan struct{})
	fmt.Println("wait read data ...")
	go func() {
		for {
			if result, err := stream.Recv(); err != nil {
				fmt.Printf("recv err:%v \n", err)
				close(status)
				return
			} else {
				fmt.Println("result => ", result)
			}
		}
	}()
	<-status
}
