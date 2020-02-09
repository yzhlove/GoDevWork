package tcp

import (
	"bufio"
	"errors"
	"io"
)

func (s *Server) readKey(reader *bufio.Reader) (value string, err error) {
	size, err := readSize(reader)
	if err != nil {
		return
	}
	buf := make([]byte, size)
	if _, err = io.ReadFull(reader, buf); err != nil {
		return
	}
	value = string(buf)
	addr, ok := s.ShouldProcess(value)
	if !ok {
		err = errors.New("redirect " + addr)
		return
	}
	return
}

func (s *Server) readKeyAndValue(reader *bufio.Reader) (key string, value []byte, err error) {
	var keySize, valueSize int
	if keySize, err = readSize(reader); err != nil {
		return
	}
	if valueSize, err = readSize(reader); err != nil {
		return
	}
	buf := make([]byte, keySize)
	if _, err = io.ReadFull(reader, buf); err != nil {
		return
	}
	key = string(buf)
	addr, ok := s.ShouldProcess(key)
	if !ok {
		err = errors.New("redirect " + addr)
		return
	}
	value = make([]byte, valueSize)
	_, err = io.ReadFull(reader, value)
	return
}
