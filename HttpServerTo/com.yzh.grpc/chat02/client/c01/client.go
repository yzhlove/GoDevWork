package main

import (
	"WorkSpace/GoDevWork/HttpServerTo/com.yzh.grpc/chat02/conf"
	"fmt"
	"net/rpc"
)

func main() {

	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		panic(err)
	}
	var replay string
	err = client.Call(conf.HelloServiceName+".Hello", "hi", &replay)
	if err != nil {
		panic(err)
	}
	fmt.Println("result => ", replay)
}
