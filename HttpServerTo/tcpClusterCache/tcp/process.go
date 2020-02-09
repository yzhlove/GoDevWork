package tcp

import (
	"bufio"
	"errors"
	"io"
	"log"
	"net"
)

type result struct {
	buf []byte
	err error
}

func getResultChan(resultChan chan chan *result) chan *result {
	result := make(chan *result)
	resultChan <- result
	return result
}

func (s *Server) get(resultChan chan chan *result, reader *bufio.Reader) {
	c := getResultChan(resultChan)
	key, err := s.readKey(reader)
	if err != nil {
		c <- &result{nil, err}
		return
	}
	go func() {
		if value, err := s.Get(key); err != nil {
			return
		} else {
			c <- &result{value, nil}
		}
	}()
}

func (s *Server) set(resultChan chan chan *result, reader *bufio.Reader) {
	c := getResultChan(resultChan)
	key, value, err := s.readKeyAndValue(reader)
	if err != nil {
		c <- &result{nil, err}
		return
	}
	go func() {
		c <- &result{nil, s.Set(key, value)}
	}()
}

func (s *Server) del(resultChan chan chan *result, reader *bufio.Reader) {
	c := getResultChan(resultChan)
	key, err := s.readKey(reader)
	if err != nil {
		c <- &result{nil, err}
		return
	}
	go func() {
		c <- &result{nil, s.Del(key)}
	}()
}

func replay(conn net.Conn, resultChan chan chan *result) {
	defer conn.Close()
	for {
		if ch, ok := <-resultChan; ok {
			if result, ok := <-ch; ok {
				if err := sendResponse(result.buf, result.err, conn); err != nil {
					log.Println("sendResponse err: ", err.Error())
				}
			}
		} else {
			return
		}
	}
}

func (s *Server) process(conn net.Conn) {
	reader := bufio.NewReader(conn)
	resultChan := make(chan chan *result, 512)
	defer close(resultChan)
	go replay(conn, resultChan)
	for {
		opt, err := reader.ReadByte()
		if err != nil {
			if !errors.Is(err, io.EOF) {
				log.Println("close connection die to err: ", err.Error())
			}
			return
		}
		switch opt {
		case 'G':
			s.get(resultChan, reader)
		case 'S':
			s.set(resultChan, reader)
		case 'D':
			s.del(resultChan, reader)
		default:
			log.Println("unknown operator: ", opt)
			return
		}
	}

}
