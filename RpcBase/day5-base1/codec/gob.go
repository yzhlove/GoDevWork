package codec

import (
	"bufio"
	"encoding/gob"
	"io"
)

type GobCodec struct {
	conn    io.ReadWriteCloser
	buffer  *bufio.Writer
	encoder *gob.Encoder
	decoder *gob.Decoder
}

func NewGobCodec(conn io.ReadWriteCloser) Coder {
	buffer := bufio.NewWriter(conn)
	return &GobCodec{conn: conn, buffer: buffer, encoder: gob.NewEncoder(buffer), decoder: gob.NewDecoder(conn)}
}

func (c *GobCodec) Close() error {
	return c.conn.Close()
}

func (c *GobCodec) ReadHeader(header *Header) error {
	return c.decoder.Decode(header)
}

func (c *GobCodec) ReadBody(body interface{}) error {
	return c.decoder.Decode(body)
}

func (c *GobCodec) Send(header *Header, body interface{}) (err error) {

	defer func() {
		c.buffer.Flush()
		if err != nil {
			c.Close()
		}
	}()

	if err = c.encoder.Encode(header); err != nil {
		return
	}

	return c.encoder.Encode(body)
}


