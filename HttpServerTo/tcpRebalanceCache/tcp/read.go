package tcp

import (
	"bufio"
	"errors"
	"io"
)

func (s *Server) readKey(reader *bufio.Reader) (key string, err error) {
	size, err := getSize(reader)
	if err != nil {
		return
	}
	buf := make([]byte, size)
	if _, err = io.ReadFull(reader, buf); err != nil {
		return
	}
	key = string(buf)
	address, ok := s.ShouldProcess(key)
	if !ok {
		err = errors.New("redirect " + address)
	}
	return
}

func (s *Server) readKeyAndValue(reader *bufio.Reader) (key string, value []byte, err error) {
	var keySize, valueSize int
	if keySize, err = getSize(reader); err != nil {
		return
	}
	if valueSize, err = getSize(reader); err != nil {
		return
	}
	buf := make([]byte, keySize)
	if _, err = io.ReadFull(reader, buf); err != nil {
		return
	}
	key = string(buf)
	address, ok := s.ShouldProcess(key)
	if !ok {
		err = errors.New("redirect " + address)
		return
	}
	value = make([]byte, valueSize)
	_, err = io.ReadFull(reader, value)
	return
}
