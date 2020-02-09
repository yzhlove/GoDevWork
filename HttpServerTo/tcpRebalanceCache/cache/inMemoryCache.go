package cache

import "sync"

type inMemoryCache struct {
	data  map[string][]byte
	mutex sync.RWMutex
	Stat
}

func (c *inMemoryCache) Set(k string, v []byte) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if v, ok := c.data[k]; ok {
		c.del(k, v)
	}
	c.data[k] = v
	c.add(k, v)
	return nil
}

func (c *inMemoryCache) Get(k string) ([]byte, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.data[k], nil
}

func (c *inMemoryCache) Del(k string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if v, ok := c.data[k]; ok {
		delete(c.data, k)
		c.del(k, v)
	}
	return nil
}

func (c *inMemoryCache) GetStat() Stat {
	return c.Stat
}

func newInMemory() *inMemoryCache {
	return &inMemoryCache{
		data:  make(map[string][]byte, 16),
		mutex: sync.RWMutex{},
		Stat:  Stat{},
	}
}
