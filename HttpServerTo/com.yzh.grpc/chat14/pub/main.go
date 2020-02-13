package main

import (
	"WorkSpace/GoDevWork/HttpServerTo/com.yzh.grpc/chat14/proto"
	"context"
	"google.golang.org/grpc"
	"math/rand"
	"strconv"
	"time"
)

func main() {

	conn, err := grpc.Dial(":1234", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := proto.NewPubSubServiceInterfaceClient(conn)
	pubMessage(client)
}

func pubMessage(cli proto.PubSubServiceInterfaceClient) {
	tags := [...]string{"hi:", "golang:", "docker:"}
	for i := 0; i < 30; i++ {
		str := tags[rand.Intn(3)] + " message to " + strconv.Itoa(i+1)
		if _, err := cli.Pub(context.Background(), &proto.String{Var: str}); err != nil {
			panic(err)
		}
		time.Sleep(time.Second * 2)
	}
}
