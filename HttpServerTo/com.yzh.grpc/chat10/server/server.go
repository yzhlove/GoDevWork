package main

import (
	"errors"
	"log"
	"net"
	"net/rpc"
)

type HelloService struct {
	conn    net.Conn
	isLogin bool
}

func ServerHelloService(conn net.Conn) {
	p := rpc.NewServer()
	p.Register(&HelloService{conn: conn})
	p.ServeConn(conn)
}

func (p *HelloService) Login(request string, replay *string) error {
	if request != "user:password" {
		log.Println("login failed")
		return errors.New("login failed")
	}
	log.Println("login succeed")
	p.isLogin = true
	return nil
}

func (p *HelloService) Hello(request string, replay *string) error {
	if !p.isLogin {
		return errors.New("please login")
	}
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
			panic("accept:" + err.Error())
		}
		go ServerHelloService(conn)
	}
}
