package main

import (
	"net"
	"net/rpc"
)

const HelloServiceName = "path/to/pkg.HelloService"

type HelloServiceInterface = interface {
	Hello(request string, replay *string) error
}

func RegisterHelloService(hsi HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, hsi)
}

type HelloService struct{}

func (h *HelloService) Hello(request string, replay *string) error {
	*replay = "server:" + request
	return nil
}

func main() {
	RegisterHelloService(new(HelloService))
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			panic(err)
		}
		go rpc.ServeConn(c)
	}
}
