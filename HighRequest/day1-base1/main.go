package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

//bad design

type SimS3Bucket struct{}

var S3Bucket = SimS3Bucket{}

func (s *SimS3Bucket) PutReader(_ string, _ *bytes.Buffer, _ int64) error {
	fmt.Println("hello world")
	return nil
}

type Payload struct {
	Path string `json:"path"`
}

type PayloadCollection struct {
	WindowsVersion string    `json:"version"`
	Token          string    `json:"token"`
	Payloads       []Payload `json:"data"`
}

func (p *Payload) UploadTos3() error {
	storagePath := fmt.Sprintf("%v/%v", p.Path, time.Now().UnixNano())
	bucket := S3Bucket
	sb := new(bytes.Buffer)
	if err := json.NewEncoder(sb).Encode(p); err != nil {
		return err
	}
	return bucket.PutReader(storagePath, sb, int64(sb.Len()))
}

func payloadHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var content = &PayloadCollection{}
	if err := json.NewDecoder(io.LimitReader(r.Body, 1<<10)).Decode(&content); err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for _, payload := range content.Payloads {
		go payload.UploadTos3()
	}
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/payload", payloadHandle)
	fmt.Println(http.ListenAndServe(":1234", nil))
}
