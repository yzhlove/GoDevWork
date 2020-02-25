package main

import (
	"WorkSpace/GoDevWork/HttpServerTo/com.yzh.grpc/chat16/proto"
	"context"
	"google.golang.org/grpc"
)

func main() {

	c, err := grpc.Dial(":1234", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer c.Close()
	cli := proto.NewPubSubServiceInterfaceClient(c)
	cli.Pub(context.TODO(), &proto.Manager_Msg{
		Zone: 0,
		Var:  "what are you doing ...",
	})

}
