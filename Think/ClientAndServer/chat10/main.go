package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"
)

//http代理

func main() {

	ch := make(chan struct{})

	go func() {
		fmt.Println(tcpLis(":2345"))
	}()

	go func() {
		s := &server{address: ":1234"}
		fmt.Println(s.httpLis())
	}()

	go func() {
		httpClient()
	}()

	<-ch

}

type server struct {
	address string
}

func tcpLis(address string) error {
	l, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			return err
		}
		go func(c net.Conn) {
			var data = make([]byte, 1024)
			for {
				n, err := c.Read(data)
				if err != nil {
					if err != io.EOF {
						panic(err)
					}
				}
				c.Write([]byte(strings.ToUpper(string(data[:n]))))
			}
		}(conn)
	}

}

func proxyHttp(w http.ResponseWriter, r *http.Request) {
	dest_conn, err := net.DialTimeout("tcp", r.Host, 10*time.Second)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)
	src_conn, _, err := w.(http.Hijacker).Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	go transfer(dest_conn, src_conn)
	go transfer(src_conn, dest_conn)
}

func handleHttp(w http.ResponseWriter, r *http.Request) {
	resp, err := http.DefaultTransport.RoundTrip(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()
	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request ->", r.Host, r.Method)
	if r.Method == http.MethodConnect {
		proxyHttp(w, r)
	} else {
		handleHttp(w, r)
	}
}

func (s *server) httpLis() error {
	http.Handle("/down", s)
	return http.ListenAndServe(s.address, nil)
}

func transfer(dest io.WriteCloser, src io.ReadCloser) {
	defer dest.Close()
	defer src.Close()
	io.Copy(dest, src)
}

func copyHeader(dst, src http.Header) {
	for k, v := range src {
		for _, r := range v {
			dst.Add(k, r)
		}
	}
}

func httpClient() {

	conn, err := net.DialTimeout("tcp", "127.0.0.1:1234", time.Second*10)
	if err != nil {
		panic(err)
	}

	fmt.Println("conn succeed ")

	io.WriteString(conn, "CONNECT 127.0.0.1:2345 HTTP/1.1\r\n\r\n")
	resp, err := http.ReadResponse(bufio.NewReader(conn), &http.Request{Method: http.MethodConnect})
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var data = make([]byte, 1024)

	go func() {

		tk := time.NewTicker(time.Second)
		count := 0
		for {
			select {
			case <-tk.C:
				count++
				str := fmt.Sprintf("send data %d", count)
				fmt.Println("str -> ", str)
				conn.Write([]byte(str))
			}
		}
	}()

	for {
		n, err := resp.Body.Read(data)
		if err != nil {
			break
		}
		fmt.Println("result -> ", string(data[:n]))
	}

}
