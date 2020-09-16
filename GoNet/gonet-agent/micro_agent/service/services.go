package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc"
	"micro_agent/config"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type client struct {
	key  string
	conn *grpc.ClientConn
}

type service struct {
	clients []client
	idx     uint32
}

type pool struct {
	root      string
	names     map[string]struct{}
	services  map[string]*service
	provided  bool
	client    *clientv3.Client
	callbacks map[string][]chan string
	sync.RWMutex
}

func (p *pool) init(cfg *config.Config) {
	var err error
	if p.client, err = clientv3.New(
		clientv3.Config{Endpoints: cfg.EtcdHost, DialTimeout: cfg.Timeout}); err != nil {
		log.Panic(err)
	}
	p.root = cfg.EtcdRoot
	p.services = make(map[string]*service)
	p.names = make(map[string]struct{})
	if len(cfg.Services) > 0 {
		p.provided = true
		for _, name := range cfg.Services {
			p.names[p.root+"/"+strings.TrimSpace(name)] = struct{}{}
		}
	}
	p.watcher()
}

func (p *pool) watcher() {
	kv := clientv3.NewKV(p.client)
	go func() {
		wait := 5 * time.Second
		for {
			ctx, cancel := context.WithTimeout(context.Background(), wait)
			if resp, err := kv.Get(ctx, p.root, clientv3.WithPrefix()); err != nil {
				cancel()
				log.Error("etcd watcher service err:", err)
				time.Sleep(wait)
			} else {
				cancel()
				for _, event := range resp.Kvs {
					if _, ok := p.names[string(event.Key)]; ok {
						p.add(string(event.Key), string(event.Value))
					}
				}
				break
			}
		}
		_watch := clientv3.NewWatcher(p.client)
		defer _watch.Close()
		for w := range _watch.Watch(context.Background(), p.root, clientv3.WithPrefix()) {
			if w.Canceled {
				return
			}
			for _, e := range w.Events {
				if _, ok := p.names[string(e.Kv.Key)]; ok {
					switch e.Type {
					case clientv3.EventTypePut:
						p.add(string(e.Kv.Key), string(e.Kv.Value))
					case clientv3.EventTypeDelete:
						p.remove(string(e.Kv.Key))
					}
				}
			}
		}
	}()
}

func (p *pool) add(key, value string) {
	p.Lock()
	defer p.Unlock()
	serviceName := filepath.Dir(key)
	if _, ok := p.names[serviceName]; p.provided && !ok {
		log.Error("not found service:", key, value)
		return
	}
	if _, ok := p.services[serviceName]; !ok {
		p.services[serviceName] = &service{}
	}
	s := p.services[serviceName]
	if conn, err := grpc.Dial(value, grpc.WithBlock(), grpc.WithInsecure()); err != nil {
		log.Error("connect err:", key, value, err)
	} else {
		s.clients = append(s.clients, client{key: key, conn: conn})
		log.Info("service add:", key, value)
		for _, ch := range p.callbacks[serviceName] {
			select {
			case ch <- key:
			default:
			}
		}
	}
}

func (p *pool) remove(key string) {
	p.Lock()
	defer p.Unlock()

	serviceName := filepath.Dir(key)
	if _, ok := p.names[serviceName]; p.provided && !ok {
		return
	}

	s := p.services[serviceName]
	if s == nil {
		log.Error("no such service:", serviceName)
		return
	}

	for k, c := range s.clients {
		if c.key == key {
			c.conn.Close()
			s.clients = append(s.clients[:k], s.clients[k+1:]...)
			log.Info("service remove:", key)
			return
		}
	}
}

func (p *pool) getServiceWithId(path, id string) *grpc.ClientConn {
	p.RLock()
	defer p.RUnlock()
	s := p.services[path]
	if s == nil || len(s.clients) == 0 {
		return nil
	}
	name := path + "/" + id
	for _, c := range s.clients {
		if c.key == name {
			return c.conn
		}
	}
	return nil
}

func (p *pool) getService(path string) (*grpc.ClientConn, string) {
	p.RLock()
	defer p.RUnlock()

	s := p.services[path]
	if s == nil || len(s.clients) == 0 {
		return nil, ""
	}

	i := int(atomic.AddUint32(&s.idx, 1)) % len(s.clients)
	return s.clients[i].conn, s.clients[i].key
}

func (p *pool) registerCallback(path string, callback chan string) {
	p.RLock()
	defer p.RUnlock()

	if p.callbacks == nil {
		p.callbacks = make(map[string][]chan string)
	}

	p.callbacks[path] = append(p.callbacks[path], callback)
	if s, ok := p.services[path]; ok {
		for _, c := range s.clients {
			callback <- c.key
		}
	}
	log.Info("register callback on:", path)
}

var (
	local_pool pool
	once       sync.Once
)

func Init(cfg *config.Config) {
	once.Do(func() {
		local_pool.init(cfg)
	})
}
