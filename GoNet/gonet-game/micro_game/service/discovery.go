package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc"
	"path/filepath"
	"strings"
	"time"

	"micro_game/config"
	"sync"
)

var (
	ErrorNotFoundService = func(name string) string {
		return "not found service:" + name
	}
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
	etcd      *clientv3.Client
	callbacks map[string][]chan string
	sync.Mutex
}

func (p *pool) init(cfg *config.Config) {
	var err error
	if p.etcd, err = clientv3.New(
		clientv3.Config{Endpoints: cfg.EtcdHosts, DialTimeout: cfg.Timeout}); err != nil {
		log.Fatal(err)
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

func (p *pool) connect() error {
	kv := clientv3.NewKV(p.etcd)
	timeout := time.Second * 5
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if resp, err := kv.Get(ctx, p.root, clientv3.WithPrefix()); err != nil {
		return err
	} else {
		for _, e := range resp.Kvs {
			if _, ok := p.names[string(e.Key)]; ok {
				p.add(string(e.Key), string(e.Value))
			}
		}
	}
	return nil
}

func (p *pool) watcher() {
	go func() {
		for {
			if err := p.connect(); err != nil {
				log.Error(err)
				time.Sleep(5 * time.Second)
				continue
			}
			_watcher := clientv3.NewWatcher(p.etcd)
			for m := range _watcher.Watch(context.Background(), p.root, clientv3.WithPrefix()) {
				if !m.Canceled {
					for _, e := range m.Events {
						if _, ok := p.names[string(e.Kv.Key)]; ok {
							switch e.Type {
							case clientv3.EventTypePut:
								p.add(string(e.Kv.Key), string(e.Kv.Value))
							case clientv3.EventTypeDelete:
								p.remove(string(e.Kv.Key))
							}
						}
					}
					continue
				}
				break
			}
			_watcher.Close()
		}
	}()
}

func (p *pool) back(service, key string) {
	for _, ch := range p.callbacks[service] {
		select {
		case ch <- key:
		default:
		}
	}
}

func (p *pool) add(key, address string) {
	p.Lock()
	defer p.Unlock()
	name := filepath.Dir(key)
	if _, ok := p.names[name]; p.provided && !ok {
		log.Error(ErrorNotFoundService(name))
		return
	}
	service := p.services[name]
	if conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock()); err != nil {
		log.Error(err)
	} else {
		service.clients = append(service.clients, client{key: key, conn: conn})
		log.Info("service add:", key, address)
		p.back(name, key)
	}
}

func (p *pool) remove(key string) {
	p.Lock()
	defer p.Unlock()
	name := filepath.Dir(key)
	if _, ok := p.names[name]; p.provided && !ok {
		log.Error(ErrorNotFoundService(name))
		return
	}
	if service, ok := p.services[name]; ok && service != nil {
		for i, c := range service.clients {
			if c.key == key {
				c.conn.Close()
				service.clients = append(service.clients[:i], service.clients[i+1:]...)
				log.Info("service remove:", key)
				return
			}
		}
	}
	log.Error(ErrorNotFoundService(name))
}

var (
	_pool pool
	once  sync.Once
)

func Init(cfg *config.Config) {
	once.Do(func() {
		_pool.init(cfg)
	})
}
