package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

func main() {
	server := znet.NewTcpServer()
	server.Register(0, &pingRouter{})
	server.Run()
}

type pingRouter struct {
	znet.AbstractRouter
}

func (pingRouter) Handle(req ziface.ReqImp) {
	fmt.Println("[ping] msg id:", req.GetMsgID(), " client say:", string(req.GetMsgData()))
	panic("test panic err")
}
