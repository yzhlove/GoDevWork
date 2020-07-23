package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Selector interface {
	Next() (Node, error)
}

type selectorServer struct {
	cli  *clientv3.Client
	node []Node
	opt  SelectorOptions
}

type SelectorOptions struct {
	name   string
	config clientv3.Config
}

func NewSelector(opt SelectorOptions) (Selector, error) {
	cli, err := clientv3.New(opt.config)
	if err != nil {
		return nil, err
	}
	var s = &selectorServer{
		opt: opt,
		cli: cli,
	}
	go s.Watch()
	return s, nil
}

func (s *selectorServer) Next() (Node, error) {
	if len(s.node) == 0 {
		return Node{}, errors.New("not found node " + s.opt.name)
	}
	i := rand.Int() % len(s.node)
	return s.node[i], nil
}

func (s *selectorServer) Watch() {
	result, err := s.cli.Get(context.Background(), s.GetKey(), clientv3.WithPrefix(), clientv3.WithSerializable())
	if err != nil {
		log.Printf("[Watch] err: %s", err.Error())
		return
	}
	for _, kv := range result.Kvs {
		node, err := s.GetVal(kv.Value)
		if err != nil {
			log.Printf("[GetVal] err: %s ", err.Error())
			continue
		}
		s.node = append(s.node, node)
	}
	for c := range s.cli.Watch(context.Background(), prefix, clientv3.WithPrefix()) {
		if !c.Canceled {
			for _, e := range c.Events {
				switch e.Type {
				case clientv3.EventTypePut:
					node, err := s.GetVal(e.Kv.Value)
					if err != nil {
						log.Printf("[EventTypePut] err: %s", err.Error())
						continue
					}
					s.AddNode(node)
				case clientv3.EventTypeDelete:
					keys := strings.Split(string(e.Kv.Key), "/")
					if len(keys) == 0 {
						log.Printf("[EventTypeDel] key split 0")
						return
					}
					nodeId, err := strconv.Atoi(keys[len(keys)-1])
					if err != nil {
						log.Printf("[EventTypeDel] transform err")
						continue
					}
					s.DelNode(uint32(nodeId))
				}
			}
		}
	}

}

func (s *selectorServer) DelNode(id uint32) {
	var tNode []Node
	for _, v := range s.node {
		if v.Id != id {
			tNode = append(tNode, v)
		}
	}
	s.node = tNode
}

func (s *selectorServer) AddNode(node Node) {
	for _, v := range s.node {
		if v.Id == node.Id {
			return
		}
	}
	s.node = append(s.node, node)
}
func (s *selectorServer) GetKey() string {
	return fmt.Sprintf("%s%s", prefix, s.opt.name)
}

func (s *selectorServer) GetVal(val []byte) (Node, error) {
	var node Node
	if err := json.Unmarshal(val, &node); err != nil {
		return Node{}, err
	}
	return node, nil
}
