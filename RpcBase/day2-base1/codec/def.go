package codec

import "io"

type Header struct {
	SvcMethod string
	Seq       uint64
	Err       string
}

type Codec interface {
	Close() error
	ReadHeader(header *Header) error
	ReadBody(body interface{}) error
	Writer(header *Header, body interface{}) error
}

type Func func(conn io.ReadWriteCloser) Codec

type Type string

const (
	GOB  Type = "application/gob"
	JSON Type = "application/json"
)

var CodecsMap map[Type]Func

func init() {
	CodecsMap = make(map[Type]Func)
	CodecsMap[GOB] = NewGOBCodec
}
