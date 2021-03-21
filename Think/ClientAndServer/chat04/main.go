package main

import (
	"fmt"
	"net/http"
)

//get file data size

func main() {

	head, err := http.Head("http://127.0.0.1:1234/assets/web.html")
	if err != nil {
		panic(err)
	}

	fmt.Println("size => ", head.ContentLength," ",float64(head.ContentLength)/1000/1000,"MB")

}
