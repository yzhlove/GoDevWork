package main

import (
	"WorkSpace/GoDevWork/HttpServerTo/com.yzh.grpc/chat16/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main() {

	c, err := grpc.Dial(":1234", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer c.Close()

	cli := proto.NewPubSubServiceInterfaceClient(c)
	stream, err := cli.Sub(context.TODO(), &proto.Manager_Zone{Var: 1})
	if err != nil {
		panic(err)
	}
	count := 0
	for {
		//if count >= 3 {
		//	break
		//}
		result, err := stream.Recv()
		if err != nil {
			panic(err)
		}
		count++
		fmt.Println("result => ", result, " count => ", count)

	}
	//stream.CloseSend()
	//fmt.Println("close stream ...")
	//time.Sleep(time.Second * 5)
	//fmt.Println("exit ...")
}
