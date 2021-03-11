package main

import (
	day2_base1 "day2-base1-example"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

func startServer(addr chan string) {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("network error", err)
	}
	log.Println("start rpc server ...")
	addr <- l.Addr().String()
	day2_base1.Accept(l)
}

func main() {

	addr := make(chan string)
	go startServer(addr)

	client, err := day2_base1.Dial("tcp", <-addr)
	if err != nil {
		log.Fatal("client start error", err)
	}
	defer client.Close()

	time.Sleep(time.Second)

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			args := fmt.Sprintf("gee rpc seq:%d", i)
			var reply string
			if err := client.Call("Foo.Sum", args, &reply); err != nil {
				log.Fatal("call err:", err)
			}
			log.Println("reply:", reply)
		}(i)
	}
	wg.Wait()
}
