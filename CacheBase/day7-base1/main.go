package main

import (
	"flag"
	"fmt"
	"geecacheseven"
	"log"
	"net/http"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

var (
	port int
	api  bool
)

func init() {
	flag.IntVar(&port, "port", 8001, "cache server port")
	flag.BoolVar(&api, "api", false, "start a api server")
	flag.Parse()
}

func main() {

	apiAddr := "http://localhost:9999"
	points := map[int]string{
		8001: "http://localhost:8001",
		8002: "http://localhost:8002",
		8003: "http://localhost:8003",
	}

	var addrs []string
	for _, addr := range points {
		addrs = append(addrs, addr)
	}

	gee := createGroup()
	if api {
		go startAPIServer(apiAddr, gee)
	}
	startCacheServer(points[port], addrs, gee)
}

func createGroup() *geecacheseven.Group {
	return geecacheseven.NewGroup("scores", 2<<10, geecacheseven.GetterFunc(func(key string) ([]byte, error) {
		log.Println("[access sql ] search key ", key)
		if v, ok := db[key]; ok {
			return []byte(v), nil
		}
		return nil, fmt.Errorf("%s not exists ", key)
	}))
}

func startCacheServer(addr string, points []string, gee *geecacheseven.Group) {
	httpContext := geecacheseven.NewHttpContext(addr)
	httpContext.Set(points...)
	gee.RegisterPeers(httpContext)
	log.Println("geecache is running at :", addr)
	log.Fatal(http.ListenAndServe(addr[len("http://"):], httpContext))
}

func startAPIServer(apiAddr string, gee *geecacheseven.Group) {
	http.Handle("/api", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		key := request.URL.Query().Get("key")
		view, err := gee.Get(key)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Println("api get key:", key, " result:", view.String())
		writer.Header().Set("Content-Type", "text/plain")
		writer.Write(view.Bytes())
	}))
	log.Println("start api server is running at ", apiAddr)
	log.Fatal(http.ListenAndServe(apiAddr[len("http://"):], nil))
}
