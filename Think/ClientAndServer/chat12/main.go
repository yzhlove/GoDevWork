package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {

	conn, err := net.DialTimeout("tcp", ":8800", time.Second*5)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(conn, "CONNECT 127.0.0.1:1238 HTTP/1.1\r\n\r\n")

	var data = make([]byte, 1024)
	for {
		n, err := conn.Read(data)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Printf("%v\n", string(data[:n]))
	}
}
