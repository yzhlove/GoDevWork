package main

import "net/http"

func main() {

	_ = http.ListenAndServe(":1234", nil)

}



