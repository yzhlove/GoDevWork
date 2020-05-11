package geecacheeight

import (
	"context"
	"errors"
	"fmt"
	"geecacheeight/consistent"
	"geecacheeight/pb"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
	"time"
)

type RpcContext struct {
	self       string
	mutex      sync.Mutex
	peers      *consistent.Map
	rpcGetters map[string]*rpcGetter
}

type rpcGetter struct {
	self string
}

func NewRpcContext(self string) *RpcContext {
	return &RpcContext{self: self}
}

func (r *RpcContext) Log(format string, a ...interface{}) {
	log.Printf("[server %s ] %s \n", r.self, fmt.Sprintf(format, a...))
}

func (r *RpcContext) Get(ctx context.Context, in *pb.Cache_Req) (*pb.Cache_Resp, error) {
	group := GetGroup(in.Group)
	if group == nil {
		return nil, errors.New("no such group:" + in.Group)
	}

	view, err := group.Get(in.Key)
	if err != nil {
		return nil, err
	}

	return &pb.Cache_Resp{Value: view.Bytes()}, nil
}

func (r *RpcContext) ServeRPC() error {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", r.self)
	if err != nil {
		return err
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return err
	}
	r.Log(" listening to running ... ")
	sev := grpc.NewServer(grpc.MaxConcurrentStreams(1024))
	pb.RegisterGroupCacheServer(sev, r)
	return sev.Serve(listener)
}

func (r *RpcContext) Set(peers ...string) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.peers = consistent.NewConsistent(replicas, nil)
	r.peers.Set(peers...)
	r.rpcGetters = make(map[string]*rpcGetter, len(peers))

	for _, peer := range peers {
		r.rpcGetters[peer] = &rpcGetter{self: peer}
	}
}

func (r *RpcContext) PickPeer(key string) (PeerGetter, bool) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if peer := r.peers.Get(key); peer != "" && peer != r.self {
		r.Log("pick peer %s ", peer)
		return r.rpcGetters[peer], true
	}
	return nil, false
}

func (rc *rpcGetter) Get(in *pb.Cache_Req, out *pb.Cache_Resp) (err error) {

	conn, err := grpc.Dial(rc.self, grpc.WithInsecure())
	if err != nil {
		return err
	}

	client := pb.NewGroupCacheClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	out, err = client.Get(ctx, in)
	fmt.Println("[rpc getter ] Get result --> ", out)
	return
}
