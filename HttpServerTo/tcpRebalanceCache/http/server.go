package http

import (
	"WorkSpace/GoDevWork/HttpServerTo/tcpRebalanceCache/cache"
	"WorkSpace/GoDevWork/HttpServerTo/tcpRebalanceCache/cluster"
	"net/http"
)

type Server struct {
	cache.Cache
	cluster.Node
}

func (s *Server) Listen() {
	http.Handle("/cache/", s.cacheHandle())
	http.Handle("/status", s.statusHandle())
	http.Handle("/cluster", s.cacheHandle())
	http.Handle("/rebalanced", s.rebalancedHandle())
	http.ListenAndServe(s.Address()+":1234", nil)
}

func New(c cache.Cache, n cluster.Node) *Server {
	return &Server{c, n}
}
