package day1_base1

import (
	"day1_base1_example/codec"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"reflect"
	"sync"
)

const MagicNumber = 0x3bf5c

type Option struct {
	MagicNumber int
	CodecType   codec.Type
}

var DefaultOption = &Option{
	MagicNumber: MagicNumber,
	CodecType:   codec.GobType,
}

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

var DefaultServer = NewServer()

func (server *Server) ServeConn(conn io.ReadWriteCloser) {
	defer func() { conn.Close() }()
	var opt Option
	if err := json.NewDecoder(conn).Decode(&opt); err != nil {
		log.Println("rpc server:options error: ", err)
		return
	}

	if opt.MagicNumber != MagicNumber {
		log.Printf("rpc server:invalid magic number %x", opt.MagicNumber)
		return
	}

	if fn := codec.NewCodecFuncMap[opt.CodecType]; fn == nil {
		log.Printf("rpc server: invalid codec type %s", opt.CodecType)
	} else {
		server.serveCodec(fn(conn))
	}

}

var errRequest = ""

func (server *Server) serveCodec(cc codec.Codec) {
	m := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	for {
		req, err := server.readRequest(cc)
		if err != nil {
			if req == nil {
				break
			}
			req.h.Error = err.Error()
			server.sendResponse(cc, req.h, errRequest, m)
			continue
		}
		wg.Add(1)
		go server.handleRequest(cc, req, m, wg)
	}
	wg.Wait()
	cc.Close()
}

type request struct {
	h           *codec.Header
	argv, reply reflect.Value
}

func (server *Server) readRequestHeader(cc codec.Codec) (*codec.Header, error) {
	var header codec.Header
	if err := cc.ReadHeader(&header); err != nil {
		if errors.Is(err, io.EOF) && errors.Is(err, io.ErrUnexpectedEOF) {
			log.Println("rpc server:read header error: ", err)
		}
		return nil, err
	}
	return &header, nil
}

func (server *Server) readRequest(cc codec.Codec) (*request, error) {
	header, err := server.readRequestHeader(cc)
	if err != nil {
		return nil, err
	}
	req := &request{h: header}
	req.argv = reflect.New(reflect.TypeOf(""))
	if err := cc.ReadBody(req.argv.Interface()); err != nil {
		log.Println("rpc server: read argv err:", err)
	}
	return req, nil
}

func (server *Server) sendResponse(cc codec.Codec, header *codec.Header, body interface{}, mutex *sync.Mutex) {
	mutex.Lock()
	defer mutex.Unlock()
	if err := cc.Write(header, body); err != nil {
		log.Println("rpc server: write response error: ", err)
	}
}

func (server *Server) handleRequest(cc codec.Codec, req *request, mutex *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println("client info data -->:", req.h, req.argv.Elem())
	req.reply = reflect.ValueOf(fmt.Sprintf("geegrpc resp %d", req.h.Seq))
	server.sendResponse(cc, req.h, req.reply.Interface(), mutex)
}

func (server *Server) Accept(l net.Listener) {
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("rcp server:accept error: ", err)
			return
		}
		go server.ServeConn(conn)
	}
}

func Accept(lis net.Listener) {
	DefaultServer.Accept(lis)
}
