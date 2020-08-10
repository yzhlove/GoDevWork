package main

import (
	"errors"
	"fmt"
	"github.com/xtaci/kcp-go"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

func main() {

	lis, err := kcp.ListenWithOptions(":1234", nil, 10, 3)
	if err != nil {
		panic(err)
	}

	go client()

	for {
		conn, err := lis.AcceptKCP()
		if err != nil {
			panic(err)
		}

		go func(conn net.Conn) {
			var buffer = make([]byte, 1024)
			for {
				n, err := conn.Read(buffer)
				if err != nil {
					if errors.Is(err, io.EOF) {
						log.Printf("io.EOF ")
						break
					}
					log.Println("read err:", err)
					break
				}
				log.Printf("show message:%+v \n", string(buffer[:n]))
			}
		}(conn)
	}
}

func client() {
	kcpconn, err := kcp.DialWithOptions("localhost:1234", nil, 10, 3)
	if err != nil {
		log.Println("client err:", err)
		panic(err)
	}
	var count int
	for {
		count++
		msg := "kcp client send message: " + strconv.Itoa(count)
		if n, err := kcpconn.Write([]byte(msg)); err != nil {
			log.Println("kcp client send message err:", err, " msg:", msg)
		} else {
			fmt.Println("client send message succeed. message:", msg, " bytes:", n)
		}
		time.Sleep(time.Second)
	}
}
