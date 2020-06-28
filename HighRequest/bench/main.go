package main

import (
	"encoding/json"
	"fmt"
	"go.uber.org/atomic"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

const url = "http://localhost:1234/payload"
const contentType = "application/x-www-form-urlencoded"

func main() {

	var count atomic.Uint32
	var succeed atomic.Uint32
	var wg sync.WaitGroup

	req, err := requestData()
	if err != nil {
		panic(err)
	}

	start := time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		fmt.Println("current => ", i+1)
		go httpRequest(string(req), &count, &succeed, &wg)
	}
	wg.Wait()

	fmt.Println("==========================")
	fmt.Println("| count:", count.Load())
	fmt.Println("| succeed:", succeed.Load())
	fmt.Println("| time:", time.Now().Sub(start))
	fmt.Println("==========================")

}

type Payload struct {
	Path string `json:"path"`
}

type PayloadCollection struct {
	Version  string    `json:"version"`
	Token    string    `json:"token"`
	Payloads []Payload `json:"data"`
}

func httpRequest(req string, count, succeed *atomic.Uint32, wg *sync.WaitGroup) {
	defer func() {
		count.Add(1)
		wg.Done()
	}()
	c := http.Client{Timeout: 5 * time.Second}
	if resp, err := c.Post(url, contentType, strings.NewReader(req)); err != nil {
		return
	} else {
		if resp.StatusCode == http.StatusOK {
			succeed.Add(1)
		}
	}
}

func requestData() ([]byte, error) {
	content := &PayloadCollection{
		Version:  "1",
		Token:    "*#06#*#*",
		Payloads: make([]Payload, 0, 100),
	}
	for i := 0; i < 100; i++ {
		content.Payloads = append(content.Payloads, Payload{
			Path: strings.Repeat(strconv.Itoa(rand.Intn(1000)+1), 5),
		})
	}
	return json.Marshal(content)
}
