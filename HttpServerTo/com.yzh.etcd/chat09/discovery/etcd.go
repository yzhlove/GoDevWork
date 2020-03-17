package discovery

import (
	"WorkSpace/GoDevWork/HttpServerTo/com.yzh.etcd/chat09/conf"
	"context"
	"errors"
	"fmt"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/sirupsen/logrus"
	"go.etcd.io/etcd/clientv3"
	"path/filepath"
	"strings"
	"time"
)

type Etcd struct {
	etcdClient *clientv3.Client
	svc        map[string]struct{}
}

func (etcd *Etcd) Init() error {
	if len(conf.Endpoints) == 0 {
		return errors.New("endpoints is nil")
	}
	if len(conf.MonitorServ) == 0 {
		return errors.New("monitor server is nil")
	}
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: conf.Endpoints, DialTimeout: conf.DialTimeout,
	})
	if err != nil {
		return err
	}
	etcd.etcdClient = cli
	etcd.svc = make(map[string]struct{}, len(conf.MonitorServ))
	for _, serviceName := range conf.MonitorServ {
		etcd.svc[conf.Root+"/"+serviceName] = struct{}{}
	}
	return nil
}

func (etcd *Etcd) Register(key, value string) {
	kv := clientv3.NewKV(etcd.etcdClient)
	lease := clientv3.NewLease(etcd.etcdClient)
	go func() {
		var id clientv3.LeaseID = 0
		for {
			if id == 0 {
				if result, err := lease.Grant(context.Background(), 5); err != nil {
					logrus.Error("[Register] grant err:", err)
				} else {
					id = result.ID
					key = conf.Root + "/" + strings.TrimSpace(key) + "/" + fmt.Sprintf("%d", id)
					if _, err = kv.Put(context.Background(), key, value, clientv3.WithLease(id)); err != nil {
						logrus.Error("[Register] put key err:", err)
					}
				}
			} else {
				if _, err := lease.KeepAliveOnce(context.Background(), id); err != nil {
					id = 0
				}
			}
			time.Sleep(2 * time.Second)
		}
	}()
}

func (etcd *Etcd) Watcher() chan Event {
	kv := clientv3.NewKV(etcd.etcdClient)
	eventChans := make(chan Event, conf.EventMax)
	go func() {
		for {
			if result, err := kv.Get(context.Background(), conf.Root, clientv3.WithPrefix()); err != nil {
				logrus.Error("[Watcher] get root err:", err)
			} else {
				for _, event := range result.Kvs {
					serviceName := filepath.Dir(string(event.Key))
					if _, ok := etcd.svc[serviceName]; ok {
						eventChans <- Event{Action: PUT, Key: string(event.Key), Addr: string(event.Value)}
					}
				}
				break
			}
			time.Sleep(time.Second)
		}
		w := clientv3.NewWatcher(etcd.etcdClient)
		watchChans := w.Watch(context.Background(), conf.Root, clientv3.WithPrefix())
		defer w.Close()
		for watchResp := range watchChans {
			if watchResp.Canceled {
				break
			}
			for _, event := range watchResp.Events {
				serviceName := filepath.Dir(string(event.Kv.Key))
				if _, ok := etcd.svc[serviceName]; ok {
					switch event.Type {
					case mvccpb.PUT:
						eventChans <- Event{Action: PUT, Key: string(event.Kv.Key), Addr: string(event.Kv.Value)}
					case mvccpb.DELETE:
						eventChans <- Event{Action: DELETE, Key: string(event.Kv.Key), Addr: string(event.Kv.Value)}
					}
				}
			}
		}
	}()
	return eventChans
}

func NewEtcd() *Etcd {
	return &Etcd{}
}
