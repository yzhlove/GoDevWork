package day3_base1

import (
	"day3-base1-example/codec"
	"encoding/json"
	"io"
	"log"
	"net"
	"reflect"
)

const MagicCode = 0x3bf5c

type Auth struct {
	Code int
	Type codec.CodecsType
}

type server struct{}

func (s *server) accept(l net.Listener) {
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("[rpc server] accept error:", err)
			return
		}
		go s.handleConn(conn)
	}
}

func (s *server) handleConn(conn io.ReadWriteCloser) {
	defer conn.Close()

	a := &Auth{}
	if err := json.NewDecoder(conn).Decode(a); err != nil {
		log.Println("[rpc server] json decoder error:", err)
		return
	}

	if fn := codec.CodecsMap[a.Type]; fn != nil {
		s.handleCodec(fn(conn))
	} else {
		log.Println("[rpc server] codec func not found")
	}
}

func (s *server) handleCodec(cc codec.Coder) {

}

type request struct {
	header      *codec.Header
	args, reply reflect.Value
}

func (s *server) readReqHeader(cc codec.Coder) (*codec.Header, error) {
	header := &codec.Header{}
	if err := cc.ReadHeader(header); err != nil {
		return nil, err
	}
	return header, nil
}

func (s *server) readReqBody(cc codec.Coder) (*request, error) {
	header, err := s.readReqHeader(cc)
	if err != nil {
		return nil, err
	}
	req := &request{header: header}
	req.args = reflect.New(reflect.TypeOf(""))
	return req, cc.ReadBody(req.args.Interface())
}
