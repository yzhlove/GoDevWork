package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

type Payload struct {
	Path string `json:"path"`
}

type PayloadCollection struct {
	Version  string    `json:"version"`
	Token    string    `json:"token"`
	Payloads []Payload `json:"data"`
}

func main() {

	data, err := generateData("123456", "1")
	if err != nil {
		panic(err)
	}
	fmt.Println(len(data))
	resp, err := http.Post("http://localhost:1234/payload",
		"application/x-www-form-urlencoded", strings.NewReader(string(data)))
	if err != nil {
		panic(err)
	}
	log.Println("code => ", resp.StatusCode)

}

func generateData(token, version string) ([]byte, error) {
	content := &PayloadCollection{
		Version:  version,
		Token:    token,
		Payloads: make([]Payload, 0, 100),
	}
	for i := 0; i < 100; i++ {
		content.Payloads = append(content.Payloads, Payload{
			Path: strings.Repeat(strconv.Itoa(rand.Intn(1000)+1), 5),
		})
	}
	return json.Marshal(content)
}
