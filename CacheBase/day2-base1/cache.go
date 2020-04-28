package geecachetwo

import (
	"day2-base1/lru"
	"sync"
)

type cache struct {
	sync.RWMutex
	cacheLRU *lru.Cache
	maxBytes int64
}

func (c *cache) set(key string, value ByteView) {
	c.Lock()
	defer c.Unlock()
	if c.cacheLRU == nil {
		c.cacheLRU = lru.NewLRU(c.maxBytes, nil)
	}
	c.cacheLRU.Set(key, value)
}

func (c *cache) get(key string) (ByteView, bool) {
	c.RLock()
	defer c.RUnlock()
	if c.cacheLRU != nil {
		if view, ok := c.cacheLRU.Get(key); ok {
			return view.(ByteView), ok
		}
	}
	return ByteView{}, false
}
