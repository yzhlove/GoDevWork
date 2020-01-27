package main

import (
	"WorkSpace/GoDevWork/HttpServerTo/httpCache/cache"
	"WorkSpace/GoDevWork/HttpServerTo/httpCache/server"
)

func main() {
	server.New(cache.New("IN_MEMORY")).Listen()
}
