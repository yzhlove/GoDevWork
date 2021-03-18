package day3_base1

import (
	"reflect"
	"testing"
)

type Foo int

type Args struct {
	Num1, Num2 int
}

func (f Foo) Sum(args Args, reply *int) error {
	*reply = args.Num1 + args.Num2
	return nil
}

func (f Foo) sum(args Args, reply *int) error {
	*reply = args.Num1 + args.Num2
	return nil
}

func TestNewService(t *testing.T) {
	var foo Foo
	s := newService(&foo)
	t.Log("methods ", len(s.method))
	t.Log(s.method)

	if mtyp := s.method["Sum"]; mtyp == nil {
		t.Error("found method error")
	}

}

func TestMethodType_Call(t *testing.T) {

	var foo Foo
	s := newService(&foo)
	mtyp := s.method["Sum"]

	argv := mtyp.newArg()
	reply := mtyp.newReply()
	argv.Set(reflect.ValueOf(Args{Num1: 100, Num2: 200}))
	if err := s.call(mtyp, argv, reply); err != nil {
		t.Error("call error", err)
	}
	t.Log(*reply.Interface().(*int))
	t.Log(mtyp.NumCalls())

}

type View int

type Reqs struct {
	S1, S2, S3 string
}

type Resps struct {
	SS []string
}

func (v *View) Set(st Reqs, reply *Resps) error {
	reply = &Resps{}
	reply.SS = append(reply.SS, []string{st.S1, st.S2, st.S3}...)
	return nil
}

func TestSetSlice(t *testing.T) {

	var v View
	s := newService(&v)
	mt := s.method["Set"]

	argv := mt.newArg()
	reply := mt.newReply()
	argv.Set(reflect.ValueOf(Reqs{"a", "b", "c"}))
	if err := s.call(mt, argv, reply); err != nil {
		panic(err)
	}
	t.Log(*(reply.Interface().(*Resps)))
}
