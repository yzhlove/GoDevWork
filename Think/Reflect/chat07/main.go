package main

import (
	"fmt"
	"go/ast"
	"reflect"
	"strings"
)

//反射调用函数
type Class struct {
	Name    string
	RefTyp  reflect.Type
	Self    reflect.Value
	Methods map[string]*Method
}

type Method struct {
	RefMethod reflect.Method
	RefArgs   reflect.Type
	RefReply  reflect.Type
}

func NewRefClass(v interface{}) *Class {
	class := &Class{}
	class.RefTyp = reflect.TypeOf(v)
	class.Name = class.RefTyp.Name()
	class.Self = reflect.ValueOf(v)
	class.load()
	return class
}

func (m *Method) NewArgs() reflect.Value {
	if m.RefArgs.Kind() == reflect.Ptr {
		return reflect.New(m.RefArgs.Elem())
	}
	return reflect.New(m.RefArgs).Elem()
}

func (m *Method) NewReply() reflect.Value {
	rep := reflect.New(m.RefReply.Elem())
	switch m.RefReply.Kind() {
	case reflect.Slice:
		rep.Elem().Set(reflect.MakeSlice(m.RefReply.Elem(), 0, 0))
	case reflect.Map:
		rep.Elem().Set(reflect.MakeMap(m.RefReply.Elem()))
	}
	return rep
}

func (c *Class) load() {
	if methods := c.RefTyp.NumMethod(); methods > 0 {
		c.Methods = make(map[string]*Method, methods)
		for i := 0; i < methods; i++ {
			method := c.RefTyp.Method(i)
			if method.Type.NumIn() != 3 {
				continue
			}
			if method.Type.NumOut() != 1 {
				continue
			}
			if method.Type.Out(0) != reflect.TypeOf((*error)(nil)).Elem() {
				continue
			}
			args, reply := method.Type.In(1), method.Type.In(2)
			if checkExport(args) && checkExport(reply) {
				c.Methods[method.Name] = &Method{
					RefMethod: method,
					RefArgs:   args,
					RefReply:  reply,
				}
			}
		}
	}
}

func checkExport(tp reflect.Type) bool {
	return ast.IsExported(tp.Name()) || tp.PkgPath() == ""
}

func (c *Class) Call(method *Method, args, replay reflect.Value) error {
	res := method.RefMethod.Func.Call([]reflect.Value{c.Self, args, replay})
	if err := res[0].Interface(); err != nil {
		return err.(error)
	}
	return nil
}

type Args struct {
	S string
	T string
	N string
}

type A struct{}

func (a *A) Append(args Args, replay *string) error {
	*replay = strings.Join([]string{args.N, args.S, args.T}, "-")
	return nil
}

func main() {

	a := &A{}
	cs := NewRefClass(a)
	mt := cs.Methods["Append"]
	args := mt.NewArgs()
	args.Set(reflect.ValueOf(Args{S: "ssss", T: "ttttt", N: "nnnnnn"}))
	reply := mt.NewReply()
	if err := cs.Call(mt, args, reply); err != nil {
		panic(err)
	}
	fmt.Print("reply => ", reply.Elem())
}
