package day5_base1

import (
	"day5/codec"
	"io"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

const MagicNumber = 0x3b5f

type CheckCode struct {
	Code           int
	Type           codec.Coder
	ConnectTimeout time.Duration
	HandleTimeout  time.Duration
}

type server struct {
	services sync.Map
}

func (s *server) serveConn(conn net.Conn) {

}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "CONNECT" {
		w.Header().Set("Content-Type", "text/plain;charset=utf-8")
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, "405 must CONNECT\n")
		return
	}
	conn, _, err := w.(http.Hijacker).Hijack()
	if err != nil {
		log.Println("rcp hijack ", r.RemoteAddr, ":", err)
		return
	}
	io.WriteString(conn, "HTTP/1.1 200 Connected to Gee RPC\r\n\r\n")
	s.serveConn(conn)
}

func (s *server) HandleHttp() {
	http.Handle("/_geerpc_", s)
	http.Handle("/debug/geerpc", s)
	log.Println("rpc server debug path:/debug/geerpc")
}
