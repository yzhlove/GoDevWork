package main

import "WorkSpace/GoDevWork/HttpServerTo/tcpClient/tcp"

func main() {
	cli := tcp.New()
	//cli.Run(&tcp.Result{Opt: "set", Key: "yzh", Value: "1996"})
	//cli.Run(&tcp.Result{Opt: "get", Key: "yzh"})
	//cli.Run(&tcp.Result{Opt: "del", Key: "yzh"})
	//cli.Run(&tcp.Result{Opt: "get", Key: "yzh"})

	rs := []*tcp.Result{
		&tcp.Result{Opt: "set", Key: "yzh", Value: "1996"},
		&tcp.Result{Opt: "get", Key: "yzh"},
		&tcp.Result{Opt: "set", Key: "yzh", Value: "1998"},
		&tcp.Result{Opt: "get", Key: "yzh"},
		&tcp.Result{Opt: "del", Key: "yzh"},
		&tcp.Result{Opt: "get", Key: "yzh"},
	}
	cli.RunPipeline(rs)
}
