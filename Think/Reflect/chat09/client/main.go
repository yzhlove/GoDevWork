package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/rpc"
	"time"
)

type Req struct {
	S string
	A int
	T time.Time
}

type Ardith struct{}

func (t *Ardith) Handle(req Req, resp *[]string) error {
	fmt.Print("req ==> ", req)
	data, err := json.Marshal(req)
	*resp = append(*resp, req.S)
	*resp = append(*resp, string(data))
	return err
}

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dail error:", err)
	}

	req := &Req{S: "hello", A: 128, T: time.Now()}
	var resp = make([]string, 0, 0)

	if err := client.Call("Ardith.Handle", req, &resp); err != nil {
		log.Fatal("call error:", err)
	}
	fmt.Print("resp => ", resp)

}
