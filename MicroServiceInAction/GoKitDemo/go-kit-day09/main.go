package main

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd/etcdv3"
	"hash/crc32"
	"time"
)

var (
	prometheus_addr = "127.0.0.1:2345"
	grpc_addr       = "127.0.0.1:1234"
	etcd_addr       = []string{"127.0.0.1:2379"}
)

func main() {

	serviceName := "service.user.agent"
	ttl := 5 * time.Second

	opts := etcdv3.ClientOptions{DialKeepAlive: ttl, DialTimeout: ttl}
	client, err := etcdv3.NewClient(context.Background(), etcd_addr, opts)
	if err != nil {
		panic(err)
	}
	register := etcdv3.NewRegistrar(client,
		etcdv3.Service{
			Key:   fmt.Sprintf("%s/%s", serviceName, crc32.ChecksumIEEE([]byte(grpc_addr))),
			Value: grpc_addr,
		}, log.NewNopLogger())

	//Jaeger
	go func() {
		
	}()

}
