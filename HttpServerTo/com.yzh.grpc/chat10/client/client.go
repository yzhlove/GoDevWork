package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	cli, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	var replay string
	if err = cli.Call("HelloService.LoginEndpoint", "abc", &replay); err != nil {
		log.Println(err)
	} else {
		log.Println("login ok")
	}

	if err = cli.Call("HelloService.LoginEndpoint", "user:password", &replay); err != nil {
		log.Println(err)
	} else {
		log.Println("login ok")
	}

	if err = cli.Call("HelloService.Hello", "username=>password", &replay); err != nil {
		log.Println(err)
	}

	fmt.Println("replay => ", replay)

}
