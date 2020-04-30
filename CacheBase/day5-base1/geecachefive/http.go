package geecachefive

import (
	"errors"
	"fmt"
	"geecachefive/consistenthash"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

const (
	defaultBasePath = "/_geecache/"
	defaultReplicas = 50
)

type httpGetter struct {
	baseURL string
}

type HttpContext struct {
	self        string
	basePath    string
	mutex       sync.Mutex
	peers       *consistenthash.Map
	httpGetters map[string]*httpGetter
}

func NewHttpContext(s string) *HttpContext {
	return &HttpContext{
		self:     s,
		basePath: defaultBasePath,
	}
}

func (c *HttpContext) Log(format string, v ...interface{}) {
	log.Printf("[server %s ] %s", c.self, fmt.Sprintf(format, v...))
}

func (c *HttpContext) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, defaultBasePath) {
		panic(errors.New("http context server unexpected path:" + r.URL.Path))
	}
	c.Log("%s %s ", r.Method, r.URL.Path)
	parts := strings.SplitN(r.URL.Path[len(c.basePath):], "/", 2)
	if len(parts) != 2 {
		http.Error(w, "bad request ", http.StatusBadRequest)
		return
	}
	group := GetGroup(parts[0])
	if group == nil {
		http.Error(w, "no such group:"+parts[0], http.StatusNotFound)
		return
	}

	view, err := group.Get(parts[1])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(view.GetBytes())
}

func (c *HttpContext) Set(peers ...string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.peers = consistenthash.NewConsistentHash(defaultReplicas, nil)
	c.peers.Set(peers...)
	c.httpGetters = make(map[string]*httpGetter, len(peers))
	for _, peer := range peers {
		c.httpGetters[peer] = &httpGetter{baseURL: peer + c.basePath}
	}
}

func (c *HttpContext) PickPeer(key string) (PeerGetter, bool) {

	c.mutex.Lock()
	defer c.mutex.Unlock()

	if peer := c.peers.Get(key); peer != "" && peer != c.self {
		c.Log("pick peer %s", peer)
		return c.httpGetters[peer], true
	}
	return nil, false
}

func (h *httpGetter) Get(group, key string) ([]byte, error) {
	urlpath := fmt.Sprintf("%v%v/%v", h.baseURL, url.QueryEscape(group), url.QueryEscape(key))
	res, err := http.Get(urlpath)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned: %v ", res.Status)
	}
	return ioutil.ReadAll(res.Body)
}
