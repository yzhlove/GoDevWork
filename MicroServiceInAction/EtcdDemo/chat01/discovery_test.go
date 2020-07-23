package main

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"testing"
	"time"
)

func TestNewRegistry(t *testing.T) {
	var op = Options{
		name: "svc.info",
		ttl:  10,
		config: clientv3.Config{
			Endpoints:   []string{"localhost:2379"},
			DialTimeout: 5 * time.Second,
		},
	}

	for i := 1; i <= 3; i++ {
		res, err := NewRegistry(op)
		if err != nil {
			t.Error(err)
			return
		}
		if err = res.RegistryNode(PutNode{Addr: fmt.Sprintf("127.0.0.0:%d%d%d%d", i, i, i, i)}); err != nil {
			t.Error(err)
			return
		}
		if i == 3 {
			go func() {
				var index int
				for range time.NewTicker(time.Second).C {
					index++
					fmt.Println("time tick index -> ", index)
					if index >= 10 {
						fmt.Printf("exit tick and unRegistry")
						res.UnRegistry()
						return
					}
				}
			}()
		}
	}
	time.Sleep(15 * time.Second)
}

func TestNewSelector(t *testing.T) {
	var op = SelectorOptions{
		name: "svc.info",
		config: clientv3.Config{
			Endpoints:   []string{"localhost:2379"},
			DialTimeout: 5 * time.Second,
		},
	}
	s, err := NewSelector(op)
	if err != nil {
		t.Error(err)
		return
	}
	for {
		time.Sleep(time.Second * 2)
		val, err := s.Next()
		if err != nil {
			t.Error(err)
			continue
		}
		t.Log(val)
	}
}
