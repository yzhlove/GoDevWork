package main

import (
	"WorkSpace/GoDevWork/HttpServerTo/com.yzh.grpc/chat03/api"
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{}

func (h *HelloService) Hello(request string, replay *string) error {
	*replay = "server: " + request
	return nil
}

func main() {
	api.RegisterHelloService(new(HelloService))
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("accept:", err)
		}
		go rpc.ServeConn(conn)
	}
}
