package codec

import (
	"bufio"
	"encoding/gob"
	"io"
	"log"
)

type GOBCodec struct {
	conn    io.ReadWriteCloser
	buf     *bufio.Writer
	decoder *gob.Decoder
	encoder *gob.Encoder
}

func (c *GOBCodec) Close() error {
	return c.conn.Close()
}

func (c *GOBCodec) ReadHeader(header *Header) error {
	return c.decoder.Decode(header)
}

func (c *GOBCodec) ReadBody(body interface{}) error {
	return c.decoder.Decode(body)
}

func (c *GOBCodec) Writer(header *Header, body interface{}) (err error) {

	defer func() {
		c.buf.Flush()
		if err != nil {
			c.Close()
		}
	}()

	if err = c.encoder.Encode(header); err != nil {
		log.Println("rpc:encoder header error->", err)
		return
	}

	if err = c.encoder.Encode(body); err != nil {
		log.Println("rpc:encoder body error->", err)
		return
	}

	return
}

func NewGOBCodec(conn io.ReadWriteCloser) Codec {
	buf := bufio.NewWriter(conn)
	return &GOBCodec{conn: conn, buf: buf, encoder: gob.NewEncoder(buf), decoder: gob.NewDecoder(conn)}
}
