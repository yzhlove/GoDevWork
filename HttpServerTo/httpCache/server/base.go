package server

import (
	"WorkSpace/GoDevWork/HttpServerTo/httpCache/cache"
	"net/http"
)

type Server struct {
	cache.Cache
}

func (s *Server) Listen() {
	http.Handle("/cache/", s.cacheHand())
	http.Handle("/status/", s.statusHand())
	_ = http.ListenAndServe(":1234", nil)
}

func New(c cache.Cache) *Server {
	return &Server{Cache: c}
}
