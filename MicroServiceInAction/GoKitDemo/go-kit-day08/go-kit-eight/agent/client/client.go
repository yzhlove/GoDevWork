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
	"go-kit-eight/agent/pb"
	"go-kit-eight/agent/src"
	"go-kit-eight/utils"
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

func NewUserAgentClient(addr []string, logger log.Logger) (*UserAgent, error) {
	server_name := "server.user.agent"
	ttl := 5 * time.Second

	opts := etcdv3.ClientOptions{DialKeepAlive: ttl, DialTimeout: ttl}
	client, err := etcdv3.NewClient(context.Background(), addr, opts)
	if err != nil {
		return nil, err
	}
	instance, err := etcdv3.NewInstancer(client, server_name, logger)
	if err != nil {
		return nil, err
	}
	return &UserAgent{instance: instance, logger: logger}, nil
}

func (u *UserAgent) UserAgentClient() (src.Service, error) {
	retryMax := 3
	retryTimeout := 5 * time.Second

	factory := u.factoryFor(src.MakeLoginEndpoint)
	pointer := sd.NewEndpointer(u.instance, factory, u.logger)
	balancer := lb.NewRoundRobin(pointer)
	retry := lb.Retry(retryMax, retryTimeout, balancer)
	return src.EndpointServer{LoginEndpoint: retry}, nil
}

func (u *UserAgent) factoryFor(makeEndpoint func(s src.Service) endpoint.Endpoint) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		conn, err := grpc.Dial(instance, grpc.WithInsecure())
		if err != nil {
			return nil, nil, err
		}
		return makeEndpoint(u.NewGrpcClient(conn)), conn, nil
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
		conn,
		"pb.User",
		"RpcUserLogin",
		u.RequestLogin,
		u.ResponseLogin,
		pb.LoginAck{},
		opts...,
	).Endpoint()}
}

func (u *UserAgent) RequestLogin(ctx context.Context, req interface{}) (interface{}, error) {
	if in, ok := req.(*pb.Login); ok {
		return in, nil
	}
	return nil, errors.New("type err:" + reflect.TypeOf(req).String())
}

func (u *UserAgent) ResponseLogin(ctx context.Context, resp interface{}) (interface{}, error) {
	if out, ok := resp.(*pb.LoginAck); ok {
		return out, nil
	}
	return nil, errors.New("type err:" + reflect.TypeOf(resp).String())
}
