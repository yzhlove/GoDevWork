package codec

import "io"

type Header struct {
	Method string
	Seq    uint64
	Err    string
}

type Coder interface {
	io.Closer
	ReadHeader(header *Header) error
	ReadBody(body interface{}) error
	Send(head *Header, body interface{}) error
}

type ParseFunc func(conn io.ReadWriteCloser) Coder

type CodecsType string

const (
	Gob  CodecsType = "application/gob"
	Json CodecsType = "application/json"
)

var CodecsMap map[CodecsType]ParseFunc

func init() {
	CodecsMap = make(map[CodecsType]ParseFunc, 2)
	CodecsMap[Gob] = NewGobCodec
	CodecsMap[Json] = NewJsonCodec
}
