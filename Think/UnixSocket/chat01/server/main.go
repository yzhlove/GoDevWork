package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {

	var file string = "test.sock"

start:
	l, err := net.Listen("unix", file)
	if err != nil {
		log.Println("unix socket create error:", err)
		if err := os.Remove(file); err != nil {
			log.Println("remove socket file error:", err)
			os.Exit(0)
		}
		goto start
	}

	fmt.Println("unix socket create succeed:", l.Addr().String())
	defer l.Close()

	idx := 0
	for {
		conn, err := l.Accept()
		idx++
		fmt.Println("current idx ==> ", idx)
		if err != nil {
			log.Println("accept error:", err)
			continue
		}
		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		io.Copy(conn, conn)
	}
}
