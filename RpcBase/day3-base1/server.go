package day3_base1

import (
	"day3-base1-example/codec"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net"
	"reflect"
	"strings"
	"sync"
)

const MagicCode = 0x3bf5c

type Auth struct {
	Code int
	Type codec.CodecsType
}

type server struct {
	services sync.Map
}

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
	send := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	for {
		req, err := s.readReq(cc)
		if err != nil {
			if req == nil {
				break
			}
			req.header.Err = err.Error()
			s.sendResp(cc, req.header, nil, send)
			continue
		}
		wg.Add(1)
		go s.handleReq(cc, req, send, wg)
	}

	wg.Wait()
	_ = cc.Close()
}

func (s *server) register(service interface{}) error {
	svc := newService(service)
	if _, ok := s.services.LoadOrStore(svc.name, svc); ok {
		return errors.New("rpc server: service already defined:" + svc.name)
	}
	return nil
}

func (s *server) findService(svcMethod string) (svc *service, mtyp *methodType, err error) {
	if idx := strings.LastIndex(svcMethod, "."); idx != -1 {
		class, method := svcMethod[:idx], svcMethod[idx+1:]
		if svi, ok := s.services.Load(class); !ok {
			err = errors.New("rpc server:not found service:" + svcMethod)
			return
		} else {
			svc = svi.(*service)
			mtyp = svc.method[method]
			if mtyp == nil {
				err = errors.New("rpc server:not found method:" + method)
			}
		}
	} else {
		err = errors.New("rpc server:service method struct invalid:" + svcMethod)
	}
	return
}

func (s *server) sendResp(cc codec.Coder, header *codec.Header, body interface{}, send *sync.Mutex) {
	send.Lock()
	defer send.Unlock()
	if err := cc.Send(header, body); err != nil {
		log.Println("rpc server:write response error:", err)
	}
}

func (s *server) handleReq(cc codec.Coder, req *request, send *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	if err := req.svc.call(req.mtyp, req.argv, req.replay); err != nil {
		req.header.Err = err.Error()
		s.sendResp(cc, req.header, nil, send)
		return
	}
	s.sendResp(cc, req.header, req.replay.Interface(), send)
}

type request struct {
	header       *codec.Header
	argv, replay reflect.Value
	mtyp         *methodType
	svc          *service
}

func (s *server) readReqHeader(cc codec.Coder) (*codec.Header, error) {
	header := &codec.Header{}
	if err := cc.ReadHeader(header); err != nil {
		return nil, err
	}
	return header, nil
}

func (s *server) readReq(cc codec.Coder) (*request, error) {
	header, err := s.readReqHeader(cc)
	if err != nil {
		return nil, err
	}
	req := &request{header: header}
	req.svc, req.mtyp, err = s.findService(header.Method)
	if err != nil {
		return nil, err
	}
	req.argv = req.mtyp.newArg()
	req.replay = req.mtyp.newReply()
	//取指针
	argvi := req.argv.Interface()
	if req.argv.Type().Kind() != reflect.Ptr {
		argvi = req.argv.Addr().Interface()
	}

	if err = cc.ReadBody(argvi); err != nil {
		log.Println("rpc server: read body err:", err)
		return nil, err
	}
	return req, nil
}

func NewServer() *server {
	return &server{}
}

var DefServer = NewServer()

func Register(class interface{}) error {
	return DefServer.register(class)
}

func Accept(l net.Listener) {
	DefServer.accept(l)
}
