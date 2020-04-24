package main

import (
	"container/list"
	"log"
	"sync"
	"sync/atomic"
)

////////////////////////////////
// LRU
////////////////////////////////

func main() {

}

type CacheStatus struct {
	Gets        int64
	Hits        int64
	MaxItemSize int
	CurrentSize int
}

type Cache interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
	Del(key string)
	Status() *CacheStatus
}

type AtomicInt int64

type MemoryCache struct {
	sync.RWMutex
	maxItemSize int
	cacheList   *list.List
	cache       map[interface{}]*list.Element
	hits, gets  AtomicInt
}

type entry struct {
	key, value interface{}
}

func NewMemoryCache(max int) *MemoryCache {
	return &MemoryCache{
		maxItemSize: max,
		cacheList:   list.New(),
		cache:       make(map[interface{}]*list.Element),
	}
}

func (c *MemoryCache) Status() *CacheStatus {
	c.RLock()
	defer c.RUnlock()
	return &CacheStatus{
		MaxItemSize: c.maxItemSize,
		CurrentSize: c.cacheList.Len(),
		Gets:        c.gets.Get(),
		Hits:        c.hits.Get(),
	}
}

func (c *MemoryCache) Get(key string) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()
	c.gets.Add(1)
	if element, ok := c.cache[key]; ok {
		c.hits.Add(1)
		c.cacheList.MoveToFront(element)
		return element.Value.(*entry).value, true
	}
	return nil, false
}

func (c *MemoryCache) Set(key string, value interface{}) {
	c.Lock()
	defer c.Unlock()
	if c.cache == nil {
		c.cache = make(map[interface{}]*list.Element)
		c.cacheList = list.New()
	}
	if element, ok := c.cache[key]; ok {
		c.cacheList.MoveToFront(element)
		element.Value.(*entry).value = value
		return
	}
	element := c.cacheList.PushFront(&entry{key: key, value: value})
	c.cache[key] = element
	if c.maxItemSize != 0 && c.cacheList.Len() > c.maxItemSize {
		c.RemoveOldest()
	}
}

func (c *MemoryCache) Del(key string) {
	c.Lock()
	defer c.Unlock()
	if c.cache != nil {
		if element, ok := c.cache[key]; ok {
			c.cacheList.Remove(element)
			key := element.Value.(*entry).key
			delete(c.cache, key)
		}
	}
}

func (c *MemoryCache) RemoveOldest() {
	if c.cache != nil {
		if element := c.cacheList.Back(); element != nil {
			c.cacheList.Remove(element)
			key := element.Value.(*entry).key
			delete(c.cache, key)
		}
	}
}

func (c *MemoryCache) ShowValue() {
	for element := c.cacheList.Front(); element != nil; element = element.Next() {
		log.Print(element.Value.(*entry).key, " - ")
	}
	log.Println()
}

func (i *AtomicInt) Add(n int64) {
	atomic.AddInt64((*int64)(i), n)
}

func (i *AtomicInt) Get() int64 {
	return atomic.LoadInt64((*int64)(i))
}
