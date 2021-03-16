package main

import (
	day3_base1 "day3-base1-example"
	"day3-base1-example/codec"
	"log"
	"net"
	"sync"
)

type Foo int

type Args struct {
	Num1, Num2 int
}

func (f Foo) Sum(args Args, reply *int) error {
	*reply = args.Num1 + args.Num2
	return nil
}

func startServer(addr chan string) {
	var foo Foo
	if err := day3_base1.Register(&foo); err != nil {
		log.Fatal("register error:", err)
	}
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("listener error:", err)
	}
	log.Println("start network tcp:", l.Addr())
	addr <- l.Addr().String()
	day3_base1.Accept(l)
}

func main() {
	addr := make(chan string, 1)
	go startServer(addr)
	startClient(<-addr)
}

func startClient(address string) {
	client, err := day3_base1.Dial("tcp", address, &day3_base1.Auth{Code: day3_base1.MagicCode, Type: codec.Gob})
	if err != nil {
		panic(err)
	}
	defer client.Close()

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			args := Args{Num2: 1000, Num1: 2000 * i}
			var reply int
			if err := client.Call("Foo.Sum", args, &reply); err != nil {
				log.Fatal("call function error:", err)
			}
			log.Printf("%d + %d = %d", args.Num1, args.Num2, reply)
		}(i)
	}
	wg.Wait()
}
