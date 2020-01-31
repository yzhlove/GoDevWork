package main

import (
	"WorkSpace/GoDevWork/HttpServerTo/tcpCache/cache"
	"WorkSpace/GoDevWork/HttpServerTo/tcpCache/tcp"
)

func main() {
	tcp.New(cache.New(cache.MEMORY)).Listen()
}
