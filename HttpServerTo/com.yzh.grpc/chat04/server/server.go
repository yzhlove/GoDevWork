package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct{}

func (h *HelloService) Hello(request string, replay *string) error {
	*replay = "hello: " + request
	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
