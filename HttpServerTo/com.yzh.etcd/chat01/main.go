package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		panic(err)
	}

	defer cli.Close()

	kv := clientv3.NewKV(cli)

	resp, err := kv.Put(context.Background(), "foo", "bar")
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)

}
