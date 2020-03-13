package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var (
	endpoints = []string{"127.0.0.1:2379"}
	timeout   = 5 * time.Second
)

func main() {

	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints, DialTimeout: timeout})
	if err != nil {
		panic(err)
	}

	kv := clientv3.NewKV(cli)
	resp, err := kv.Get(context.Background(), "/gift", clientv3.WithPrefix())
	if err != nil {
		panic(err)
	}

	for k, v := range resp.Kvs {
		fmt.Println(" k => ", k)
		fmt.Printf("%v \n", v)
	}

	fmt.Println(resp)

}
