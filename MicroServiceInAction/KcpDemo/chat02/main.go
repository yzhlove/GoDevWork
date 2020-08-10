package main

import (
	"errors"
	"fmt"
	"github.com/xtaci/kcp-go"
	"io"
	"log"
	"math/rand"
	"net"
	"strings"
	"time"
)

func main() {

	go server()
	for i := 0; i < 5; i++ {
		tp := i
		go client(tp)
	}
	select {}
}

func server() {
	lis, err := kcp.ListenWithOptions(":1234", nil, 10, 3)
	if err != nil {
		panic(err)
	}

	_exit := make(chan error, 1)
	_stringBuffer := make(chan string, 1)

	_Reader := func(conn net.Conn) {
		var buffer = make([]byte, 1024)
		for {
			if n, err := conn.Read(buffer); err != nil {
				if errors.Is(err, io.EOF) {
					_exit <- fmt.Errorf("read eof : %v", err)
					return
				}
				_exit <- err
				return
			} else {
				log.Println("server read client message:", string(buffer[:n]))
				_stringBuffer <- string(buffer[:n])
			}
		}
	}

	_Writer := func(conn net.Conn) {
		for msg := range _stringBuffer {
			if _, err := conn.Write([]byte(strings.ToUpper(msg))); err != nil {
				_exit <- err
				return
			}
		}
	}

	for {
		conn, err := lis.AcceptKCP()
		if err != nil {
			panic(err)
		}

		if _, err := conn.Write([]byte("welcome to kcp ...")); err != nil {
			log.Println("first conn err:", err)
			continue
		}

		go _Writer(conn)
		go _Reader(conn)
	}

}

func client(id int) {

	kcpconn, err := kcp.DialWithOptions("localhost:1234", nil, 10, 3)
	if err != nil {
		panic(err)
	}

	go func(conn net.Conn) {
		for {
			var buffer = make([]byte, 1024)
			if n, err := conn.Read(buffer); err != nil {
				log.Println("read err:", err)
				panic(err)
			} else {
				log.Println("client read server message:", string(buffer[:n]))
			}
		}
	}(kcpconn)

	var count int
	for {
		count++
		msg := fmt.Sprintf("[%d] rand dict : %c  count: %d", id, rune(rand.Intn(26)+97), count)
		if _, err := kcpconn.Write([]byte(msg)); err != nil {
			log.Println("send message err:", err)
		}
		time.Sleep(time.Second * 2)
	}
}
