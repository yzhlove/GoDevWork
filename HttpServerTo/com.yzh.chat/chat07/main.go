package main

import "net/http"

func main() {

	_, err := http.Get("http://dev.gmgate.net:8080")
	if err != nil {
		panic(err)
	}

}
