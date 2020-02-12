package main

import (
	"fmt"
	"net"
	"net/rpc"
	"time"
)

//反向rpc

type HelloService struct{}

func (h *HelloService) Hello(request string, replay *string) error {
	*replay = "server to: " + request
	return nil
}

func main() {
	rpc.Register(new(HelloService))
	for {
		conn, _ := net.Dial("tcp", ":1234")
		if conn == nil {
			fmt.Println("waiting ...")
			time.Sleep(time.Second)
			continue
		}
		rpc.ServeConn(conn)
		conn.Close()
	}
}
