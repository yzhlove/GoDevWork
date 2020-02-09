package main

import (
	"WorkSpace/GoDevWork/HttpServerTo/tcpRebalanceCache/cache"
	"WorkSpace/GoDevWork/HttpServerTo/tcpRebalanceCache/cluster"
	"WorkSpace/GoDevWork/HttpServerTo/tcpRebalanceCache/http"
	"WorkSpace/GoDevWork/HttpServerTo/tcpRebalanceCache/tcp"
	"flag"
	"log"
)

var (
	typ      int
	addr     string
	clusters string
)

func init() {
	flag.IntVar(&typ, "type", cache.MEMORY, "cache type")
	flag.StringVar(&addr, "addr", "127.0.0.1", "ip address")
	flag.StringVar(&clusters, "cluster", "", "cluster address")
	flag.Parse()
	log.Println("type:", typ)
	log.Println("address:", addr)
	log.Println("cluster:", clusters)
}

func main() {
	c := cache.New(cache.MEMORY)
	node, err := cluster.New(addr, clusters)
	if err != nil {
		panic(err)
	}
	go tcp.New(c, node).Listen()
	http.New(c, node).Listen()
}
