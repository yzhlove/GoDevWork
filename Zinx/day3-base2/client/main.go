package main

import (
	"fmt"
	"io"
	"net"
	"time"
	"zinx-day3-base2/znet"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		panic(err)
	}
	var cid uint32
	for {
		dp := znet.NewDataPack()
		msg, err := dp.Pack(znet.NewMessagePackage(cid, []byte("zinx server v0.5 client test message")))
		if err != nil {
			panic(err)
		}
		if _, err = conn.Write(msg); err != nil {
			fmt.Println("write message err: ", err)
			return
		}

		head := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, head); err != nil {
			fmt.Println("read head err :", err)
			return
		}
		msgHead, err := dp.Unpack(head)
		if err != nil {
			fmt.Println("client unpack err:", err)
			return
		}

		if msgHead.GetDataLen() > 0 {
			if msg, ok := msgHead.(*znet.Message); ok {
				msg.Data = make([]byte, msg.GetDataLen())
				if _, err := io.ReadFull(conn, msg.Data); err != nil {
					fmt.Println("server unpack data err:", err)
					return
				}
				fmt.Printf("[message] ID:%d len:%d message:%v \n", msg.ID, msg.Length, string(msg.Data))
			} else {
				panic("type of to message err")
			}
		}
		cid++
		time.Sleep(time.Second)
	}
}
