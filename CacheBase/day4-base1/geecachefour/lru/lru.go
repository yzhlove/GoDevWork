package lru

import "container/list"

var _typeErrString = "[cache ] type is err"

func _TypeErr() {
	panic(_typeErrString)
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
	maxBytes    int64
	cntBytes    int64
	cacheList   *list.List
	hashMap     map[string]*list.Element
	OnEventFunc Event
}

func NewLRU(maxBytes int64, event Event) *Cache {
	return &Cache{
		maxBytes:    maxBytes,
		cacheList:   list.New(),
		hashMap:     make(map[string]*list.Element),
		OnEventFunc: event,
	}
}

func (c *Cache) Set(key string, value Value) {
	if v, ok := c.hashMap[key]; ok {
		c.cacheList.MoveToFront(v)
		if kv, ok := v.Value.(*entry); ok {
			c.cntBytes += int64(value.Len()) - int64(kv.value.Len())
			kv.value = value
		} else {
			_TypeErr()
		}
	} else {
		e := c.cacheList.PushFront(&entry{key: key, value: value})
		c.cntBytes += int64(len(key)) + int64(value.Len())
		c.hashMap[key] = e
	}
	for c.maxBytes != 0 && c.maxBytes < c.cntBytes {
		c.RemoveOldest()
	}
}

func (c *Cache) Get(key string) (Value, bool) {
	if v, ok := c.hashMap[key]; ok {
		c.cacheList.MoveToFront(v)
		if e := v.Value.(*entry); ok {
			return e.value, true
		} else {
			_TypeErr()
		}
	}
	return nil, false
}

func (c *Cache) RemoveOldest() {
	if v := c.cacheList.Back(); v != nil {
		c.cacheList.Remove(v)
		if kv, ok := v.Value.(*entry); ok {
			delete(c.hashMap, kv.key)
			c.cntBytes -= int64(len(kv.key)) + int64(kv.value.Len())
			if c.OnEventFunc != nil {
				c.OnEventFunc(kv.key, kv.value)
			}
		} else {
			_TypeErr()
		}
	}
}

func (c *Cache) Len() int {
	return c.cacheList.Len()
}
