package main

import (
	"fmt"
	"zinx-day3-base2/ziface"
	"zinx-day3-base2/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

func (p *PingRouter) Handle(req ziface.RequestInterface) {
	fmt.Printf("[PingRouter] request to handle ...")
	fmt.Printf("request msg id:%v data:%v \n", req.GetMessageID(), string(req.GetData()))

	if err := req.GetConn().Send(1, []byte("ping...ping...ping")); err != nil {
		fmt.Println("send err:", err)
	}

}

func main() {

	server := znet.NewServer()
	server.RegisterRouter(&PingRouter{})
	server.Run()
}
