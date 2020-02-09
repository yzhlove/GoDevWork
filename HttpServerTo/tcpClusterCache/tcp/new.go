package tcp

import (
	"WorkSpace/GoDevWork/HttpServerTo/tcpClusterCache/cache"
	"WorkSpace/GoDevWork/HttpServerTo/tcpClusterCache/cluster"
	"net"
)

type Server struct {
	cache.Cache
	cluster.Node
}

func (s *Server) Listen() {
	l, err := net.Listen("tcp", ":"+s.TcpAddress())
	if err != nil {
		panic(err)
	}
	for {
		if conn, err := l.Accept(); err != nil {
			panic(err)
		} else {
			go s.process(conn)
		}
	}
}

func New(c cache.Cache, n cluster.Node) *Server {
	return &Server{c, n}
}
