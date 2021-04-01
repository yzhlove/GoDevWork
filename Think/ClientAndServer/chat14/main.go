package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	log.Println(service(":8800"))
}

func service(address string) error {

	conn, err := net.DialTimeout("tcp", address, 5*time.Second)
	if err != nil {
		return err
	}
	io.WriteString(conn, "CONNECT http://127.0.0.1:1234/ HTTP/1.1\r\n\r\n")

	var data = make([]byte, 1024)

	for {
		n, err := conn.Read(data)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		if bytes.Contains(data[:n], []byte("Connection Established")) {
			conn.Write([]byte("GET / HTTP/1.1\r\nHost: 127.0.0.1:1234\r\n\r\n"))
		}
		fmt.Printf("%v\n", string(data[:n]))
	}
	fmt.Println("over.")
	return nil
}
