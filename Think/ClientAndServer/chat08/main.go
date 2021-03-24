package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
)

var path = "/Users/yostar/workSpace/gowork/src/GoDevWork/1.jpg"

func main() {

	ch := make(chan struct{})

	go func() {
		s := server{address: ":1234"}
		fmt.Println(s.httpListen())
	}()

	go func() {
		httpGet()
		fmt.Println("down over.")
	}()

	<-ch
}

type server struct {
	address string
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	f, err := os.Open(path)
	if err != nil {
		log.Println("open files error:", err)
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
	defer conn.Close()

	setChunkHeader(conn)
	var data = make([]byte, 1024)
	for {
		n, err := f.Read(data)
		if err != nil {
			if err != io.EOF {
				fmt.Println("send data error:", err)
			}
			break
		}
		setData(conn, data[:n])
	}
	setStopChunk(conn)
}

func setChunkHeader(conn net.Conn) {
	conn.Write([]byte("HTTP/1.1 200 OK\r\n"))
	conn.Write([]byte("Transfer-Encoding: chunked\r\n\r\n"))
}

func setData(conn net.Conn, data []byte) {
	conn.Write([]byte(fmt.Sprintf("%x\r\n", len(data))))
	conn.Write(data)
	conn.Write([]byte("\r\n"))
}

func setStopChunk(conn net.Conn) {
	conn.Write([]byte("0\r\n\r\n"))
}

func (s *server) httpListen() error {
	http.Handle("/down", s)
	return http.ListenAndServe(s.address, nil)
}

func httpGet() {
	f, err := os.Create("./succeed.jpg")
	if err != nil {
		log.Println("create file error:", err)
		return
	}
	defer f.Close()

	resp, err := http.Get("http://localhost:1234/down")
	if err != nil {
		log.Println("http error:", err)
		return
	}
	defer resp.Body.Close()

	var data = make([]byte, 1024)
	for {
		n, err := resp.Body.Read(data)
		if err != nil {
			if err != io.EOF {
				fmt.Println("resp body error:", err)
			}
			break
		}
		f.Write(data[:n])
	}
}
