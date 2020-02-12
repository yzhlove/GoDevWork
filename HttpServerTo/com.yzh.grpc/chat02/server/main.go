package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{}

func (h *HelloService) Hello(request string, replay *string) error {
	*replay = "hello :" + request
	return nil
}

func main() {
	RegisterHelloService(&HelloService{})
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("Accept:", err)
		}
		go rpc.ServeConn(conn)
	}
}
