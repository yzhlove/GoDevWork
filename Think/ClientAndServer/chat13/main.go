package main

import (
	"fmt"
	"net"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		if err := httpService(":1234"); err != nil {
			fmt.Println("http service error:", err)
		}
	}()
	go func() {
		defer wg.Done()
		if err := tcpService(":1238"); err != nil {
			fmt.Println("tcp service eerror:", err)
		}
	}()
	wg.Wait()
}

func httpService(addr string) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("<<<< http request >>>>")
		count := 0
		for i := 0; i < 100; i++ {
			w.Write([]byte(fmt.Sprintf("http send data %d", count+i)))
			count++
		}
	})
	fmt.Println("monitor http to ", addr)
	return http.ListenAndServe(addr, nil)
}

func tcpService(addr string) error {
	fmt.Println("monitor tcp to ", addr)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	for {
		if conn, err := l.Accept(); err != nil {
			fmt.Println("take conn error:", err)
		} else {
			go func(c net.Conn) {
				defer func() {
					fmt.Println("关闭连接closing...")
					c.Close()
				}()
				fmt.Println("<<<< tcp connection >>>>")
				for i := 0; i < 100; i++ {
					c.Write([]byte(fmt.Sprintf("tcp send data %d", i)))
				}
			}(conn)
		}
	}
}
