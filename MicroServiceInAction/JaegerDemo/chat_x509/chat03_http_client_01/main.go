package main

import (
	"crypto/tls"
	"io"
	"net/http"
	"os"
)

//不验证的客户端

func main() {

	c := http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}

	if resp, err := c.Get("https://localhost:1234/hello"); err != nil {
		panic(err)
	} else {
		defer resp.Body.Close()
		io.Copy(os.Stdout, resp.Body)
	}
}
