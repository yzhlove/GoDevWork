package geeseven

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"
)

type HandleFunc func(ctx *Context)

type RouterGroup struct {
	prefix     string
	components []HandleFunc
	parent     *RouterGroup
	engine     *Engine
}

type Engine struct {
	*RouterGroup
	router       *router
	groups       []*RouterGroup
	htmlTemplate *template.Template
	htmlFuncMap  template.FuncMap
}

func NewEngine() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = append(engine.groups, engine.RouterGroup)
	return engine
}

func NewDefaultEngine() *Engine {
	engine := NewEngine()
	engine.Use(Logger(), Recovery())
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

func (group *RouterGroup) Use(components ...HandleFunc) {
	group.components = append(group.components, components...)
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

func (group *RouterGroup) staticHandle(prefix string, system http.FileSystem) HandleFunc {
	filter := path.Join(group.prefix, prefix)
	server := http.StripPrefix(filter, http.FileServer(system))
	return func(ctx *Context) {
		file := ctx.GetParamValue("filepath")
		if _, err := system.Open(file); err != nil {
			ctx.SetCode(http.StatusNotFound)
			log.Printf("%s not found:%s \n", file, err.Error())
			return
		}
		server.ServeHTTP(ctx.Resp, ctx.Req)
	}
}

func (group *RouterGroup) StaticRouter(prefix, root string) {
	pattern := path.Join(prefix, "/*filepath")
	group.GET(pattern, group.staticHandle(prefix, http.Dir(root)))
}

func (engine *Engine) SetFuncMap(funcMap template.FuncMap) {
	engine.htmlFuncMap = funcMap
}

func (engine *Engine) LoadHtmlGlob(pattern string) {
	engine.htmlTemplate = template.Must(template.New("").Funcs(engine.htmlFuncMap).ParseGlob(pattern))
}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var comps []HandleFunc
	for _, group := range engine.groups {
		if strings.HasPrefix(r.URL.Path, "/"+group.prefix) {
			comps = append(comps, group.components...)
		}
	}
	ctx := newContext(w, r)
	ctx.handles = comps
	ctx.engine = engine
	engine.router.handle(ctx)
}
