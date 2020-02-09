package tcp

import (
	"WorkSpace/GoDevWork/HttpServerTo/tcpRebalanceCache/cache"
	"WorkSpace/GoDevWork/HttpServerTo/tcpRebalanceCache/cluster"
	"net"
)

type Server struct {
	cache.Cache
	cluster.Node
}

func (s *Server) Listen() {
	lis, err := net.Listen("tcp", s.Address()+":1235")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			panic(err)
		}
		go s.process(conn)
	}
}

func New(c cache.Cache, n cluster.Node) *Server {
	return &Server{c, n}
}
