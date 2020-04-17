package geeseven

import "strings"

type router struct {
	roots   map[string]*trieTreeNode
	handles map[string]HandleFunc
}

func newRouter() *router {
	return &router{
		roots:   make(map[string]*trieTreeNode, 4),
		handles: make(map[string]HandleFunc),
	}
}

func parse(part string) []string {
	strs := strings.Split(part, "/")
	parts := make([]string, 0, len(strs)>>1)
	for _, str := range strs {
		if v := strings.TrimSpace(str); v != "" {
			parts = append(parts, v)
			if strings.HasPrefix(v, "*") {
				break
			}
		}
	}
	return parts
}

func (r *router) addRouter(method, pattern string, handle HandleFunc) {
	parts := parse(pattern)
	key := method + "-" + pattern
	if _, ok := r.roots[method]; !ok {
		r.roots[method] = &trieTreeNode{}
	}
	r.roots[method].treeInsert(pattern, parts)
	r.handles[key] = handle
}

func (r *router) getRouter(method, pattern string) (*trieTreeNode, map[string]string) {
	parts := parse(pattern)
	params := make(map[string]string, len(parts)>>1)
	if root, ok := r.roots[method]; ok {
		if node := root.treeSearch(parts); node != nil {
			for index, part := range parse(node.pattern) {
				if strings.HasPrefix(part, ":") {
					params[strings.TrimLeft(part, ":")] = parts[index]
				}
				if strings.HasPrefix(part, "*") {
					params[strings.TrimLeft(part, "*")] = strings.Join(parts[index:], "/")
				}
			}
			return node, params
		}
	}
	return nil, nil
}

func (r *router) handle(ctx *Context) {
	if node, params := r.getRouter(ctx.Method, ctx.Path); node != nil {
		key := ctx.Method + "-" + node.pattern
		ctx.Params = params
		ctx.handles = append(ctx.handles, r.handles[key])
	} else {
		ctx.handles = append(ctx.handles, ctx.For404Func())
	}
	ctx.Next()
}
