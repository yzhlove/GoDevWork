package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net"
	"net/rpc"
	"sync"
	"time"
)

const KVServiceName = "path/to/pkg.KVStorageService"

type KeyChangeFunc = func(key string)

type KVStorageService struct {
	data   map[string]string
	filter map[string]KeyChangeFunc
	mutex  sync.Mutex
}

func NewKVStorageService() *KVStorageService {
	return &KVStorageService{
		data:   make(map[string]string),
		filter: make(map[string]KeyChangeFunc),
		mutex:  sync.Mutex{},
	}
}

func (p *KVStorageService) Get(key string, value *string) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if v, ok := p.data[key]; ok {
		*value = v
		return nil
	}
	return errors.New("nof found key:" + key)
}

func (p *KVStorageService) Set(kv [2]string, _ *string) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	key, value := kv[0], kv[1]
	if old := p.data[key]; old != value {
		for _, f := range p.filter {
			f(key)
		}
	}
	p.data[key] = value
	return nil
}

func (p *KVStorageService) Watch(ts int, keyChanged *string) error {
	id := fmt.Sprintf("watch-%s-%03d", time.Now(), rand.Int())
	ch := make(chan string, 10)
	p.mutex.Lock()
	p.filter[id] = func(key string) { ch <- key }
	p.mutex.Unlock()
	select {
	case <-time.After(time.Duration(ts) * time.Second):
		return errors.New("timeout")
	case key := <-ch:
		*keyChanged = key
	}
	return nil
}

func main() {
	rpc.RegisterName(KVServiceName, NewKVStorageService())
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		go rpc.ServeConn(conn)
	}
}
