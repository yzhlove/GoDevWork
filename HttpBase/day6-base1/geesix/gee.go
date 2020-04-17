package geesix

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
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
	router        *router
	groups        []*RouterGroup
	htmlTemplates *template.Template
	funcMap       template.FuncMap
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

//静态路由（文件服务器）
func (group *RouterGroup) createStaticHandle(filepath string, fs http.FileSystem) HandleFunc {
	fpath := path.Join(group.prefix, filepath)
	fileServer := http.StripPrefix(fpath, http.FileServer(fs))
	return func(ctx *Context) {
		//if _, err := fs.Open(file); err != nil {
		//	ctx.SetCode(http.StatusNotFound)
		//	return
		//}
		fmt.Println("file path ==> ", ctx.Params)
		fileServer.ServeHTTP(ctx.Resp, ctx.Req)
	}
}

func (group *RouterGroup) Static(filepath string, root string) {
	handler := group.createStaticHandle(filepath, http.Dir(root))
	pattern := path.Join(filepath, "/*filepath")
	group.GET(pattern, handler)
}

func (engine *Engine) SetFuncMap(funcMap template.FuncMap) {
	engine.funcMap = funcMap
}

func (engine *Engine) LoadHTMLGlob(pattern string) {
	engine.htmlTemplates = template.Must(template.New("").Funcs(engine.funcMap).ParseGlob(pattern))
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
	ctx.engine = engine
	engine.router.handle(ctx)
}
