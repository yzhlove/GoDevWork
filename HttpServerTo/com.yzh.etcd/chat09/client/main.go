package main

import (
	"WorkSpace/GoDevWork/HttpServerTo/com.yzh.etcd/chat09/discovery"
	"flag"
	"time"
)

var (
	key  string
	addr string
)

func init() {
	flag.StringVar(&key, "key", "snowflak", "register key")
	flag.StringVar(&addr, "addr", "127.0.0.1:1234", "address")
	flag.Parse()
}

func main() {

	manager, err := discovery.New(discovery.ETCD)
	if err != nil {
		panic(err)
	}
	manager.Register(key, addr)
	for {
		time.Sleep(time.Second * 5)
	}
}
