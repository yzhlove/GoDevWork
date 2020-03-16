package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"go.etcd.io/etcd/clientv3"
	"net/rpc"
	"sync"
	"time"
)

var (
	endpoints = []string{"localhost:2379"}
	timeout   = 5 * time.Second
)

type rpcService struct {
	root  string
	nodes map[string]string
	mutex sync.RWMutex
}

func getNodes(etcdClient *clientv3.Client) *rpcService {
	s := &rpcService{root: "hello/", nodes: make(map[string]string, 4)}
	kv := clientv3.NewKV(etcdClient)

	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		rangeResp, err := kv.Get(ctx, s.root, clientv3.WithPrefix())
		cancel()
		if err != nil {
			panic(err)
		}
		if len(rangeResp.Kvs) == 0 {
			fmt.Println(">>>>>>> node list is 0 <<<<<<<")
			time.Sleep(time.Second)
			continue
		}
		s.mutex.Lock()
		for _, kv := range rangeResp.Kvs {
			s.nodes[string(kv.Key)] = string(kv.Value)
		}
		s.mutex.Unlock()
		break
	}
	go watchServiceUpdate(etcdClient, s)
	return s
}

func watchServiceUpdate(etcdClient *clientv3.Client, s *rpcService) {
	watcher := clientv3.NewWatcher(etcdClient)
	watchChans := watcher.Watch(context.Background(), s.root, clientv3.WithPrefix())
	for watchResp := range watchChans {
		if watchResp.Canceled {
			break
		}
		for _, event := range watchResp.Events {
			s.mutex.Lock()
			switch event.Type {
			case mvccpb.PUT:
				s.nodes[string(event.Kv.Key)] = string(event.Kv.Value)
			case mvccpb.DELETE:
				delete(s.nodes, string(event.Kv.Key))
			}
			s.mutex.Unlock()
		}
	}
	fmt.Println("watch exit ....")
}

func main() {

	client, err := clientv3.New(clientv3.Config{Endpoints: endpoints, DialTimeout: timeout})
	if err != nil {
		panic(err)
	}
	s := getNodes(client)

	for {
		fmt.Println("================================================")
		for _, node := range s.nodes {
			rpcClient, err := rpc.Dial("tcp", node)
			if err != nil {
				fmt.Println(node, " => connection ", err.Error())
				continue
			}
			var replay string
			err = rpcClient.Call("HelloService.Hello", node, &replay)
			if err != nil {
				fmt.Println(node, " => rpc call ", err.Error())
				continue
			}
			fmt.Println(node, " rpc succeed => ", replay)
			rpcClient.Close()
		}
		fmt.Println("================================================")
		time.Sleep(time.Second * 5)
	}

}
