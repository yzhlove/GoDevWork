package codec

import (
	"bufio"
	"encoding/gob"
	"io"
	"log"
)

type GobCodec struct {
	conn    io.ReadWriteCloser
	buf     *bufio.Writer
	encoder *gob.Encoder
	decoder *gob.Decoder
}

func init() {
	registry(GobType, ConstructGobCodec)
}

func ConstructGobCodec(conn io.ReadWriteCloser) Codec {
	buf := bufio.NewWriter(conn)
	return &GobCodec{
		conn:    conn,
		buf:     buf,
		encoder: gob.NewEncoder(buf),
		decoder: gob.NewDecoder(conn),
	}
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

func (c *GobCodec) Write(header *Header, body interface{}) (err error) {

	defer func() {
		c.buf.Flush()
		if err != nil {
			_ = c.Close()
		}
	}()

	if err = c.encoder.Encode(header); err != nil {
		log.Println("rpc server: encoder header error: ", err)
		return
	}
	if err = c.encoder.Encode(body); err != nil {
		log.Println("rpc server: encoder body error: ", err)
		return
	}
	return
}
