package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {

	ch := make(chan struct{})

	go func() {
		defer func() {
			ch <- struct{}{}
		}()
		s := &server{address: ":1234"}
		fmt.Println(s.httpListener())
	}()

	go func() {
		httpGet()
		fmt.Println("read over.")
	}()

	<-ch
}

type server struct {
	ch      chan string
	address string
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	conn, _, err := w.(http.Hijacker).Hijack()
	if err != nil {
		log.Println("hijack error:", err)
	}
	s.testChunk(conn)
	fmt.Println("close connection")
	conn.Close()
}

func (s *server) testChunk(conn net.Conn) {
	conn.Write([]byte("HTTP/1.1 200 OK\r\n"))
	conn.Write([]byte("Transfer-Encoding: chunked\r\n"))
	conn.Write([]byte("Content-Type: text/html;charset=UTF-8\r\n\r\n"))

	for i := 0; i < 100; i++ {
		conn.Write([]byte("7\r\n"))
		conn.Write([]byte("Mozilla\r\n"))
		conn.Write([]byte("9\r\n"))
		conn.Write([]byte("Developer\r\n"))
	}
	conn.Write([]byte("0\r\n\r\n"))
}

func (s *server) setChunkHead(conn net.Conn) {
	conn.Write([]byte("HTTP/1.1 200 OK\r\n"))
	conn.Write([]byte("Content-Type: text/html;charset=utf8\r\n"))
	conn.Write([]byte("X-Content-Type-Options: nosniff\r\n"))
	conn.Write([]byte("Connection: Keep-Alive\r\n"))
	conn.Write([]byte("Transfer-Encoding: chunked\r\n\r\n"))
}

func (s *server) sendBytes(conn net.Conn, str string) {
	ss := fmt.Sprintf(fmt.Sprintf("%d\r\n\r\n<h1>%s</h1>\r\n\r\n", len(str), str))
	fmt.Printf("(%q)\n", ss)
	conn.Write([]byte(ss))
}

func (s *server) httpListener() error {
	http.Handle("/test", s)
	return http.ListenAndServe(s.address, nil)
}

func httpGet() {
	request, err := http.Get("http://127.0.0.1:1234/test")
	if err != nil {
		panic(err)
	}
	defer request.Body.Close()

	reader := bufio.NewReader(request.Body)
	data := make([]byte, 64)

	for {
		if n, err := reader.Read(data); err != nil {
			log.Println("read error:", err)
			return
		} else {
			fmt.Printf("read data:(%q)\n", string(data[:n]))
		}
	}

}
