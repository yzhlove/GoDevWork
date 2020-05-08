package lru

import (
	"container/list"
)

func typeofError() {
	panic("[cache] typeof err")
}

type Value interface {
	Len() int
}

type Event func(key string, value Value)

type entry struct {
	key   string
	value Value
}

type Cache struct {
	capBytes int64
	cntBytes int64
	vector   *list.List
	hashMap  map[string]*list.Element
	Callback Event
}

func NewLRU(capicaty int64, event Event) *Cache {
	return &Cache{
		capBytes: capicaty,
		vector:   list.New(),
		hashMap:  make(map[string]*list.Element),
		Callback: event,
	}
}
