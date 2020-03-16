package main

import (
	"context"
	"flag"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/etcdserver/api/v3rpc/rpctypes"
	"strings"
	"time"
)

var (
	endpoints = []string{"localhost:2379"}
	timeout   = 5 * time.Second
	hostIp    = "127.0.0.1:1234"
)

func init() {
	flag.StringVar(&hostIp, "ip", "-", "ip ")
	flag.Parse()
}

func Register(dir, value string) {
	dir = strings.TrimRight(dir, "/") + "/"
	client, err := clientv3.New(clientv3.Config{Endpoints: endpoints, DialTimeout: timeout})
	if err != nil {
		panic(err)
	}
	kv := clientv3.NewKV(client)
	lease := clientv3.NewLease(client)

	var leaseId clientv3.LeaseID = 0
	for {
		if leaseId == 0 {
			resp, err := lease.Grant(context.Background(), 10)
			if err != nil {
				break
			}
			key := dir + fmt.Sprintf("%d", resp.ID)
			fmt.Println("lease key => ", key)
			if _, err := kv.Put(context.Background(), key, value, clientv3.WithLease(resp.ID)); err != nil {
				break
			}
			leaseId = resp.ID
		} else {
			fmt.Println("keep lease id => ", leaseId)
			if _, err := lease.KeepAliveOnce(context.Background(), leaseId); err == rpctypes.ErrLeaseNotFound {
				leaseId = 0
			}
		}
		time.Sleep(time.Second * 2)
	}
	fmt.Println("lease invalid ...")
}

func main() {
	go Register("/agent", hostIp)
	for {
		time.Sleep(60 * time.Second)
	}
}
