package main

import (
	"log"
	"net/http"
)

////////////////////////////
// HttpFileServer
////////////////////////////

func main() {
	path := "."
	fs := http.StripPrefix("/assets", http.FileServer(http.Dir(path)))
	log.Println(http.ListenAndServe(":1234", fs))

}
