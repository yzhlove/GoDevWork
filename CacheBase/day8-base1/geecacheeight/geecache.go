package geecacheeight

import (
	"fmt"
	"geecacheeight/pb"
	"geecacheeight/sign"
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
	return g.load(key)
}

func (g *Group) RegisterPeers(peers PeerPick) {
	if g.peers != nil {
		panic("RegisterPeerPick called more than once")
	}
	g.peers = peers
}

func (g *Group) load(key string) (value ByteView, err error) {
	v, err := g.loader.Do(key, func() (interface{}, error) {
		if g.peers != nil {
			if peer, ok := g.peers.PickPeer(key); ok {
				if value, err = g.getFromPeer(peer, key); err == nil {
					return value, nil
				}
				log.Println("[GeeCache] Failed to get from peer: ", err)
			}
		}
		return g.getLocally(key)
	})
	if err == nil {
		return v.(ByteView), nil
	}
	return
}

func (g *Group) populateCache(key string, value ByteView) {
	g.c.set(key, value)
}

func (g *Group) getLocally(key string) (ByteView, error) {
	bytes, err := g.getter.Get(key)
	if err != nil {
		return ByteView{}, err
	}
	value := ByteView{buf: _copy(bytes)}
	g.populateCache(key, value)
	return value, nil
}

func (g *Group) getFromPeer(peer PeerGetter, key string) (ByteView, error) {
	req := &pb.Cache_Req{Group: g.name, Key: key}
	resp := &pb.Cache_Resp{}
	if err := peer.Get(req, resp); err != nil {
		return ByteView{}, err
	}
	return ByteView{buf: resp.Value}, nil
}
