package geecacheseven

import (
	"fmt"
	"geecacheseven/sign"
	"log"
	"sync"
)

type Getter interface {
	Get(key string) ([]byte, error)
}

type Group struct {
	name   string
	getter Getter
	c      cache
	peers  PeerPick
	loader *sign.Group
}

type GetterFunc func(key string) ([]byte, error)

func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

var (
	_mutex  sync.Mutex
	_groups = make(map[string]*Group)
)

func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	if getter == nil {
		panic("Getter is nil")
	}
	_mutex.Lock()
	defer _mutex.Unlock()
	g := &Group{
		name:   name,
		getter: getter,
		c:      cache{capBytes: cacheBytes},
		loader: &sign.Group{},
	}
	_groups[name] = g
	return g
}

func GetGroup(name string) *Group {
	_mutex.Lock()
	defer _mutex.Unlock()
	return _groups[name]
}

func (g *Group) Get(key string) (ByteView, error) {
	if key == "" {
		return ByteView{}, fmt.Errorf("key is required")
	}
	if v, ok := g.c.get(key); ok {
		log.Println("[GeeCache] hit.")
		return v, nil
	}

}

func (g *Group) load(key string) (value ByteView, err error) {
	v, err := g.loader.Do(key, func() (interface{}, error) {

	})
	if err != nil {
		return v.(ByteView), nil
	}
	return
}
