package main

import (
	"WorkSpace/GoDevWork/HttpServerTo/com.yzh.etcd/chat09/discovery"
	"fmt"
	"time"
)

func main() {

	mgr, err := discovery.New(discovery.ETCD)
	if err != nil {
		panic(err)
	}

	for {
		fmt.Println("nodes => ", mgr.GetNodes())
		time.Sleep(time.Second * 5)
	}

}
