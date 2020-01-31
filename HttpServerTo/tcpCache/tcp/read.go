package tcp

import (
	"bufio"
	"io"
)

func (s *Server) readKey(r *bufio.Reader) (str string, err error) {
	length, err := readLen(r)
	if err != nil {
		return
	}
	key := make([]byte, length)
	if _, err = io.ReadFull(r, key); err != nil {
		return
	}
	str = string(key)
	return
}

func (s *Server) readKeyAndValue(r *bufio.Reader) (key string, value []byte, err error) {
	keyLen, err := readLen(r)
	if err != nil {
		return
	}
	valueLen, err := readLen(r)
	if err != nil {
		return
	}
	keyBytes := make([]byte, keyLen)
	if _, err = io.ReadFull(r, keyBytes); err != nil {
		return
	}
	value = make([]byte, valueLen)
	if _, err = io.ReadFull(r, value); err != nil {
		return
	}
	key = string(keyBytes)
	return
}
