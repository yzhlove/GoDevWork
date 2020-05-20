package znet

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func ClientTest() {
	fmt.Println("Client Test ... Start")
	time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("client start err,exit !")
		return
	}

	for {
		_, err := conn.Write([]byte("hahahaaaa"))
		if err != nil {
			fmt.Println("write error err ", err)
			return
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf error ", err)
			return
		}
		fmt.Printf("server call cnt = %d ,black = %s \n", cnt, buf)
		time.Sleep(time.Second)
	}
}

func Test_Server(t *testing.T) {
	s := NewServer("[zinx] v0.1")
	go ClientTest()
	s.Serve()
}
