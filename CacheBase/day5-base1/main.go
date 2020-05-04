package main

import (
	"flag"
	"fmt"
	"geecachefive"
	"log"
	"net/http"
)

var (
	db = map[string]string{
		"Tom":  "630",
		"Jack": "589",
		"Sam":  "567",
	}

	groupName = "scores"
)

func createGroup() *geecachefive.Group {
	return geecachefive.NewGroup(groupName, 2<<10, geecachefive.GetterFunc(func(key string) ([]byte, error) {
		log.Println("[slow-db] search key ", key)
		if v, ok := db[key]; ok {
			return []byte(v), nil
		}
		return nil, fmt.Errorf("%s not exist ", key)
	}))
}

func startCacheServer(addr string, addrs []string, gee *geecachefive.Group) {
	peers := geecachefive.NewHttpContext(addr)
	peers.Set(addrs...)
	gee.RegisterPeers(peers)
	log.Println("geecache is running at --> ", addr)
	log.Fatal(http.ListenAndServe(addr[7:], peers))
}

func startApiServer(apiAddr string, gee *geecachefive.Group) {
	http.Handle("/api", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		view, err := gee.Get(key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write(view.GetBytes())
	}))
	log.Println("api servet is running at", apiAddr)
	log.Fatal(http.ListenAndServe(apiAddr[7:], nil))
}

var (
	_port int
	_api  bool
)

func init() {
	flag.IntVar(&_port, "port", 8001, "gee cache server port")
	flag.BoolVar(&_api, "api", false, "start a api server")
	flag.Parse()
}

func main() {

	apiAddr := "http://localhost:9999"
	addrMap := map[int]string{
		8001: "http://localhost:8001",
		8002: "http://localhost:8002",
		8003: "http://localhost:8003",
	}

	var addrs []string
	for _, v := range addrMap {
		addrs = append(addrs, v)
	}

	gee := createGroup()
	if _api {
		go startApiServer(apiAddr, gee)
	}
	startCacheServer(addrMap[_port], addrs, gee)

}
