package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	cli, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	var replay string
	if err = cli.Call("HelloService.Hello", "abab", &replay); err != nil {
		panic(err)
	}
	fmt.Println("replay => ", replay)
}
