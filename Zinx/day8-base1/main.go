package main

import (
	"fmt"
	"zinx-day8-base1/ziface"
	"zinx-day8-base1/znet"
)

func main() {

	server := znet.NewTcpServer()
	server.SetOnConnStart(DoConnStart)
	server.SetOnConnStop(DoConnStop)
	server.RegisterRouter(0, &pingRouter{})
	server.RegisterRouter(1, &helloRouter{})
	server.Run()
}

type pingRouter struct {
	znet.BaseRouter
}

func (ping *pingRouter) Handle(req ziface.RequestInterface) {
	fmt.Println("[ping] msg id:", req.GetMsgID(), " data:", string(req.GetData()))
	if err := req.GetConn().SendBuf(0, []byte("server:ping ...")); err != nil {
		fmt.Printf("[ping] err:%v \n", err)
	}
}

type helloRouter struct {
	znet.BaseRouter
}

func (hello *helloRouter) Handle(req ziface.RequestInterface) {
	fmt.Println("[hello] msg id:", req.GetMsgID(), " data:", string(req.GetData()))
	if err := req.GetConn().SendBuf(1, []byte("server:hello world .")); err != nil {
		fmt.Printf("[hello] err:%v \n", err)
	}
}

func DoConnStart(conn ziface.ConnectionInterface) {
	fmt.Println("--- conn start ---")
	conn.SetAttribute("name", "yzh")
	conn.SetAttribute("email", "lcmm5201314@gmail.com")
	if err := conn.SendBuf(2, []byte("do conn start and set attribute .")); err != nil {
		fmt.Println("do conn send msg err:", err)
	}
}

func DoConnStop(conn ziface.ConnectionInterface) {
	if name, ok := conn.GetAttribute("name"); ok {
		fmt.Println("Get attribute name => ", name)
	}
	if email, ok := conn.GetAttribute("email"); ok {
		fmt.Println("Get attribute email => ", email)
	}
	fmt.Println("do conn stop .")
}
