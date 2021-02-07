package codec

import (
	"io"
	"log"
)

type Header struct {
	Method string
	Seq    uint64
	Error  string
}

type Codec interface {
	io.Closer
	ReadHeader(header *Header) error
	ReadBody(body interface{}) error
	Write(header *Header, body interface{}) error
}

type Type string

const (
	GobType  Type = "application/gob"
	JsonType Type = "application/json"
)

var funcMaps map[Type]func(closer io.ReadWriteCloser) Codec

func init() {
	funcMaps = make(map[Type]func(closer io.ReadWriteCloser) Codec)
}

func registry(t Type, fn func(closer io.ReadWriteCloser) Codec) {
	if funcMaps != nil {
		funcMaps[t] = fn
	} else {
		log.Println("registry error: funcMaps no init")
	}
}
