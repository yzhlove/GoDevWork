package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var MaxQueue = 128
var MaxBuffer int64 = 4096
var JobQueue chan Payload

type Payload struct {
	Path string `json:"path"`
}

type PayloadCollection struct {
	Version  string    `json:"version"`
	Token    string    `json:"token"`
	Payloads []Payload `json:"data"`
}

func init() {
	JobQueue = make(chan Payload, MaxQueue)
}

func processor() {
	for pay := range JobQueue {
		pay.Update()
	}
	log.Println("[JobQueue] exit .")
}

func (p *Payload) String() string {
	return fmt.Sprintf("[Payload] path:%s", p.Path)
}

func (p *Payload) Update() {
	time.Sleep(100 * time.Millisecond)
	log.Println(p)
}

func payloadH(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var content = &PayloadCollection{}
	if err := json.NewDecoder(io.LimitReader(r.Body, MaxBuffer)).Decode(&content); err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for _, pay := range content.Payloads {
		JobQueue <- pay
	}
	w.WriteHeader(http.StatusOK)
}

func main() {
	go processor()
	http.HandleFunc("/payload", payloadH)
	log.Println("start server listen by port -> 0.0.0.0:1234")
	if err := http.ListenAndServe(":1234", nil); err != nil {
		panic("[server] start server err:" + err.Error())
	}
}
