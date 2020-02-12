package main

import (
	"WorkSpace/GoDevWork/HttpServerTo/com.yzh.grpc/chat03/api"
	"fmt"
)

func main() {

	cli, err := api.DialHelloService("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	var replay string
	if err = cli.Hello("i am client", &replay); err != nil {
		panic(err)
	}
	fmt.Println("server => ", replay)
}
