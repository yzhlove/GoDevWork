package gee

import "net/http"

type LogicFunc func(ctx *Context)

type Engine struct {
	r *router
}

func New() *Engine {
	return &Engine{r: newRouter()}
}

func (e *Engine) addRouter(method, pattern string, h LogicFunc) {
	e.r.addRouter(method, pattern, h)
}

func (e *Engine) GET(pattern string, h LogicFunc) {
	e.addRouter("GET", pattern, h)
}

func (e *Engine) POST(pattern string, h LogicFunc) {
	e.addRouter("POST", pattern, h)
}

func (e *Engine) Run(add string) error {
	return http.ListenAndServe(add, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	e.r.handle(newContext(w, r))
}
