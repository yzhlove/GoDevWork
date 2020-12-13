package main

import "encoding/json"

type Sender interface {
	Send() []byte
}

type TextData struct {
	Content string `json:"content"`
}

type MarkData struct {
	Tag     string `json:"title"`
	Content string `json:"text"`
}

type MarkDoc struct {
	MsgType  string `json:"msgtype"`
	MarkData `json:"markdown"`
}

type TextDoc struct {
	MsgType  string `json:"msgtype"`
	TextData `json:"text"`
}

func (doc *MarkDoc) Send() []byte {
	return trans(doc)
}

func (doc *TextDoc) Send() []byte {
	return trans(doc)
}

func trans(sender Sender) []byte {
	data, err := json.Marshal(sender)
	if err != nil {
		panic(err)
	}
	return data
}
