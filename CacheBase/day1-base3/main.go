package main

import (
	"container/list"
	"sync"
)

////////////////////////////////
// LRU
////////////////////////////////

func main() {

}

type AtomicInt = int64

type MemoryCache struct {
	mutex       sync.RWMutex
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
		cacheList: list.New(),
		cache: make(map[interface{}]*list.Element),
	}
}

func (c *MemoryCache) Status() {

}
