package main

import (
	"fmt"
	"zinx-day6-base1/ziface"
	"zinx-day6-base1/znet"
)

func main() {
	server := znet.NewTcpServer()
	server.RegisterRouter(0, &pingRouter{id: 0})
	server.RegisterRouter(1, &helloRouter{id: 1})
	server.Run()
}

type pingRouter struct {
	znet.BaseRouter
	id uint32
}

type helloRouter struct {
	znet.BaseRouter
	id uint32
}

func (r pingRouter) Handle(req ziface.RequestInterface) {
	fmt.Printf("[ping] msg id:%d data:%s \n", req.GetMsgID(), string(req.GetData()))
	if err := req.GetConn().Send(r.id, []byte("ping...ping")); err != nil {
		fmt.Println("[ping] send err:", err)
	}
}

func (r helloRouter) Handle(req ziface.RequestInterface) {
	fmt.Printf("[hello] msg id:%d data:%s \n", req.GetMsgID(), string(req.GetData()))
	if err := req.GetConn().Send(r.id, []byte("what can i do for you ?")); err != nil {
		fmt.Println("[hello] send err:", err)
	}
}
