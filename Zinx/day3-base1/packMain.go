package main

import (
	"fmt"
	"io"
	"net"
	"zinx-day3-base1/znet"
)

func serverTest() {

	listener, err := net.Listen("tcp", "0.0.0.0:7777")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("server accept err:", err)
			continue
		}

		go func(conn net.Conn) {
			dp := znet.NewDataPack()
			for {
				headData := make([]byte, dp.GetHeadLen())
				if _, err := io.ReadFull(conn, headData); err != nil {
					fmt.Println("read headData err:", err)
					return
				}
				msgHead, err := dp.Unpack(headData)
				if err != nil {
					fmt.Println("server unpack err:", err)
					return
				}

				if msgHead.GetDataLen() > 0 {
					msg := msgHead.(*znet.Message)
					msg.Data = make([]byte, msg.DataLen)
					if _, err := io.ReadFull(conn, msg.Data); err != nil {
						fmt.Println("server unpack data err:", err)
						return
					}
					fmt.Println("==> read ID:", msg.ID, " Len:", msg.DataLen, " Data:", string(msg.Data))
				}
			}
		}(conn)
	}

}
