package lru

import (
	"container/list"
)

type EvictFunc func(key string, value Value)

type Value interface {
	Len() int
}

type Cache struct {
	maxBytes, cntBytes int64
	cacheList          *list.List
	hashMap            map[string]*list.Element
	OnEvicted          EvictFunc
}

type Entry struct {
	key   string
	value Value
}

func NewLRU(maxBytes int64, event EvictFunc) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		cacheList: list.New(),
		hashMap:   make(map[string]*list.Element),
		OnEvicted: event,
	}
}

func (c *Cache) Set(key string, value Value) {
	if e, ok := c.hashMap[key]; ok {
		c.cacheList.MoveToFront(e)
		kv := e.Value.(*Entry)
		c.cntBytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		e := c.cacheList.PushFront(&Entry{key: key, value: value})
		c.hashMap[key] = e
		c.cntBytes += int64(len(key)) + int64(value.Len())
	}
	for c.maxBytes != 0 && c.maxBytes < c.cntBytes {
		c.RemoveOldest()
	}
}

func (c *Cache) Get(key string) (Value, bool) {
	if e, ok := c.hashMap[key]; ok {
		c.cacheList.MoveToFront(e)
		kv := e.Value.(*Entry)
		return kv.value, true
	}
	return nil, false
}

func (c *Cache) RemoveOldest() {
	if e := c.cacheList.Back(); e != nil {
		c.cacheList.Remove(e)
		kv := e.Value.(*Entry)
		delete(c.hashMap, kv.key)
		c.cntBytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

func (c *Cache) Len() int {
	return c.cacheList.Len()
}
