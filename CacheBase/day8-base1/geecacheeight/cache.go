package geecacheeight

import (
	"geecacheeight/lru"
	"sync"
)

type cache struct {
	mutex    sync.RWMutex
	lru      *lru.Cache
	capBytes int64
}

func (c *cache) set(key string, value ByteView) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.lru == nil {
		c.lru = lru.NewLRU(c.capBytes, nil)
	}
	c.lru.Set(key, value)
}

func (c *cache) get(key string) (value ByteView, ok bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	if c.lru != nil {
		if v, ok := c.lru.Get(key); ok {
			return v.(ByteView), ok
		}
	}
	return
}
