package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type SimS3Bucket struct{}

var S3Bucket = &SimS3Bucket{}

type Payload struct {
	Path string `json:"path"`
}

type PayloadCollection struct {
	WindowsVersion string    `json:"version"`
	Token          string    `json:"token"`
	Payloads       []Payload `json:"data"`
}

func (s *SimS3Bucket) PutReader(path string, _ io.ReadWriter, _ int64) error {
	fmt.Println("path:", path)
	return nil
}

func (p *Payload) UploadS3() error {
	storageLocation := fmt.Sprintf("%s/%v", p.Path, time.Now().UnixNano())
	bucket := S3Bucket
	sb := new(bytes.Buffer)
	if err := json.NewEncoder(sb).Encode(p); err != nil {
		return err
	}
	return bucket.PutReader(storageLocation, sb, int64(sb.Len()))
}
