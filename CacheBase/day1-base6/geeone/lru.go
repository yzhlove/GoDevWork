package geeone

import (
	"container/list"
	"fmt"
	"strings"
)

type Value interface {
	Len() int
}

type entry struct {
	key   string
	value Value
}

type EvictFunc func(key string, value Value)

type Cache struct {
	maxBytes  int64
	cntBytes  int64
	cacheList *list.List
	hashMap   map[string]*list.Element
	OnEvicted EvictFunc
}

func New(maxBytes int64, onEvicted EvictFunc) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		cacheList: list.New(),
		hashMap:   make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

func (e *entry) String() string {
	return fmt.Sprintf(" key:%v value:%v \n", e.key, e.value)
}

func (c *Cache) Show() {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("====================[cache]====================\n"))
	sb.WriteString(fmt.Sprintf("|\tmaxBytes:%d cntBytes:%d \n", c.maxBytes, c.cntBytes))
	sb.WriteString(fmt.Sprintf("|\tHashMap--> \n"))
	for key, element := range c.hashMap {
		sb.WriteString(fmt.Sprintf("|\t\t key:%v  => %v \n", key, element.Value.(*entry).String()))
	}
	sb.WriteString(fmt.Sprintf("|\tcacheList--> \n"))
	for e := c.cacheList.Front(); e != nil; e = e.Next() {
		sb.WriteString(fmt.Sprintf("|\t\t value:%v \n", e.Value.(*entry).String()))
	}
	sb.WriteString(fmt.Sprintf("===============================================\n"))
	fmt.Print(sb.String())
}

func (c *Cache) Set(key string, value Value) {
	if e, ok := c.hashMap[key]; ok {
		c.cacheList.MoveToFront(e)
		kv := e.Value.(*entry)
		c.cntBytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		element := c.cacheList.PushFront(&entry{key: key, value: value})
		c.hashMap[key] = element
		c.cntBytes += int64(len(key)) + int64(value.Len())
	}
	for c.maxBytes != 0 && c.maxBytes < c.cntBytes {
		c.RemoveOldest()
	}
}

func (c *Cache) Get(key string) (value Value, ok bool) {
	if e, ok := c.hashMap[key]; ok {
		c.cacheList.MoveToFront(e)
		kv := e.Value.(*entry)
		return kv.value, true
	}
	return
}

func (c *Cache) RemoveOldest() {
	if e := c.cacheList.Back(); e != nil {
		c.cacheList.Remove(e)
		kv := e.Value.(*entry)
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
