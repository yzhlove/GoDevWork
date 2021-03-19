package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	var address = "127.0.0.1:1234"
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("start server ...")
		var s server
		errCh := make(chan error)
		go func() {
			if err := s.listener(address); err != nil {
				log.Println("server tcp error:", err)
				errCh <- err
			}
		}()
		go func() {
			if err := s.httpListener(); err != nil {
				log.Println("server http error:", err)
				errCh <- err
			}
		}()
		if err := <-errCh; err != nil {
			log.Println("server error:", err)
		}
	}()
	time.Sleep(time.Second)
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("start client ...")
		var c client
		errCh := make(chan error)
		go func() {
			if err := c.listener(address); err != nil {
				log.Println("client tcp error:", err)
				errCh <- err
			}
		}()
		go func() {
			if err := c.httpListener(); err != nil {
				log.Println("client http error:", err)
				errCh <- err
			}
		}()
		if err := <-errCh; err != nil {
			log.Println("client error:", err)
		}
	}()
	wg.Wait()
	fmt.Println("running over ...")
}

type server struct{}

func (s *server) listener(address string) error {
	l, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	return s.accept(l)
}

func (s *server) accept(l net.Listener) error {
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("accept error", err)
			return err
		}
		go s.serverConn(conn)
	}
}

func (s *server) serverConn(conn net.Conn) {
	reader := bufio.NewReader(conn)
	write := bufio.NewWriter(conn)
	for {
		if res, err := reader.ReadString('\n'); err != nil {
			log.Println("reader error:", err)
			return
		} else {
			if _, err := write.WriteString(strings.ToUpper(res)); err != nil {
				log.Println("write str:", err)
				return
			}
			if err := write.Flush(); err != nil {
				log.Println("write flush error:", err)
				return
			}
		}
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("receive http request.")
	//Transfer-Encoding:chunk
	conn, _, err := w.(http.Hijacker).Hijack()
	if err != nil {
		log.Println("hijack error:", err)
		return
	}
	conn.Write([]byte("HTTP/1.1 200 OK\r\n"))
	conn.Write([]byte("Transfer-Encoding:chunked\r\n\r\n"))
	if _, err := conn.Write([]byte("connection succeed.\n")); err != nil {
		log.Println("<write data error >")
		return
	}
	s.serverConn(conn)
}

func (s *server) httpListener() error {
	http.Handle("/rpc", s)
	return http.ListenAndServe(":6666", nil)
}

type client struct{}

func (c *client) listener(address string) error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}

	go func() {
		reader := bufio.NewReader(conn)
		for {
			str, err := reader.ReadString('\n')
			if err != nil {
				log.Println("read string error:", err)
				panic(err)
			}
			fmt.Println("read string:", str)
		}
	}()

	writer := bufio.NewWriter(conn)
	count := 0
	for {
		str := fmt.Sprintf("(%d) writer message.\n", count)
		fmt.Println("--> ", str)
		count++
		if _, err := writer.WriteString(str); err != nil {
			return err
		}
		if err := writer.Flush(); err != nil {
			return err
		}
		time.Sleep(time.Second * 5)
	}
}

func (c *client) httpListener() error {

	var count = 0
	var client = &http.Client{Timeout: time.Second * 5}
	for {
		str := []byte(fmt.Sprintf("(%d) http request to message.\n", count))
		request, err := http.NewRequest("GET", "http://127.0.0.1:6666/rpc", bytes.NewBuffer(str))
		if err != nil {
			return err
		}
		request.TransferEncoding = []string{"identity"}
		resp, err := client.Do(request)
		if err != nil {
			return err
		}
		buffer := bufio.NewReader(resp.Body)
		for {
			str, err := buffer.ReadString('\n')
			if err != nil {
				log.Println("read data error:", err)
				break
			}
			fmt.Println("http read => ", str)
		}
		resp.Body.Close()
		count++
		time.Sleep(time.Second * 2)
	}
}
