package geecacheseven

import (
	"fmt"
	"geecacheseven/consistent"
	"log"
	"net/http"
	"strings"
	"sync"
)

const (
	basePath = "/_geecache/"
	replicas = 50
)

type HttpContext struct {
	self        string
	basePath    string
	mutex       sync.Mutex
	peers       *consistent.Map
	httpGetters map[string]*httpGetter
}

type httpGetter struct {
	baseURL string
}

func NewHttpContext(self string) *HttpContext {
	return &HttpContext{
		self:     self,
		basePath: basePath,
	}
}

func (h *HttpContext) Log(format string, a ...interface{}) {
	log.Printf("[server %s] %s", h.self, fmt.Sprintf(format, a...))
}

func (h *HttpContext) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pattern := r.URL.Path
	if !strings.HasPrefix(pattern, h.basePath) {
		panic("HttpContext serving unexpected path:" + pattern)
	}
	h.Log("%s %s ", r.Method, pattern)
	// request=> /<basepath>/<groupname>/<key>
	parts := strings.SplitN(pattern[len(h.basePath):], "/", 2)
	if len(parts) != 2 {
		http.Error(w, "bad request ", http.StatusBadRequest)
		return
	}

}
