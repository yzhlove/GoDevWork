package geethree

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

const defaultPath = "/_geecache/"

type HttpContext struct {
	self     string
	basePath string
}

func NewHttpContext(s string) *HttpContext {
	return &HttpContext{
		self:     s,
		basePath: defaultPath,
	}
}

func (c *HttpContext) Log(format string, v ...interface{}) {
	log.Printf("[Server %s] %s", c.self, fmt.Sprintf(format, v))
}

func (c *HttpContext) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, c.basePath) {
		http.Error(w, fmt.Sprintf("path err:%v", r.URL.Path), http.StatusBadRequest)
		return
	}
	c.Log("%s %s", r.Method, r.URL.Path)
	parts := strings.SplitN(r.URL.Path[len(c.basePath):], "/", 2)
	if len(parts) != 2 {
		http.Error(w, "bad request ", http.StatusBadRequest)
		return
	}
	name := parts[0]
	key := parts[1]
	fmt.Printf("%q %q \n", name, key)
	group := GetGroup(name)
	if group == nil {
		http.Error(w, "no such group :"+name, http.StatusBadRequest)
		return
	}
	if view, err := group.Get(key); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write(view.GetBytes())
	}

}
