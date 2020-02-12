package main

import "fmt"

func main() {
	cli, err := DialHelloService("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	var replay string
	err = cli.Hello("hello 212312313", &replay)
	if err != nil {
		panic(err)
	}
	fmt.Println("result => ", replay)
}
