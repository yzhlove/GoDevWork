package main

import (
	"net"
	"net/rpc"
)

type HelloService struct{}

func (p *HelloService) Hello(request string, replay *string) error {
	*replay = "hello:" + request
	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic("listen tcp err: " + err.Error())
	}
	conn, err := l.Accept()
	if err != nil {
		panic("accept err: " + err.Error())
	}
	rpc.ServeConn(conn)
}
