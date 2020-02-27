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
	_, err = client.Generate(context.Background(), &proto.Manager_GenReq{
		FixCode:      "abcdefg",
		Num:          10,
		StartTime:    0,
		EndTime:      0,
		TimesPerCode: 0,
		TimesPerUser: 0,
		ZoneIds:      nil,
		Items:        nil,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("generate ok !")

}
