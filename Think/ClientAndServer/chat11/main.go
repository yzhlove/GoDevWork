package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"sync"
	"time"
)

func main() {
	fmt.Println(start())
}

func start() error {

	server := &http.Server{Addr: ":8800"}
	server.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("method:%s host:%s remote:%s \n", r.Method, r.Host, r.RemoteAddr)
		fmt.Printf("headers:%v \n", r.Header)

		if r.Method != http.MethodConnect {
			http.Error(w, "this is http proxy", http.StatusMethodNotAllowed)
			return
		}

		conn, _, err := w.(http.Hijacker).Hijack()
		if err != nil {
			http.Error(w, "http not support hijack", http.StatusMethodNotAllowed)
			return
		}
		fmt.Println("host addr:", r.URL.Host)
		tcpAddr, err := net.ResolveTCPAddr("tcp4", r.URL.Host)
		if err != nil {
			http.Error(w, "resolve tcp addr error", http.StatusInternalServerError)
			return
		}
		fmt.Println("parse host:", tcpAddr.String())
		back, err := net.DialTimeout("tcp", tcpAddr.String(), time.Second*5)
		if err != nil {
			http.Error(w, "tcp conn error", http.StatusInternalServerError)
			return
		}
		conn.Write([]byte("HTTP/1.0 200 Connection Established\r\n\r\n"))
		//数据透传
		tunnelTransport(conn, back)
	})
	return server.ListenAndServe()
}

func transfer(in, out net.Conn) {
	defer in.Close()
	defer out.Close()
	io.Copy(in, out)
}

func tunnelTransport(in, out net.Conn) {
	var wg sync.WaitGroup
	wg.Add(2)
	f := func(in, out net.Conn) {
		defer wg.Done()
		transfer(in, out)
	}
	go f(in, out)
	go f(out, in)
	wg.Wait()
}
