package main

import (
	"fmt"
	"gee"
	"log"
	"net/http"
)

func main() {

	r := gee.New()
	r.GET("/", indexHandle)
	r.POST("/hello", helloHandle)
	log.Fatal(r.Run(":1234"))

}

func indexHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.path=%q\n", r.URL.Path)
}

func helloHandle(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Fprintf(w, "Head[%q]=%q\n", k, v)
	}
}
