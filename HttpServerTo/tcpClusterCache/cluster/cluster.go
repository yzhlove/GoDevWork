package cluster

import (
	"github.com/hashicorp/memberlist"
	"io/ioutil"
	"stathat.com/c/consistent"
	"strconv"
	"time"
)

type Node interface {
	ShouldProcess(key string) (string, bool)
	Members() []string
	TcpAddress() string
	HttpAddress() string
}

type node struct {
	*consistent.Consistent
	tcpAddress  string
	httpAddress string
}

func (n *node) TcpAddress() string {
	return n.tcpAddress
}

func (n *node) HttpAddress() string {
	return n.httpAddress
}

func New(tcpAddr, httpAddr, cluster string, gossipPort int) (Node, error) {
	if cluster == "" {
		cluster = "127.0.0.1:" + strconv.Itoa(gossipPort)
	}
	conf := memberlist.DefaultLANConfig()
	conf.Name = "127.0.0.1:" + strconv.Itoa(gossipPort)
	conf.BindAddr = "127.0.0.1"
	conf.BindPort = gossipPort
	conf.LogOutput = ioutil.Discard
	list, err := memberlist.Create(conf)
	if err != nil {
		return nil, err
	}
	if _, err = list.Join([]string{cluster}); err != nil {
		return nil, err
	}
	circle := consistent.New()
	circle.NumberOfReplicas = 256
	go func() {
		for {
			members := list.Members()
			nodes := make([]string, len(members))
			for i, n := range members {
				nodes[i] = n.Name
			}
			circle.Set(nodes)
			time.Sleep(time.Second)
		}
	}()
	return &node{circle, tcpAddr, httpAddr}, nil
}

func (n *node) ShouldProcess(key string) (string, bool) {
	addr, _ := n.Get(key)
	return addr, addr == n.tcpAddress
}
