package main

import (
	"fmt"
	"net/rpc"
)

const HelloServiceName = "path/to/pkg.HelloService"

func main() {
	cli, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	doClientWork(cli)
}

func doClientWork(cli *rpc.Client) {
	helloCall := cli.Go(HelloServiceName+".Hello", "hihihi", new(string), nil)
	h := <-helloCall.Done
	if err := h.Error; err != nil {
		panic(err)
	}
	args := h.Args.(string)
	replay := h.Reply.(*string)
	fmt.Println("args ==> ", args, " replay ==>", *replay)
}
