package cluster

import (
	"github.com/hashicorp/memberlist"
	"io/ioutil"
	"stathat.com/c/consistent"
	"time"
)

type Node interface {
	ShouldProcess(string) (string, bool)
	Members() []string
	Address() string
}

type node struct {
	*consistent.Consistent
	address string
}

func (n *node) Address() string {
	return n.address
}

func (n *node) ShouldProcess(key string) (string, bool) {
	addr, _ := n.Get(key)
	return addr, addr == n.address
}

func New(addr, cluster string) (Node, error) {
	conf := memberlist.DefaultLANConfig()
	conf.Name = addr
	conf.BindAddr = addr
	conf.LogOutput = ioutil.Discard
	if cluster == "" {
		cluster = addr
	}
	members, err := memberlist.Create(conf)
	if err != nil {
		return nil, err
	}
	if _, err = members.Join([]string{cluster}); err != nil {
		return nil, err
	}
	c := consistent.New()
	c.NumberOfReplicas = 256
	go func() {
		for {
			address := members.Members()
			nodes := make([]string, len(address))
			for i, ip := range address {
				nodes[i] = ip.Name
			}
			c.Set(nodes)
			time.Sleep(time.Second)
		}
	}()
	return &node{c, addr}, nil
}
