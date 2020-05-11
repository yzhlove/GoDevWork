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

func getEntry(e *list.Element) *entry {
	if kv, ok := e.Value.(*entry); ok {
		return kv
	}
	typeofError()
	return nil
}

func NewLRU(capacity int64, event Event) *Cache {
	return &Cache{
		capBytes: capacity,
		vector:   list.New(),
		hashMap:  make(map[string]*list.Element),
		Callback: event,
	}
}

func (c *Cache) Set(key string, value Value) {
	if v, ok := c.hashMap[key]; ok {
		c.vector.MoveToFront(v)
		if kv := getEntry(v); kv != nil {
			c.cntBytes += int64(value.Len()) - int64(kv.value.Len())
			kv.value = value
		}
	} else {
		e := c.vector.PushFront(&entry{key: key, value: value})
		c.cntBytes += int64(len(key)) + int64(value.Len())
		c.hashMap[key] = e
	}
	for c.capBytes != 0 && c.capBytes < c.cntBytes {
		c.RemoveOldest()
	}
}

func (c *Cache) Get(key string) (Value, bool) {
	if v, ok := c.hashMap[key]; ok {
		c.vector.MoveToFront(v)
		if kv := getEntry(v); kv != nil {
			return kv.value, true
		}
	}
	return nil, false
}

func (c *Cache) RemoveOldest() {
	if v := c.vector.Back(); v != nil {
		c.vector.Remove(v)
		if kv := getEntry(v); kv != nil {
			delete(c.hashMap, kv.key)
			c.cntBytes -= int64(len(kv.key)) + int64(kv.value.Len())
			if c.Callback != nil {
				c.Callback(kv.key, kv.value)
			}
		}
	}
}

func (c *Cache) Len() int {
	return c.vector.Len()
}
