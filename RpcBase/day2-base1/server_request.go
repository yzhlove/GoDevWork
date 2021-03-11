package day2_base1

import (
	"day2-base1-example/codec"
	"reflect"
)

type req struct {
	header       *codec.Header
	args, replay reflect.Value
}

func (s *server) readReqHead(cc codec.Codec) (*codec.Header, error) {
	var header = &codec.Header{}
	if err := cc.ReadHeader(header); err != nil {
		return nil, err
	}
	return header, nil
}

func (s *server) readReq(cc codec.Codec) (*req, error) {
	header, err := s.readReqHead(cc)
	if err != nil {
		return nil, err
	}
	r := &req{header: header}
	r.args = reflect.New(reflect.TypeOf(""))
	return r, cc.ReadBody(r.args.Interface())
}
