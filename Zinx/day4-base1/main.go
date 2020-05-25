package main

import (
	"fmt"
	"zinx-day4-base1/ziface"
	"zinx-day4-base1/znet"
)

func main() {

	server := znet.NewServer()
	server.RegisterRouter(0, &PingRouter{})
	server.RegisterRouter(1, &HelloRouter{})
	server.Run()
}

type PingRouter struct {
	znet.BaseRouter
}

func (PingRouter) Handle(req ziface.RequestInterface) {
	fmt.Println("[ROUTER] request to ping handle .")
	fmt.Printf("msg id:%d data:%s \n", req.GetMessageID(), string(req.GetData()))

	if err := req.GetConn().Send(0, []byte("ping-pong")); err != nil {
		fmt.Println("ping:", err)
	}
}

type HelloRouter struct {
	znet.BaseRouter
}

func (HelloRouter) Handle(req ziface.RequestInterface) {
	fmt.Println("[ROUTER] request to hello handle .")
	fmt.Printf("msg id:%d data:%s \n", req.GetMessageID(), string(req.GetData()))

	if err := req.GetConn().Send(1, []byte("hello world")); err != nil {
		fmt.Println("hello:", err)
	}

}
