package cache

import "sync"

type inMemoryCache struct {
	data map[string][]byte
	sync.RWMutex
	Stat
}

func (c *inMemoryCache) Set(key string, value []byte) error {
	c.Lock()
	defer c.Unlock()
	if v, ok := c.data[key]; ok {
		c.del(key, v)
	}
	c.data[key] = value
	c.add(key, value)
	return nil
}

func (c *inMemoryCache) Get(key string) ([]byte, error) {
	c.Lock()
	defer c.Unlock()
	return c.data[key], nil
}

func (c *inMemoryCache) Del(key string) error {
	c.Lock()
	defer c.Unlock()
	if v, ok := c.data[key]; ok {
		delete(c.data, key)
		c.del(key, v)
	}
	return nil
}

func (c *inMemoryCache) GetStat() Stat {
	return c.Stat
}

func newInMemoryCache() *inMemoryCache {
	return &inMemoryCache{make(map[string][]byte), sync.RWMutex{}, Stat{}}
}
