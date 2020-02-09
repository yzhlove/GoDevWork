package tcp

import (
	"bufio"
	"errors"
	"io"
	"log"
	"net"
)

type result struct {
	value []byte
	err   error
}

func newResultChan(resultCh chan chan *result) chan *result {
	r := make(chan *result)
	resultCh <- r
	return r
}

func (s *Server) get(resultCh chan chan *result, reader *bufio.Reader) {
	r := newResultChan(resultCh)
	k, err := s.readKey(reader)
	if err != nil {
		r <- &result{nil, err}
		return
	}
	go func() {
		if v, err := s.Get(k); err != nil {
			return
		} else {
			r <- &result{v, nil}
		}
	}()
}

func (s *Server) set(resultCh chan chan *result, reader *bufio.Reader) {
	r := newResultChan(resultCh)
	k, v, err := s.readKeyAndValue(reader)
	if err != nil {
		r <- &result{nil, err}
		return
	}
	go func() {
		r <- &result{nil, s.Set(k, v)}
	}()
}

func (s *Server) del(resultCh chan chan *result, reader *bufio.Reader) {
	r := newResultChan(resultCh)
	k, err := s.readKey(reader)
	if err != nil {
		r <- &result{nil, err}
		return
	}
	go func() {
		r <- &result{nil, s.Del(k)}
	}()
}

func replay(conn net.Conn, resultCh chan chan *result) {
	defer conn.Close()
	for {
		if ch, ok := <-resultCh; ok {
			r := <-ch
			if err := setResp(r.value, r.err, conn); err != nil {
				log.Println("set resp err: ", err.Error())
			}
		} else {
			return
		}
	}
}

func (s *Server) process(conn net.Conn) {
	reader := bufio.NewReader(conn)
	resultCh := make(chan chan *result, 512)
	defer close(resultCh)
	go replay(conn, resultCh)
	for {
		opt, err := reader.ReadByte()
		if err != nil {
			if !errors.Is(err, io.EOF) {
				log.Println("close conn err: ", err.Error())
			}
			return
		}
		switch opt {
		case 'G':

		case 'S':

		case 'D':

		default:

		}
	}
}
