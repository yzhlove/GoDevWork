package main

import (
	"fmt"
	"geethree"
	"log"
	"net/http"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func main() {

	geethree.NewGroup("scores", 2<<10, geethree.GetterFunc(func(key string) ([]byte, error) {
		log.Println("[SlowDB] search key ", key)
		if v, ok := db[key]; ok {
			return []byte(v), nil
		}
		return nil, fmt.Errorf("%s not exist", key)
	}))

	addr := "localhost:1234"
	peers := geethree.NewHttpContext(addr)
	log.Println("geecache running ... to ", addr)
	log.Fatal(http.ListenAndServe(addr, peers))

}
