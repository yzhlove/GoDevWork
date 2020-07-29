package main

import (
	"crypto/tls"
	"fmt"
	"strconv"
	"time"
)

func main() {

	conf := &tls.Config{InsecureSkipVerify: true}

	conn, err := tls.Dial("tcp", ":1234", conf)
	if err != nil {
		panic("dial err: " + err.Error())
	}
	defer conn.Close()

	status := make(chan struct{})

	go func() {
		var count int
		for range time.NewTicker(time.Second * 1).C {
			count++
			msg := "hello." + strconv.Itoa(count) + "\n"
			if _, err := conn.Write([]byte(msg)); err != nil {
				status <- struct{}{}
				return
			}
		}
	}()

	go func() {
		buf := make([]byte, 1024)
		for {
			if n, err := conn.Read(buf); err != nil {
				status <- struct{}{}
				return
			} else {
				fmt.Println("read data=>", string(buf[:n]))
			}
		}
	}()

	<-status

}
