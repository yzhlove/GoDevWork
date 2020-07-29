package client

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/etcdv3"
	"github.com/go-kit/kit/sd/lb"
	grpc_transport "github.com/go-kit/kit/transport/grpc"
	"go-kit-seven/agent/pb"
	"go-kit-seven/agent/src"
	"go-kit-seven/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	"reflect"
	"time"
)

type UserAgent struct {
	instance *etcdv3.Instancer
	logger   log.Logger
}

func NewAgentClient(addr []string, logger log.Logger) (*UserAgent, error) {
	var (
		etcd_addr = addr
		svcName   = "svc.user.agent"
		ttl       = 5 * time.Second
	)
	opts := etcdv3.ClientOptions{
		DialTimeout:   ttl,
		DialKeepAlive: ttl,
	}
	client, err := etcdv3.NewClient(context.Background(), etcd_addr, opts)
	if err != nil {
		return nil, err
	}
	instance, err := etcdv3.NewInstancer(client, svcName, logger)
	if err != nil {
		return nil, err
	}
	return &UserAgent{instance: instance, logger: logger}, nil
}

func (u *UserAgent) UserAgentClient() (src.Service, error) {
	var (
		retryMax     = 3
		retryTimeout = 5 * time.Second
		endpoints    src.EndpointServer
	)
	factory := u.factoryFor(src.MakeLoginPoint)
	endpointer := sd.NewEndpointer(u.instance, factory, u.logger)
	blancer := lb.NewRoundRobin(endpointer)
	retry := lb.Retry(retryMax, retryTimeout, blancer)
	endpoints.LoginEndpoint = retry
	return endpoints, nil
}

func (u *UserAgent) factoryFor(makeEndpoint func(s src.Service) endpoint.Endpoint) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		conn, err := grpc.Dial(instance, grpc.WithInsecure())
		if err != nil {
			return nil, nil, err
		}
		svc := u.NewGrpcClient(conn)
		endpoints := makeEndpoint(svc)
		return endpoints, conn, nil
	}
}

func (u *UserAgent) NewGrpcClient(conn *grpc.ClientConn) src.Service {
	opts := []grpc_transport.ClientOption{
		grpc_transport.ClientBefore(func(ctx context.Context, md *metadata.MD) context.Context {
			UID := utils.GetUID()
			md.Set(src.CONTEXT_UID, UID)
			return metadata.NewOutgoingContext(ctx, *md)
		}),
	}
	return src.EndpointServer{LoginEndpoint: grpc_transport.NewClient(
		conn, "pb.User", "RpcUserLogin", u.Request, u.Response, pb.LoginAck{}, opts...,
	).Endpoint()}
}

func (u *UserAgent) Request(ctx context.Context, in interface{}) (interface{}, error) {
	if req, ok := in.(*pb.Login); ok {
		return req, nil
	}
	return nil, errors.New("type err:" + reflect.TypeOf(in).Name())
}

func (u *UserAgent) Response(ctx context.Context, in interface{}) (interface{}, error) {
	if resp, ok := in.(*pb.LoginAck); ok {
		return resp, nil
	}
	return nil, errors.New("type err:" + reflect.TypeOf(in).Name())
}
