package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/sirupsen/logrus"
	"go.etcd.io/etcd/clientv3"
	"strings"
	"sync"
	"time"
)

var (
	endpoints = []string{"127.0.0.1:2379"}
	timeout   = time.Second * 5
)

type ServiceDiscovery struct {
	dir   string
	mutex sync.Mutex
	nodes map[string]string
}

func NewDiscovery(dir string) (discovery *ServiceDiscovery) {
	discovery = &ServiceDiscovery{
		dir:   strings.TrimRight(dir, "/") + "/",
		nodes: make(map[string]string, 4),
	}
	return
}

func (discovery *ServiceDiscovery) watch(ctx context.Context) {
	client, err := clientv3.New(clientv3.Config{Endpoints: endpoints, DialTimeout: timeout})
	if err != nil {
		panic(err)
	}
	defer client.Close()

	//var curRevision int64 = 0
	kv := clientv3.NewKV(client)
	for {
		getResp, err := kv.Get(context.Background(), discovery.dir, clientv3.WithPrefix())
		if err != nil {
			logrus.Info("watch key ", discovery.dir, " waiting .")
			time.Sleep(time.Second)
			continue
		}
		discovery.mutex.Lock()
		for _, kv := range getResp.Kvs {
			discovery.nodes[string(kv.Key)] = string(kv.Value)
		}
		discovery.mutex.Unlock()
		//从当前版本开始watch
		//curRevision = getResp.Header.Revision + 1
		break
	}
	watcher := clientv3.NewWatcher(client)
	defer watcher.Close()
	watchChans := watcher.Watch(ctx, discovery.dir, clientv3.WithPrefix())
	for watchResp := range watchChans {
		if watchResp.Canceled {
			break
		}
		for _, event := range watchResp.Events {
			discovery.mutex.Lock()
			switch event.Type {
			case mvccpb.PUT:
				fmt.Println("PUT => ", string(event.Kv.Value))
				discovery.nodes[string(event.Kv.Key)] = string(event.Kv.Value)
			case mvccpb.DELETE:
				fmt.Println("DELETE => ", string(event.Kv.Value))
				delete(discovery.nodes, string(event.Kv.Key))
			}
			discovery.mutex.Unlock()
		}
	}
	fmt.Println("watcher close ...")
}

func (discovery *ServiceDiscovery) Nodes() []string {

	dupNodes := make(map[string]bool, len(discovery.nodes))
	nodes := make([]string, 0, len(discovery.nodes))
	discovery.mutex.Lock()
	for _, endpoint := range discovery.nodes {
		dupNodes[endpoint] = true
	}
	discovery.mutex.Unlock()
	for endpoint := range dupNodes {
		nodes = append(nodes, endpoint)
	}
	return nodes
}

func main() {

	discovery := NewDiscovery("/agent")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	go discovery.watch(ctx)
	cnt := 0
	for {
		cnt++
		if cnt > 60 {
			cancel()
			break
		}
		time.Sleep(time.Second)
		fmt.Println("nodes => ", discovery.Nodes())
	}
	time.Sleep(5 * time.Second)

}
