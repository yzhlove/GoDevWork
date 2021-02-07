package main

import (
	rpcbase "day1_base1_example"
	"day1_base1_example/codec"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"
)

func startServer(addr chan string) {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("start rpc server on --> ", l.Addr())
	addr <- l.Addr().String()
	rpcbase.Accept(l)
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	addr := make(chan string)
	go startServer(addr)

	conn, err := net.Dial("tcp", <-addr)
	if err != nil {
		panic(err)
	}
	defer func() { conn.Close() }()

	time.Sleep(time.Second)
	json.NewEncoder(conn).Encode(rpcbase.DefaultOption)
	cc := codec.NewGobCodec(conn)

	for i := 0; i < 5; i++ {
		h := &codec.Header{
			ServiceMethod: "Foo.Lucky",
			Seq:           uint64(i),
		}
		cc.Write(h, fmt.Sprintf("rpcbase req %d", h.Seq))
		cc.ReadHeader(h)
		var reply string
		cc.ReadBody(&reply)
		log.Println("reply:----->", reply)
	}
}
