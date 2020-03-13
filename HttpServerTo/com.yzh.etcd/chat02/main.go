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

	resp, err := kv.Put(context.Background(), "/gift/zone_one", "192.168.0.1:1234")
	if err != nil {
		panic(err)
	}
	fmt.Println("resp => ", resp)
	resp, err = kv.Put(context.Background(), "/gift/zone_two", "192.168.0.2:1234")
	if err != nil {
		panic(err)
	}
	fmt.Println("resp => ", resp)

	resp, err = kv.Put(context.Background(), "/giftservice", "192.168.0.2:1234")
	if err != nil {
		panic(err)
	}
	fmt.Println("resp => ", resp)

}
