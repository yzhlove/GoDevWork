package main

import (
	"fmt"
	"zinx-day2-base1/ziface"
	"zinx-day2-base1/znet"
)

func main() {

	server := znet.NewServer("[zinx] v2.0")
	server.RegisterRouter(&PingRouter{})
	server.Run()

}

type PingRouter struct {
	znet.BaseRouter
}

func (p *PingRouter) BeforeHandle(req ziface.RequestInterface) {
	fmt.Println("[PingHandle] before ...")
	if _, err := req.GetConn().GetTcp().Write([]byte("before ping ...\n")); err != nil {
		fmt.Println("[Handle] before err ", err)
	}
}

func (p *PingRouter) Handle(req ziface.RequestInterface) {
	fmt.Println("[PingHandle] handle ...")
	if _, err := req.GetConn().GetTcp().Write([]byte("ping...ping...ping\n")); err != nil {
		fmt.Println("[Handle] handle err ", err)
	}
}

func (p *PingRouter) AfterHandle(req ziface.RequestInterface) {
	fmt.Println("[PingHandle] after ...")
	if _, err := req.GetConn().GetTcp().Write([]byte("after ping ...\n")); err != nil {
		fmt.Println("[Handle] after err ", err)
	}
}
