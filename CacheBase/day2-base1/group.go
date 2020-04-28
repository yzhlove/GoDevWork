package geecachetwo

import (
	"errors"
	"log"
	"sync"
)

type Getter interface {
	Get(key string) ([]byte, error)
}

type GetterFunc func(key string) ([]byte, error)

type Group struct {
	name     string
	getter   Getter
	newcache cache
}

func (fn GetterFunc) Get(key string) ([]byte, error) {
	return fn(key)
}

var (
	_mutex  sync.RWMutex
	_groups = make(map[string]*Group)
)

func NewGroup(name string, maxBytes int64, getter Getter) *Group {
	if getter == nil {
		panic("nil Getter")
	}
	_mutex.Lock()
	defer _mutex.Unlock()
	group := &Group{
		name:     name,
		getter:   getter,
		newcache: cache{maxBytes: maxBytes},
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
		return ByteView{}, errors.New("key is required.")
	}
	if v, ok := g.newcache.get(key); ok {
		log.Println("[GeeCache] hit")
		return v, nil
	}
	return g.load(key)
}

func (g *Group) load(key string) (ByteView, error) {
	return g.getLocally(key)
}

func (g *Group) getLocally(key string) (view ByteView, err error) {
	var bytes []byte
	if bytes, err = g.getter.Get(key); err != nil {
		return
	}
	view = ByteView{buf: cloneBytes(bytes)}
	g.populateCache(key, view)
	return
}

func (g *Group) populateCache(key string, value ByteView) {
	g.newcache.set(key, value)
}
