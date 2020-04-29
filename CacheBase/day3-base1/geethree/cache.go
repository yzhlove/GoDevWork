package geethree

import (
	"geethree/lru"
	"sync"
)

type manager struct {
	sync.RWMutex
	cache    *lru.Cache
	maxBytes int64
}

func (m *manager) set(key string, value ByteView) {
	m.Lock()
	defer m.Unlock()
	if m.cache == nil {
		m.cache = lru.NewLRU(m.maxBytes, nil)
	}
	m.cache.Set(key, value)
}

func (m *manager) get(key string) (ByteView, bool) {
	m.RLock()
	defer m.RUnlock()
	if m.cache != nil {
		if v, ok := m.cache.Get(key); ok {
			if value, ok := v.(ByteView); ok {
				return value, true
			}
		}
	}
	return ByteView{}, false
}
