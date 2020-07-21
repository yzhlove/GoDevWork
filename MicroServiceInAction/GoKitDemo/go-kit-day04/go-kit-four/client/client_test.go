package client

import (
	"encoding/json"
	"go.uber.org/atomic"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"testing"
	"time"
)

type Login struct {
	Account string `json:"account"`
	Passwd  string `json:"passwd"`
}

func Test_Login(t *testing.T) {

	data, err := json.Marshal(Login{Account: "yzh", Passwd: "123456789"})
	if err != nil {
		t.Error(err)
		return
	}
	var wg sync.WaitGroup
	var succeed atomic.Int32
	var count atomic.Int32
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				count.Add(1)
				wg.Done()
			}()

			resp, err := http.Post("http://localhost:1234/login",
				"application/json", strings.NewReader(string(data)))
			if err != nil {
				t.Error(err)
				return
			}
			defer resp.Body.Close()
			content, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Error(err)
				return
			}
			t.Log("result => ", string(content))
			succeed.Add(1)
		}()
	}
	wg.Wait()
	t.Log("[100 request] count:", count.Load(), " succeed:", succeed.Load(), " failed:", count.Sub(succeed.Load()))
}

func Test_Add(t *testing.T) {

	url := "http://localhost:1234/sum?a=100&b=80"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJOYW1lIjoieXpoIiwiRGNJZCI6MSwiZXhwIjoxNTk1MzI2ODIxLCJpYXQiOjE1OTUzMjY3OTEsImlzcyI6ImdvLWtpdCIsIm5iZiI6MTU5NTMyNjc5MSwic3ViIjoibG9naW4ifQ.xt8PmQkgfkoie76ZmIP0h6rL3w-WAkSrOZG9Zg5ZOMA"
	var wg sync.WaitGroup
	var succeed atomic.Int32
	var count atomic.Int32
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				count.Add(1)
				wg.Done()
			}()
			c := http.Client{Timeout: time.Second * 5}
			request, err := http.NewRequest(http.MethodGet, url, nil)
			if err != nil {
				t.Error(err)
				return
			}
			request.Header.Set("X-Token", token)
			response, err := c.Do(request)
			if err != nil {
				t.Error(err)
				return
			}
			defer response.Body.Close()
			result, err := ioutil.ReadAll(response.Body)
			if err != nil {
				t.Error(err)
				return
			}
			t.Log("result => ", string(result))
			succeed.Add(1)
		}()
	}
	wg.Wait()
	t.Log("[100 request]count:", count.Load(), "succeed:", succeed.Load(), "failed:", count.Sub(succeed.Load()))
}
