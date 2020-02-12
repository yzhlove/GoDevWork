package main

import (
	"net"
	"net/rpc"
)

type HelloService struct {
	conn net.Conn
}

func ServerHelloService(conn net.Conn) {
	p := rpc.NewServer()
	p.Register(&HelloService{conn: conn})
	p.ServeConn(conn)
}

func (p *HelloService) Hello(request string, replay *string) error {
	*replay = "hello:" + request + " to " + p.conn.RemoteAddr().String()
	return nil
}

func main() {
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			panic("Accept : " + err.Error())
		}
		go ServerHelloService(conn)
	}
}
