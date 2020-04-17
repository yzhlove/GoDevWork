package main

import (
	"log"
	"net/http"
)

////////////////////////////
// HttpFileServer
////////////////////////////

func main() {
	path := "/Users/yurisa/Develop/GoWork/src/WorkSpace/GoDevWork/HttpBase/day6-base1/static/"
	fs := http.StripPrefix("/assets", http.FileServer(http.Dir(path)))
	log.Println(http.ListenAndServe(":1234", fs))

}
