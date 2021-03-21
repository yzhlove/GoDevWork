package main

import (
	"bufio"
	"fmt"
	"net/http"
	"testing"
)

func Test_HttpGet(t *testing.T) {

	request, err := http.Get("http://127.0.0.1:1234/test")
	if err != nil {
		panic(err)
	}
	defer request.Body.Close()

	scan := bufio.NewScanner(request.Body)
	for scan.Scan() {
		fmt.Printf("resp:(%q)", scan.Text())
	}

}