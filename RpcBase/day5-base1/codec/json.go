package codec

import (
	"bufio"
	"encoding/json"
	"io"
)

type JsonCodec struct {
	conn    io.ReadWriteCloser
	buffer  *bufio.Writer
	decoder *json.Decoder
	encode  *json.Encoder
}

func NewJsonCodec(conn io.ReadWriteCloser) Coder {
	buffer := bufio.NewWriter(conn)
	return &JsonCodec{conn: conn, buffer: buffer, encode: json.NewEncoder(buffer), decoder: json.NewDecoder(conn)}
}

func (c *JsonCodec) Close() error {
	return c.conn.Close()
}

func (c *JsonCodec) ReadHeader(header *Header) error {
	return c.decoder.Decode(header)
}

func (c *JsonCodec) ReadBody(body interface{}) error {
	return c.decoder.Decode(body)
}

func (c *JsonCodec) Send(header *Header, body interface{}) (err error) {

	defer func() {
		c.buffer.Flush()
		if err != nil {
			c.Close()
		}
	}()

	if err = c.encode.Encode(header); err != nil {
		return
	}

	return c.encode.Encode(body)
}
