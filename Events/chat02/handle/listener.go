package handle

import (
	"WorkSpace/GoDevWork/Events/chat02/context"
	"WorkSpace/GoDevWork/Events/chat02/event"
	"container/list"
)

type (
	Func       = func(ctx *context.Context, evt event.Event) error
	CancelFunc = func()
)

type EventContext struct {
	Source map[uint32]*list.List
	Queue  *list.List
}

func Construct() *EventContext {
	return &EventContext{
		Source: make(map[uint32]*list.List),
		Queue:  list.New(),
	}
}

func (e *EventContext) AddListener(evt event.Event, fn Func) CancelFunc {
	if _, ok := e.Source[evt.ID()]; !ok {
		e.Source[evt.ID()] = list.New()
	}
	element := e.Source[evt.ID()].PushBack(fn)
	return func() { e.Source[evt.ID()].Remove(element) }
}

func (e *EventContext) Notice(evt event.Event) {
	e.Queue.PushBack(evt)
}

func (e *EventContext) Awake(ctx *context.Context) error {
	for element := e.Queue.Front(); element != nil; element = element.Next() {
		if evt, ok := element.Value.(event.Event); ok {
			if conds, ok := e.Source[evt.ID()]; ok {
				for ele := conds.Front(); ele != nil; {
					next := ele.Next()
					if handleFunc, ok := ele.Value.(Func); ok {
						if err := handleFunc(ctx, evt); err != nil {
							return err
						}
					}
					conds.Remove(ele)
					ele = next
				}
			}
		}
	}
	e.Clean()
	return nil
}

func (e *EventContext) Clean() {
	e.Queue.Init()
}
