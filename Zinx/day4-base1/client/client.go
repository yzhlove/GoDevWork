package main

import (
	"fmt"
	"io"
	"net"
	"time"
	"zinx-day4-base1/znet"
)

var (
	IPVersion = "tcp"
	Address   = "127.0.0.1:7777"
)

func main() {

	die := make(chan struct{})
	go send(0, []byte("ping...pong"))
	go send(1, []byte("are you ok?"))
	<-die

}

func send(api uint32, data []byte) {
	conn, err := net.Dial(IPVersion, Address)
	if err != nil {
		panic(err)
	}
	for {
		dp := znet.NewDataPack()
		msg, err := dp.Pack(znet.NewMessagePackage(api, data))
		if err != nil {
			fmt.Println("write message err:", err)
			return
		}

		if _, err := conn.Write(msg); err != nil {
			fmt.Println("write data err:", err)
			return
		}

		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, headData); err != nil {
			fmt.Println("read head err:", err)
			return
		}

		msgHead, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("client unpack err:", err)
			return
		}

		if msgHead.GetDataLen() > 0 {
			if msg, ok := msgHead.(*znet.Message); ok {
				msg.Data = make([]byte, msg.GetDataLen())
				if _, err := io.ReadFull(conn, msg.Data); err != nil {
					fmt.Println("server read full err:", err)
					return
				}
				fmt.Printf("[client] id:%d data:%s \n", msg.ID, string(msg.Data))
			} else {
				panic("type of message error")
			}
		}

		time.Sleep(time.Second)
	}
}
