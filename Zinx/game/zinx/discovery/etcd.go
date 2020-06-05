package discovery

import (
	"context"
	"errors"
	"fmt"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"go.etcd.io/etcd/clientv3"
	"log"
	"path/filepath"
	"time"
	"zinx/config"
)

type Etcd struct {
	cli      *clientv3.Client    //etcd client
	monitors map[string]struct{} //需要监控的服务
}

func NewEtcd() (*Etcd, error) {
	if len(config.Endpoints) == 0 {
		return nil, errors.New("endpoints is nil")
	}
	if len(config.Monitors) == 0 {
		return nil, errors.New("monitor services is nil")
	}
	cli, err := clientv3.New(
		clientv3.Config{Endpoints: config.Endpoints, DialTimeout: config.DialTimeout})
	if err != nil {
		return nil, err
	}
	etcd := &Etcd{
		cli:      cli,
		monitors: make(map[string]struct{}, len(config.Monitors)),
	}
	for _, service := range config.Monitors {
		etcd.monitors[service] = struct{}{}
	}
	return etcd, nil
}

func (etcd *Etcd) Register(key, value string) {
	kv := clientv3.NewKV(etcd.cli)
	lease := clientv3.NewLease(etcd.cli)
	go func() {
		var _leaseID clientv3.LeaseID
		for {
			if _leaseID == 0 {
				grant, err := lease.Grant(context.Background(), 5)
				if err != nil {
					log.Println("register service: lease grant err:", err)
					return
				}
				_leaseID = grant.ID
				key = fmt.Sprintf("%s/%s/%d", config.Root, key, _leaseID)
				log.Println("==> register key:", key, " value:", value)
				if _, err := kv.Put(context.Background(), key, value, clientv3.WithLease(_leaseID)); err != nil {
					log.Println("register service: put value err:", err)
					return
				}
			} else {
				keep, err := lease.KeepAlive(context.Background(), _leaseID)
				if err != nil {
					log.Println("register service: keep err:", err)
					return
				}
				for ch := range keep {
					if ch == nil {
						break
					}
				}
				_leaseID = 0
			}
		}
	}()
}

func (etcd *Etcd) Watcher() <-chan Event {
	kv := clientv3.NewKV(etcd.cli)
	_ch := make(chan Event, config.EventMax)
	go func() {
		for {
			result, err := kv.Get(context.Background(), config.Root, clientv3.WithPrefix())
			if err != nil {
				log.Println("watcher service: get root err:", err)
				time.Sleep(time.Second * 5)
				continue
			}
			for _, event := range result.Kvs {
				if _, ok := etcd.monitors[filepath.Dir(string(event.Key))]; ok {
					_ch <- Event{Action: EventPut, Key: string(event.Key), Addr: string(event.Value)}
				}
			}
			break
		}
		watcher := clientv3.NewWatcher(etcd.cli)
		defer watcher.Close()
		for watch := range watcher.Watch(context.Background(), config.Root, clientv3.WithPrefix()) {
			if watch.Canceled {
				return
			}
			for _, event := range watch.Events {
				if _, ok := etcd.monitors[filepath.Dir(string(event.Kv.Key))]; ok {
					var eventType EventType
					switch event.Type {
					case mvccpb.PUT:
						eventType = EventPut
					case mvccpb.DELETE:
						eventType = EventDel
					}
					_ch <- Event{Action: eventType, Key: string(event.Kv.Key), Addr: string(event.Kv.Value)}
				}
			}
		}
	}()
	return _ch
}
