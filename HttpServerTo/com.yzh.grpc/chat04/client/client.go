package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	cli := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	var replay string
	if err = cli.Call("HelloService.Hello", "ccccdddd", &replay); err != nil {
		panic(err)
	}
	fmt.Println("server result => ", replay)
}
