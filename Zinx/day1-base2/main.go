package main

import (
	"log"
	"net"
	"strings"
)

func main() {

	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("[conn accept ] err:" + err.Error())
			return
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	//defer conn.Close()
	log.Println("[server] ", conn.RemoteAddr())
	conn.Write([]byte(strings.ToUpper("ping") + "\r\n"))
}
