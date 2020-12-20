package main

import "container/list"

type EventFunc func(evt Cond)

type Listener struct {
	Id     uint32
	Source EventFunc
}

type EventManager struct {
	Data  map[uint32]*list.List
	Queue *list.List
}

func Construct() *EventManager {
	return &EventManager{Data: make(map[uint32]*list.List), Queue: list.New()}
}

func (e *EventManager) AddListener(evt Event, cancel EventFunc) func() {
	if _, ok := e.Data[evt.Id()]; !ok {
		e.Data[evt.Id()] = list.New()
	}
	element := e.Data[evt.Id()].PushBack(cancel)
	return func() { e.Data[evt.Id()].Remove(element) }
}

func (e *EventManager) Trigger(evt Event) {
	e.Queue.PushBack(evt)
}

func (e *EventManager) Dispatch() error {
	for element := e.Queue.Front(); element != nil; element = element.Next() {
		if res, ok := e.Data[element.Value.(Event).Id()]; ok {
			for ele := res.Front(); ele != nil; {
				next := ele.Next()
				res.Remove(ele)
				ele = next
			}
		}
	}
	return nil
}
