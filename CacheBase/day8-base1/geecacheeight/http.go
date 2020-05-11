package geecacheeight

import (
	"fmt"
	"geecacheeight/consistent"
	"geecacheeight/pb"
	"github.com/gogo/protobuf/proto"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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

	group := GetGroup(parts[0])
	if group == nil {
		http.Error(w, "no such group :"+parts[0], http.StatusNotFound)
		return
	}

	view, err := group.Get(parts[1])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("http cache get key ===> ", view.String())
	body, err := proto.Marshal(&pb.Cache_Resp{Value: view.Bytes()})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(body)

}

func (h *HttpContext) Set(peers ...string) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	h.peers = consistent.NewConsistent(replicas, nil)
	h.peers.Set(peers...)
	h.httpGetters = make(map[string]*httpGetter, len(peers))
	for _, peer := range peers {
		// url -> http://localhost:port/_geecache/
		h.httpGetters[peer] = &httpGetter{baseURL: peer + h.basePath}
	}
}

func (h *HttpContext) PickPeer(key string) (PeerGetter, bool) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	if peer := h.peers.Get(key); peer != "" && peer != h.self {
		h.Log("Pick peer %s", peer)
		return h.httpGetters[peer], true
	}
	return nil, false
}

func (hp *httpGetter) Get(in *pb.Cache_Req, out *pb.Cache_Resp) error {
	u := fmt.Sprintf("%v%v/%v", hp.baseURL,
		url.QueryEscape(in.Group),
		url.QueryEscape(in.Key))

	result, err := http.Get(u)
	if err != nil {
		return err
	}

	defer result.Body.Close()
	if result.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned:%v ", result.Status)
	}
	bytes, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return fmt.Errorf("reading response body: %v", err)
	}
	if err := proto.Unmarshal(bytes, out); err != nil {
		return fmt.Errorf("decoding response body: %v ", err)
	}
	return nil
}
