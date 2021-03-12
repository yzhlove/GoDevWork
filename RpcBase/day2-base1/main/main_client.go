package main

import (
	day2_base1 "day2-base1-example"
	"day2-base1-example/codec"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"
)

//func main() {
//	testMain()
//}

func testMain() {
	address := make(chan string, 1)
	go startRpcSvc(address)
	startRpcClient(<-address)
}

func startRpcSvc(address chan string) {
	if l, err := net.Listen("tcp", ":0"); err != nil {
		log.Fatal("rpc server error:", err)
	} else {
		log.Println("start rpc server on listener:", l.Addr().String())
		address <- l.Addr().String()
		day2_base1.Accept(l)
	}
}

func startRpcClient(address string) {

	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal("rpc client error:", err)
	}
	defer conn.Close()

	//发送消息头
	var defMsg = &day2_base1.MsgHead{Type: codec.GOB, Code: day2_base1.MagicCode}
	if err := json.NewEncoder(conn).Encode(defMsg); err != nil {
		log.Fatal("json encoder error:", err)
	}

	cc := codec.NewGOBCodec(conn)

	for i := 0; i < 10; i++ {

		//发送数据
		method := &codec.Header{SvcMethod: fmt.Sprintf("A:[%d]", i+1), Seq: uint64(i) + 1}
		params := fmt.Sprintf("A func is args:(%d)", i)
		if err := cc.Writer(method, params); err != nil {
			log.Fatal("call error:", err)
		}

		time.Sleep(time.Second)

		//读取数据
		toHead := &codec.Header{}
		if err := cc.ReadHeader(toHead); err != nil {
			log.Fatal("read server head error:", err)
		}
		log.Println("server head ---> ", toHead)

		var reply string
		if err := cc.ReadBody(&reply); err != nil {
			log.Fatal("read server body error:", err)
		}
		log.Println("server body --->", reply)
		fmt.Println()
	}

}
