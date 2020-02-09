package http

import (
	"WorkSpace/GoDevWork/HttpServerTo/tcpClusterCache/cache"
	"WorkSpace/GoDevWork/HttpServerTo/tcpClusterCache/cluster"
	"net/http"
)

type Server struct {
	cache.Cache
	cluster.Node
}

func (s *Server) Listen() {
	//http.Handle("/cache/", nil)
	//http.Handle("/status", nil)
	http.Handle("/cluster", s.clusterHandle())
	_ = http.ListenAndServe(":"+s.HttpAddress(), nil)
}

func New(c cache.Cache, n cluster.Node) *Server {
	return &Server{c, n}
}
