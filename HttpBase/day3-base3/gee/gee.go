package gee

import (
	"log"
	"net/http"
)

type LogicFunc func(ctx *Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (e *Engine) addRouter(method, path string, handle LogicFunc) {
	log.Printf("Router %4s - %s", method, path)
	e.router.addRouter(method, path, handle)
}

func (e *Engine) GET(path string, handle LogicFunc) {
	e.addRouter("GET", path, handle)
}

func (e *Engine) POST(path string, handle LogicFunc) {
	e.addRouter("POST", path, handle)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	e.router.handle(newContext(w, r))
}
