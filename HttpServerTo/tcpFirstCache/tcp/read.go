package tcp

import (
	"bufio"
	"io"
)

func (s *Server) readKey(buf *bufio.Reader) (key string, err error) {
	var size int
	if size, err = readSize(buf); err != nil {
		return
	} else {
		bs := make([]byte, size)
		if _, err = io.ReadFull(buf, bs); err != nil {
			return
		}
		key = string(bs)
	}
	return
}

func (s *Server) readKeyAndValue(buf *bufio.Reader) (key string, value []byte, err error) {
	var keySize, valueSize int
	if keySize, err = readSize(buf); err != nil {
		return
	}
	if valueSize, err = readSize(buf); err != nil {
		return
	}
	_read := func(size int) (bs []byte, err error) {
		bs = make([]byte, size)
		_, err = io.ReadFull(buf, bs)
		return
	}
	var keyBytes []byte
	if keyBytes, err = _read(keySize); err != nil {
		return
	}
	key = string(keyBytes)
	value, err = _read(valueSize)
	return
}
