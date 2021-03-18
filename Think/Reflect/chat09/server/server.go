package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
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

	ardith := &Ardith{}
	rpc.Register(ardith)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listener error:", err)
	}
	log.Println(http.Serve(l, nil))
}
