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

	client, err := clientv3.New(clientv3.Config{Endpoints: endpoints, DialTimeout: timeout})
	if err != nil {
		panic(err)
	}

	lease := clientv3.NewLease(client)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	resp, err := lease.Grant(ctx, 10)
	if err != nil {
		lease.Close()
		panic(err)
	}

	key := fmt.Sprintf("%s_%d", "chat", resp.ID)
	fmt.Println("generate key => ", key)

	putResp, err := clientv3.NewKV(client).Put(context.Background(), key, "hello world", clientv3.WithLease(resp.ID))
	if err != nil {
		panic(err)
	}

	fmt.Println("putResp => ", putResp)

	for i := 0; i < 5; i++ {
		fmt.Println("keep lease id => ", resp.ID)
		lease.KeepAliveOnce(context.Background(), resp.ID)
		time.Sleep(5 * time.Second)
	}

	fmt.Println("keep over ...")

}
