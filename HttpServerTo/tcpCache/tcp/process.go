package tcp

import (
	"bufio"
	"errors"
	"io"
	"log"
	"net"
)

func (s *Server) get(conn net.Conn, r *bufio.Reader) error {
	key, err := s.readKey(r)
	if err != nil {
		return err
	}
	value, err := s.Get(key)
	return sendResp(value, err, conn)
}

func (s *Server) set(conn net.Conn, r *bufio.Reader) error {
	key, value, err := s.readKeyAndValue(r)
	if err != nil {
		return err
	}
	return sendResp(nil, s.Set(key, value), conn)
}

func (s *Server) del(conn net.Conn, r *bufio.Reader) error {
	key, err := s.readKey(r)
	if err != nil {
		return err
	}
	return sendResp(nil, s.Del(key), conn)
}

func (s *Server) process(conn net.Conn) {
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("conn close err: " + err.Error())
		}
	}()
	in := bufio.NewReader(conn)
	for {
		opt, err := in.ReadByte()
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Println("close connection err: " + err.Error())
			}
			return
		}
		switch opt {
		case 'S':
			err = s.set(conn, in)
		case 'G':
			err = s.get(conn, in)
		case 'D':
			err = s.del(conn, in)
		default:
			log.Println("close connection invalid option")
			return
		}
		if err != nil {
			log.Println("close connection err:" + err.Error())
		}
	}
}
