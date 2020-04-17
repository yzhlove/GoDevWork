package geefour

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
		roots:   make(map[string]*trieNode),
		handles: make(map[string]HandleFunc),
	}
}

func parseUrl(path string) []string {
	urls := strings.Split(path, "/")
	values := make([]string, 0, len(urls)-1)
	for _, v := range urls {
		if url := strings.TrimSpace(v); len(v) > 0 {
			values = append(values, url)
			if strings.HasPrefix(url, "*") {
				break
			}
		}
	}
	return values
}

func (r *router) addRouter(method string, path string, handle HandleFunc) {
	values := parseUrl(path)
	key := method + "-" + path
	if _, ok := r.roots[method]; !ok {
		r.roots[method] = &trieNode{}
	}
	r.roots[method].Insert(path, values)
	r.handles[key] = handle
}

func (r *router) getRouter(method string, path string) (*trieNode, map[string]string) {
	values := parseUrl(path)
	params := make(map[string]string, len(values)>>1)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	if tNode := root.Search(values); tNode != nil {
		for index, part := range parseUrl(tNode.path) {
			if strings.HasPrefix(part, ":") {
				params[strings.TrimLeft(path, ":")] = values[index]
			}
			if strings.HasPrefix(part, "*") && len(part) > 1 {
				params[strings.TrimLeft(part, "*")] = strings.Join(values[index:], "/")
				break
			}
		}
		return tNode, params
	}
	return nil, nil
}

func (r *router) getRouters(method string) []*trieNode {
	if root, ok := r.roots[method]; !ok {
		return nil
	} else {
		return root.travel()
	}
}

func (r *router) handle(ctx *Context) {
	if node, params := r.getRouter(ctx.Method, ctx.Path); node != nil {
		ctx.Params = params
		key := ctx.Method + "-" + node.path
		r.handles[key](ctx)
	} else {
		ctx.String(http.StatusNotFound, "404 NOT FOUND:%s\n", ctx.Path)
	}
}
