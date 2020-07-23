package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"hash/crc32"
	"log"
	"time"
)

var prefix = "/register/server/"

type Registry interface {
	RegistryNode(node PutNode) error
	UnRegistry()
}

type registryServer struct {
	cli        *clientv3.Client
	stop       chan struct{}
	isRegistry bool
	opt        Options
	leaseID    clientv3.LeaseID
}

type PutNode struct {
	Addr string `json:"addr"`
}

type Node struct {
	Id   uint32 `json:"id"`
	Addr string `json:"addr"`
}

type Options struct {
	name   string
	ttl    int64
	config clientv3.Config
}

func NewRegistry(opts Options) (Registry, error) {
	cli, err := clientv3.New(opts.config)
	if err != nil {
		return nil, err
	}
	return &registryServer{
		stop:       make(chan struct{}),
		opt:        opts,
		isRegistry: false,
		cli:        cli,
	}, nil
}

func (s *registryServer) RegistryNode(put PutNode) error {
	if s.isRegistry {
		return errors.New("only one node can be registered")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.opt.ttl)*time.Second)
	defer cancel()

	grant, err := s.cli.Grant(context.Background(), s.opt.ttl)
	if err != nil {
		return err
	}
	var node = Node{
		Id:   s.HashKey(put.Addr),
		Addr: put.Addr,
	}
	nodeVal, err := s.GetVal(node)
	if err != nil {
		return err
	}
	_, err = s.cli.Put(ctx, s.GetKey(node), nodeVal, clientv3.WithLease(s.leaseID))
	if err != nil {
		return err
	}
	s.leaseID = grant.ID
	s.isRegistry = true
	go s.KeepAlive()
	return nil
}

func (s *registryServer) UnRegistry() {
	s.stop <- struct{}{}
}

func (s *registryServer) Revoke() (err error) {
	if _, err = s.cli.Revoke(context.Background(), s.leaseID); err != nil {
		log.Printf("[Revoke] err:%s", err.Error())
	}
	s.isRegistry = false
	return err
}

func (s *registryServer) KeepAlive() error {
	keepAliveCh, err := s.cli.KeepAlive(context.Background(), s.leaseID)
	if err != nil {
		log.Printf("[KeepAlive] err:%s", err.Error())
		return err
	}
	for {
		select {
		case <-s.stop:
			s.Revoke()
			return nil
		case ch, ok := <-keepAliveCh:
			if !ok || ch == nil {
				s.Revoke()
				return nil
			}
		}
	}
}

func (s *registryServer) GetKey(node Node) string {
	return fmt.Sprintf("%s%s/%d", prefix, s.opt.name, s.HashKey(node.Addr))
}

func (s *registryServer) GetVal(node Node) (string, error) {
	data, err := json.Marshal(&node)
	return string(data), err
}

func (e *registryServer) HashKey(addr string) uint32 {
	return crc32.ChecksumIEEE([]byte(addr))
}
