package main

import (
	"bufio"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var path = "/Users/yostar/workSpace/gowork/src/GoDevWork/1.jpg"

func main() {

	ch := make(chan struct{})

	go func() {
		s := server{address: ":1234"}
		fmt.Println(s.httpListen())
	}()

	time.Sleep(time.Second)

	go func() {
		//defer func() {
		//	ch <- struct{}{}
		//}()
		httpGet()
		fmt.Println("write file ok.")
	}()

	<-ch
	time.Sleep(time.Second)
	fmt.Println("exit...")
}

type server struct {
	address string
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open(path)
	if err != nil {
		log.Println("open file error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()

	conn, _, err := w.(http.Hijacker).Hijack()
	if err != nil {
		log.Println("hijack error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	conn.Write([]byte("HTTP/1.1 200 OK\r\n"))
	conn.Write([]byte("X-Content-Type-Options: nosniff\r\n"))
	conn.Write([]byte("Transfer-Encoding: chunked\r\n\r\n"))

	var reader = bufio.NewReader(f)
	var buffer = make([]byte, 1024)
	var cnt int

	for {
		if n, err := reader.Read(buffer); err != nil {
			log.Println("read error:", err)
			break
		} else {
			cnt += n
			str := buffer[:n]
			conn.Write([]byte(
				fmt.Sprintf("%d\r\n%s\r\n", len(str), str)))
		}
	}

	fmt.Println("send data size=>", cnt)
	binary.Write(conn, binary.BigEndian, []byte("0\r\n\r\n"))
	conn.Close()
}

func (s *server) httpListen() error {
	http.Handle("/down", s)
	return http.ListenAndServe(s.address, nil)
}

func httpGet() {

	f, err := os.Create("./copy.jpg")
	if err != nil {
		log.Println("create file error:", err)
		return
	}
	defer f.Close()

	resp, err := http.Get("http://127.0.0.1:1234/down")
	if err != nil {
		log.Println("get error:", err)
		return
	}
	defer resp.Body.Close()
	var reader = bufio.NewReader(resp.Body)
	var cnt int

	for {
		s, err := reader.ReadString('=')
		if err != nil {
			log.Println("reader error:", err)
			break
		} else {
			fmt.Println("read s ===> ", s)
			cnt += len(s)
			if res, err := base64.StdEncoding.DecodeString(s); err != nil {
				log.Println("base64 error:", err)
				break
			} else {
				f.Write(res)
			}
		}
	}

	fmt.Println("receive data size=>", cnt)
}
