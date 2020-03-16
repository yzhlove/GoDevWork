package main

import (
	"context"
	"flag"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"net"
	"net/rpc"
	"strings"
	"time"
)

var (
	endpoints = []string{"localhost:2379"}
	timeout   = 5 * time.Second
	hostIp    = "localhost:1234"
)

func init() {
	flag.StringVar(&hostIp, "ip", "127.0.0.1:1234", "ip address")
	flag.Parse()
}

type HelloService struct{}

func (p *HelloService) Hello(req string, replay *string) error {
	*replay = hostIp + " => " + req
	return nil
}

func RegisterServiceToEtcd(service, value string) {
	root := strings.TrimRight(service, "/") + "/"
	c, err := clientv3.New(clientv3.Config{Endpoints: endpoints, DialTimeout: timeout})
	if err != nil {
		panic(err)
	}
	kv := clientv3.NewKV(c)
	lease := clientv3.NewLease(c)
	var leaseId clientv3.LeaseID
	for {
		if leaseId == 0 {
			leaseResp, err := lease.Grant(context.Background(), 5)
			if err != nil {
				panic(err)
			}
			key := root + fmt.Sprintf("%d", leaseResp.ID)
			fmt.Println("newKey ==> ", key)
			if _, err := kv.Put(context.Background(), key, value, clientv3.WithLease(leaseResp.ID)); err != nil {
				panic(err)
			}
			leaseId = leaseResp.ID
		} else {
			if _, err := lease.KeepAliveOnce(context.Background(), leaseId); err != nil {
				leaseId = 0
				continue
			}
		}
		time.Sleep(2 * time.Second)
	}
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))
	l, err := net.Listen("tcp", hostIp)
	if err != nil {
		panic(err)
	}
	go RegisterServiceToEtcd("hello", hostIp)
	for {
		c, err := l.Accept()
		if err != nil {
			panic(err)
		}
		rpc.ServeConn(c)
	}
}
