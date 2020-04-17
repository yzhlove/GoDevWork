package geefive

import (
	"log"
	"net/http"
	"strings"
)

type HandleFunc func(ctx *Context)

type RouterGroup struct {
	prefix  string
	middles []HandleFunc
	parent  *RouterGroup
	engine  *Engine
}

type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

func NewEngine() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (group *RouterGroup) NewGroup(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: engine.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) Use(middles ...HandleFunc) {
	group.middles = append(group.middles, middles...)
}

func (group *RouterGroup) addRouter(method, comp string, handle HandleFunc) {
	pattern := group.prefix + comp
	log.Printf("Router %4s - %s", method, pattern)
	group.engine.router.addRouter(method, pattern, handle)
}

func (group *RouterGroup) GET(pattern string, handle HandleFunc) {
	group.addRouter("GET", pattern, handle)
}

func (group *RouterGroup) POST(pattern string, handle HandleFunc) {
	group.addRouter("POST", pattern, handle)
}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var middles []HandleFunc
	for _, group := range engine.groups {
		if strings.HasPrefix(r.URL.Path, "/"+group.prefix) {
			middles = append(middles, group.middles...)
		}
	}
	ctx := newContext(w, r)
	ctx.handlers = middles
	engine.router.handle(ctx)
}
