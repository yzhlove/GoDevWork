package cond

import "WorkSpace/GoDevWork/Events/chat02/event"

type CancelFunc = func()

type Condition interface {
	Event() event.Event
	SetTarget(data []uint32)
	SetCancel(fn CancelFunc)
	Ok() bool
	Cancel()
}

type base struct {
	target []uint32
	fn     CancelFunc
}

func (base) Event() event.Event {
	return event.Unknown{}
}

func (b base) SetTarget(data []uint32) {
	b.target = data
}

func (b base) SetCancel(fn CancelFunc) {
	b.fn = fn
}

func (b base) Ok() bool {
	return false
}

func (b base) Cancel() {
	if b.fn != nil {
		b.fn()
		b.fn = nil
	}
}

var conditions map[string]creatorFunc

type creatorFunc func() Condition

func Get(name string) (fn creatorFunc, ok bool) {
	fn, ok = conditions[name]
	return
}
