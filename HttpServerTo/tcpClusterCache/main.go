package main

import (
	"WorkSpace/GoDevWork/HttpServerTo/tcpClusterCache/cache"
	"WorkSpace/GoDevWork/HttpServerTo/tcpClusterCache/cluster"
	"WorkSpace/GoDevWork/HttpServerTo/tcpClusterCache/http"
	"WorkSpace/GoDevWork/HttpServerTo/tcpClusterCache/tcp"
	"flag"
	"log"
)

var (
	typ        int
	nodePort   string
	httpPort   string
	gossipPort int
	clusters   string
)

func init() {
	flag.IntVar(&typ, "type", cache.MEMORY, "cache type")
	flag.StringVar(&nodePort, "tcp", "1234", "node port address")
	flag.StringVar(&httpPort, "http", "1235", "http port address")
	flag.IntVar(&gossipPort, "gossip", 7777, "gossip port")
	flag.StringVar(&clusters, "cluster", "", "cluster address")
	flag.Parse()
	log.Println("type =>", typ)
	log.Println("nodePort =>", nodePort)
	log.Println("httpPort =>", httpPort)
	log.Println("gossipPort =>", gossipPort)
	log.Println("cluster =>", clusters)
}

func main() {
	c := cache.New(cache.MEMORY)
	n, err := cluster.New(nodePort, httpPort, clusters, gossipPort)
	if err != nil {
		panic(err)
	}
	go tcp.New(c, n).Listen()
	http.New(c, n).Listen()
}
