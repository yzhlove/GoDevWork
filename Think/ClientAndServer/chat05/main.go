package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var path = "/Users/yurisa/Develop/GoWork/src/WorkSpace/GoDevWork/lumberjack.log"

func main() {

	stat := make(chan struct{})

	go func() {
		fmt.Println(server{address: ":1234"}.httpListen())
	}()

	time.Sleep(time.Second)

	go func() {
		defer func() {
			stat <- struct{}{}
		}()
		httpDown()
		fmt.Println("down succeed!")
	}()

	<-stat
	time.Sleep(time.Second)
}

type server struct {
	address string
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	conn, _, err := w.(http.Hijacker).Hijack()
	if err != nil {
		panic(err)
	}
	var reader = bufio.NewReader(f)
	//var buffer = make([]byte, 128)
	conn.Write([]byte("HTTP/1.1 200 OK\r\n"))
	conn.Write([]byte("Transfer-Encoding: chunked\r\n\r\n"))

	var number int

	for {
		if res, err := reader.ReadBytes('\n'); err != nil {
			log.Println("reader error:", err)
			break
		} else {
			res = bytes.Trim(res, "\n")
			fmt.Println("read server => ", string(res))
			conn.Write([]byte(fmt.Sprintf("%d\r\n%s\r\n", len(res), string(res))))
		}
	}
	fmt.Println("send data => ", number)
	conn.Write([]byte("0\r\n\r\n"))
	conn.Close()
}

func (s server) httpListen() error {
	http.Handle("/down", &s)
	return http.ListenAndServe(s.address, nil)
}

func httpDown() {

	f, err := os.Create("./copy.log")
	if err != nil {
		panic(err)
	}

	resp, err := http.Get("http://127.0.0.1:1234/down")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var reader = bufio.NewReader(resp.Body)
	var buffer = make([]byte, 1024)
	var cnt int
	for {
		if n, err := reader.Read(buffer); err != nil {
			log.Println("http reader error:", err)
			break
		} else {
			cnt += n
			f.Write(buffer[:n])
		}
	}
	f.Close()
	fmt.Println("write file size =>", cnt)
}
