package gee

import (
	"fmt"
	"net/http"
)

type LogicFunc func(w http.ResponseWriter, r *http.Request)

type Engine struct {
	router map[string]LogicFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]LogicFunc)}
}

func (e *Engine) addRouter(method, pattern string, handle LogicFunc) {
	key := method + "-" + pattern
	e.router[key] = handle
}

func (e *Engine) GET(pattern string, handle LogicFunc) {
	e.addRouter("GET", pattern, handle)
}

func (e *Engine) POST(pattern string, handle LogicFunc) {
	e.addRouter("POST", pattern, handle)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if h, ok := e.router[key]; ok {
		h(w, r)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND:%s\n", r.URL)
	}
}
