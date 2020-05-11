package main

import (
	"flag"
	"fmt"
	"geecacheeight"
	"net/http"

	"log"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
	"yzh":  "19960201",
	"lcm":  "19960222",
	"wyq":  "19961234",
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
		8001: ":8001",
		8002: ":8002",
		8003: ":8003",
	}

	address := make([]string, 0, len(points))
	for _, addr := range points {
		address = append(address, addr)
	}

	group := createGroup()
	if api {
		go startAPIServer(apiAddr, group)
	}
	startRpcServer(points[port], address, group)
}

func createGroup() *geecacheeight.Group {
	return geecacheeight.NewGroup("scores", 2<<10, geecacheeight.GetterFunc(func(key string) ([]byte, error) {
		log.Println("[access sql ] search key ", key)
		if v, ok := db[key]; ok {
			return []byte(v), nil
		}
		return nil, fmt.Errorf("%s not exists ", key)
	}))
}

func startRpcServer(addr string, points []string, gee *geecacheeight.Group) {
	rpcContext := geecacheeight.NewRpcContext(addr)
	rpcContext.Set(points...)
	gee.RegisterPeers(rpcContext)
	log.Fatal(rpcContext.ServeRPC())
}

func startAPIServer(apiAddr string, gee *geecacheeight.Group) {
	http.Handle("/api", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		key := request.URL.Query().Get("key")
		view, err := gee.Get(key)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-Type", "application/octet-stream")
		writer.Write(view.Bytes())
	}))
	log.Println("[start api server] is running at ", apiAddr)
	log.Fatal(http.ListenAndServe(apiAddr[len("http://"):], nil))
}
