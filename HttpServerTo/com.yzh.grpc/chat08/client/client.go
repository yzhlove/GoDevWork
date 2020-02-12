package main

import (
	"fmt"
	"net"
	"net/rpc"
	"time"
)

func main() {

	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	clientChan := make(chan *rpc.Client)
	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				panic(err)
			}
			clientChan <- rpc.NewClient(conn)
		}
	}()
	doClientWork(clientChan)
	time.Sleep(3 * time.Second)
}

func doClientWork(clientChan <-chan *rpc.Client) {
	client := <-clientChan
	defer client.Close()
	var replay string
	if err := client.Call("HelloService.Hello", "client", &replay); err != nil {
		panic(err)
	}
	fmt.Println("replay => ", replay)
}
