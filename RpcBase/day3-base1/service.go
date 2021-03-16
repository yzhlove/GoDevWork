package day3_base1

import (
	"go/ast"
	"log"
	"reflect"
	"sync/atomic"
)

type methodType struct {
	method    reflect.Method
	ArgType   reflect.Type
	ReplyType reflect.Type
	numCalls  uint64
}

func (m *methodType) NumCalls() uint64 {
	return atomic.LoadUint64(&m.numCalls)
}

func (m *methodType) newArg() reflect.Value {
	var argv reflect.Value
	if m.ArgType.Kind() == reflect.Ptr {
		argv = reflect.New(m.ArgType.Elem())
	} else {
		argv = reflect.New(m.ArgType).Elem()
	}
	return argv
}

func (m *methodType) newReply() reflect.Value {
	reply := reflect.New(m.ReplyType.Elem())
	switch m.ReplyType.Elem().Kind() {
	case reflect.Map:
		reply.Elem().Set(reflect.MakeMap(m.ReplyType.Elem()))
	case reflect.Slice:
		reply.Elem().Set(reflect.MakeSlice(m.ReplyType.Elem(), 0, 0))
	}
	return reply
}

type service struct {
	name   string
	rtyp   reflect.Type
	this   reflect.Value
	method map[string]*methodType
}

func newService(st interface{}) *service {
	s := &service{}
	s.this = reflect.ValueOf(st)
	s.name = reflect.Indirect(s.this).Type().Name()
	s.rtyp = reflect.TypeOf(st)
	if !ast.IsExported(s.name) {
		log.Fatalf("rpc server: %s is not a valid service name ", s.name)
	}
	s.registerMethods()
	return s
}

func (s *service) registerMethods() {
	s.method = make(map[string]*methodType, s.rtyp.NumMethod())
	for i := 0; i < s.rtyp.NumMethod(); i++ {
		ms := s.rtyp.Method(i)
		mt := ms.Type
		if mt.NumIn() != 3 || mt.NumOut() != 1 {
			continue
		}
		if mt.Out(0) != reflect.TypeOf((*error)(nil)).Elem() {
			continue
		}
		argType, replyTyp := mt.In(1), mt.In(2)
		if !isCheck(argType) || !isCheck(replyTyp) {
			continue
		}
		s.method[ms.Name] = &methodType{
			method:    ms,
			ArgType:   argType,
			ReplyType: replyTyp,
		}
		log.Printf("==> rpc server: register %s.%s\n", s.name, ms.Name)
	}
}

func (s *service) call(m *methodType, argv, reply reflect.Value) error {
	atomic.AddUint64(&m.numCalls, 1)
	fn := m.method.Func
	res := fn.Call([]reflect.Value{s.this, argv, reply})
	if err := res[0].Interface(); err != nil {
		return err.(error)
	}
	return nil
}

func isCheck(t reflect.Type) bool {
	return ast.IsExported(t.Name()) || t.PkgPath() == ""
}
