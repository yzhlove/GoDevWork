package gee

import "net/http"

type router struct {
	handlers map[string]LogicFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]LogicFunc)}
}

func (r *router) addRouter(method, pattern string, h LogicFunc) {
	key := method + "-" + pattern
	r.handlers[key] = h
}

func (r *router) handle(ctx *Context) {
	k := ctx.Method + "-" + ctx.Path
	if h, ok := r.handlers[k]; ok {
		h(ctx)
	} else {
		ctx.String(http.StatusNotFound, "404 NOT FOUND:%s\n", ctx.Path)
	}
}
