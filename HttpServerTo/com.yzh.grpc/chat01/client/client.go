package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		panic("dail err:" + err.Error())
	}

	var replay string
	err = client.Call("HelloService.Hello", "hi", &replay)
	if err != nil {
		panic("rpc call err: " + err.Error())
	}
	fmt.Println("replay => " + replay)
}
