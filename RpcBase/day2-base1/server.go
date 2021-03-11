package day2_base1

import (
	"day2-base1-example/codec"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"reflect"
	"sync"
)

const MagicCode = 0x3bf5c

type MsgHead struct {
	Code int
	Type codec.Type
}

type server struct{}

func (s *server) accept(l net.Listener) {
	for {
		if conn, err := l.Accept(); err != nil {
			log.Println("rpc server:accept error:", err)
			return
		} else {
			go s.svcConn(conn)
		}
	}
}

func (s *server) svcConn(conn io.ReadWriteCloser) {
	defer conn.Close()

	var slot MsgHead
	if err := json.NewDecoder(conn).Decode(&slot); err != nil {
		log.Println("rpc server:slot error->", err)
		return
	}

	if slot.Code != MagicCode {
		log.Println("rpc server:coder error->", slot.Code)
		return
	}

	if f := codec.CodecsMap[slot.Type]; f != nil {
		s.svcCodec(f(conn))
	} else {
		log.Println("rpc server:codec func is nil")
	}
}

func (s *server) svcCodec(cc codec.Codec) {
	var (
		m  sync.Mutex
		wg sync.WaitGroup
	)
	for {
		if r, err := s.readReq(cc); err != nil {
			if r == nil { //header read error
				break
			}
			r.header.Err = err.Error()
			if err := s.send(cc, r.header, "", &m); err != nil {
				log.Println("rpc server:send resp error->", err)
			}
		} else {
			wg.Add(1)
			go s.sendReply(cc, r, &m, &wg)
		}
	}
	wg.Wait()
	cc.Close()
}

func (s *server) send(cc codec.Codec, header *codec.Header, body interface{}, m *sync.Mutex) error {
	m.Lock()
	defer m.Unlock()
	if err := cc.Writer(header, body); err != nil {
		return err
	}
	return nil
}

func (s *server) sendReply(cc codec.Codec, r *req, m *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println("[server->client]", r.header, r.args.Elem())
	r.replay = reflect.ValueOf(fmt.Sprintf("[server] header seq:(%d)", r.header.Seq))
	if err := s.send(cc, r.header, r.replay.Interface(), m); err != nil {
		log.Println("rpc server:send client error->", err)
	}
}

var DefaultServer = &server{}

func Accept(l net.Listener) {
	DefaultServer.accept(l)
}
