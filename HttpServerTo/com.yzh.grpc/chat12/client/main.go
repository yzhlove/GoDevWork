package main

import (
	"WorkSpace/GoDevWork/HttpServerTo/com.yzh.grpc/chat12/proto"
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
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

	client := proto.NewHelloServiceInterfaceClient(conn)
	replay, err := client.Hello(context.Background(), &proto.String{Var: "love you"})
	if err != nil {
		log.Println("Hello Err: " + err.Error())
	}
	fmt.Println("Replay => ", replay)

	stream, err := client.Channel(context.Background())
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			if err := stream.Send(&proto.String{
				Var: "client:" + strconv.Itoa(rand.Intn(1000)+100),
			}); err != nil {
				panic("client send err :" + err.Error())
			}
			time.Sleep(time.Second * 3)
		}
	}()

	for {
		replay, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic("client recv err:" + err.Error())
		}
		fmt.Println("replay =>", replay)
	}
}
