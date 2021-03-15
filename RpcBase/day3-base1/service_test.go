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
