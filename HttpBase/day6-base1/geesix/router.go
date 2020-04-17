package geesix

import (
	"net/http"
	"strings"
)

type router struct {
	roots   map[string]*trieNode
	handles map[string]HandleFunc
}

func newRouter() *router {
	return &router{
		roots:   make(map[string]*trieNode, 4),
		handles: make(map[string]HandleFunc, 4),
	}
}

func parseParent(parent string) []string {
	parents := strings.Split(parent, "/")
	values := make([]string, 0, len(parents)>>1)
	for _, item := range parents {
		if value := strings.TrimSpace(item); len(value) > 0 {
			values = append(values, value)
			if strings.HasPrefix(value, "*") {
				break
			}
		}
	}
	return values
}

func (r *router) addRouter(method, pattern string, handle HandleFunc) {
	parents := parseParent(pattern)
	key := method + "-" + pattern
	if _, ok := r.roots[method]; !ok {
		r.roots[method] = &trieNode{}
	}
	r.roots[method].InsertNode(pattern, parents)
	r.handles[key] = handle
}

func (r *router) getRouter(method, path string) (*trieNode, map[string]string) {
	parents := parseParent(path)
	params := make(map[string]string, len(parents)>>1)
	if root, ok := r.roots[method]; ok {
		if tNode := root.SearchNode(parents); tNode != nil {
			values := parseParent(tNode.path)
			for k, v := range values {
				if strings.HasPrefix(v, ":") {
					params[strings.TrimLeft(v, ":")] = parents[k]
				}
				if strings.HasPrefix(v, "*") {
					params[strings.TrimLeft(v, "*")] = strings.Join(parents[k:], "/")
				}
			}
			return tNode, params
		}
	}
	return nil, nil
}

func (r *router) getRouters(method string) []*trieNode {
	if root, ok := r.roots[method]; ok {
		return root.Travel()
	}
	return nil
}

func (r *router) handle(ctx *Context) {
	if node, params := r.getRouter(ctx.Method, ctx.Path); node != nil {
		key := ctx.Method + "-" + node.path
		ctx.Params = params
		ctx.handlers = append(ctx.handlers, r.handles[key])
	} else {
		ctx.handlers = append(ctx.handlers, func(ctx *Context) {
			ctx.String(http.StatusNotFound, "404 Not Found! [%s] \n", ctx.Path)
		})
	}
	ctx.Next()
}
