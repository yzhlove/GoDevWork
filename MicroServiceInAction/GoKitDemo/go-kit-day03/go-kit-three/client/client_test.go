package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

type Login struct {
	Account string `json:"account"`
	Passwd  string `json:"passwd"`
}

func Test_Add(t *testing.T) {
	resp, err := http.Get("httP://localhost:1234/sum?a=100&b=20")
	if err != nil {
		t.Error(err)
		return
	}
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("result => ", string(result))
}

func Test_Login(t *testing.T) {

	data, err := json.Marshal(Login{Account: "yzh", Passwd: "123456789"})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("data => ", string(data))

	resp, err := http.Post("http://localhost:1234/login",
		"appliaction/json", strings.NewReader(string(data)))
	if err != nil {
		t.Error(err)
		return
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("Header => ", resp.Header)
	t.Log("Body => ", string(content))

}

func Test_HeadAdd(t *testing.T) {

	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, "http://localhost:1234/sum?a=100&b=50", nil)
	if err != nil {
		t.Error(err)
		return
	}
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJOYW1lIjoieXpoIiwiRGNJZCI6MSwiZXhwIjoxNTk0OTc3ODgwLCJpYXQiOjE1OTQ5Nzc4NTAsImlzcyI6ImdvLWtpdCIsIm5iZiI6MTU5NDk3Nzg1MCwic3ViIjoibG9naW4ifQ.icog1k9H_DrdSNMU9MyC4JcCUBrw0lvMCGIUS9yJ__s"
	request.Header.Add("X-Token", token)
	resp, err := client.Do(request)
	if err != nil {
		t.Error(err)
		return
	}
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("result => ", string(result))
}
