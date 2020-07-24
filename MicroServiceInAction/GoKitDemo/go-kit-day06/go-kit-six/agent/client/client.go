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
	"go-kit-six/agent/pb"
	"go-kit-six/agent/src"
	"go-kit-six/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	"reflect"
	"time"
)

type endpointFunc func(s src.Service) endpoint.Endpoint

type Agent struct {
	instance *etcdv3.Instancer
	log      log.Logger
}

func NewAgentClient(addr []string, logger log.Logger) (*Agent, error) {
	var (
		ttl    = 5 * time.Second
		prefix = "/register/svc.user.agent"
	)

	opts := etcdv3.ClientOptions{
		DialTimeout:   ttl,
		DialKeepAlive: ttl,
	}
	cli, err := etcdv3.NewClient(context.Background(), addr, opts)
	if err != nil {
		return nil, err
	}
	instance, err := etcdv3.NewInstancer(cli, prefix, logger)
	if err != nil {
		return nil, err
	}
	return &Agent{instance: instance, log: logger}, nil
}

func (agent *Agent) AgentClient() (src.Service, error) {
	var (
		retryMax      = 3
		retryTimeout  = 500 * time.Millisecond
		loginEndpoint src.LoginEndpoint
	)
	factory := agent.factoryFor(src.MakeLoginPoint)
	ept := sd.NewEndpointer(agent.instance, factory, agent.log)
	balancer := lb.NewRandom(ept, time.Now().UnixNano())
	retry := lb.Retry(retryMax, retryTimeout, balancer)
	loginEndpoint.LoginPoint = retry
	return loginEndpoint, nil
}

func (agent *Agent) factoryFor(makeEndpoint endpointFunc) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		conn, err := grpc.Dial(instance, grpc.WithInsecure())
		if err != nil {
			return nil, nil, err
		}
		ept := makeEndpoint(agent.NewGrpcClient(conn))
		return ept, conn, nil
	}
}

func (agent *Agent) NewGrpcClient(conn *grpc.ClientConn) src.Service {
	opts := []grpc_transport.ClientOption{
		grpc_transport.ClientBefore(func(ctx context.Context, md *metadata.MD) context.Context {
			UID := utils.GetUID()
			md.Set(src.CONTEXT_REQ_UID, UID)
			ctx = metadata.NewOutgoingContext(context.Background(), *md)
			return ctx
		}),
	}
	var loginEndpoint endpoint.Endpoint
	{
		loginEndpoint = grpc_transport.NewClient(
			conn, "pb.User", "RpcUserLogin", agent.request, agent.response, pb.LoginAck{}, opts...,
		).Endpoint()
	}
	return src.LoginEndpoint{LoginPoint: loginEndpoint}
}

func (agent *Agent) request(ctx context.Context, req interface{}) (interface{}, error) {
	if in, ok := req.(*pb.Login); ok {
		return in, nil
	}
	return nil, errors.New("client.request type err:" + reflect.TypeOf(req).Name())
}

func (agent *Agent) response(ctx context.Context, resp interface{}) (interface{}, error) {
	if out, ok := resp.(*pb.LoginAck); ok {
		return out, nil
	}
	return nil, errors.New("client.response type err:" + reflect.TypeOf(resp).Name())
}
