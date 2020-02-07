package main

import (
	"WorkSpace/GoDevWork/HttpServerTo/tcpFirstCache/cache"
	"WorkSpace/GoDevWork/HttpServerTo/tcpFirstCache/tcp"
)

func main() {
	tcp.New(cache.New(cache.MEMORY)).Listen()
}
