package geethree

import (
	"errors"
	"log"
	"sync"
)

type Getter interface {
	Get(key string) ([]byte, error)
}

type GetterFunc func(key string) ([]byte, error)

func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

type Group struct {
	name   string
	getter Getter
	m      manager
}

var (
	_mutex         sync.RWMutex
	_groups        = make(map[string]*Group)
	requiredKeyErr = errors.New("key is required")
)

func NewGroup(name string, maxBytes int64, getter Getter) *Group {
	if getter == nil {
		panic("getter not is nil")
	}
	_mutex.Lock()
	defer _mutex.Unlock()
	group := &Group{
		name:   name,
		getter: getter,
		m:      manager{maxBytes: maxBytes},
	}
	_groups[name] = group
	return group
}

func GetGroup(name string) *Group {
	_mutex.RLock()
	defer _mutex.RUnlock()
	return _groups[name]
}

func (g *Group) Get(key string) (ByteView, error) {
	if key == "" {
		return ByteView{}, requiredKeyErr
	}
	if v, ok := g.m.get(key); ok {
		log.Println("hit.")
		return v, nil
	}
	return g.load(key)
}

func (g *Group) load(key string) (view ByteView, err error) {
	var bytes []byte
	if bytes, err = g.getter.Get(key); err != nil {
		return
	}
	view = ByteView{buffer: _copy(bytes)}
	g.m.set(key, view)
	return
}
