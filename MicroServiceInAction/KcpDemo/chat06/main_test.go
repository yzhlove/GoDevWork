package main

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func Test_Client(t *testing.T) {
	resp, err := http.Get("http://localhost:1234/hello")
	if err != nil {
		t.Error(err)
		return
	}
	defer resp.Body.Close()
	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("message :", string(d))
}
