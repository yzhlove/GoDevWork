package main

import (
	"zinx/ziface"
	"zinx/zlog"
	"zinx/znet"
)

func main() {
	server := znet.NewTcpServer()
	server.Register(0, &panicRouter{})
	server.Run()
}

type panicRouter struct {
	znet.AbstractRouter
}

func (panicRouter) Handle(req ziface.ReqImp) {
	zlog.Info("[ping] msg id:", req.GetMsgID(), " client say:", string(req.GetMsgData()))
	panic("test panic err")
}
