package tcp

import (
	"WorkSpace/GoDevWork/HttpServerTo/tcpCache/cache"
	"net"
)

type Server struct {
	cache.Cache
}

func (s *Server) Listen() {
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic("tcpErr: " + err.Error())
	}
	for {
		if conn, err := l.Accept(); err != nil {
			panic(err)
		} else {
			go s.process(conn)
		}
	}
}

func New(c cache.Cache) *Server {
	return &Server{c}
}
