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
	grpc_middle "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_tracer "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
	"go-kit-nine/agent/pb"
	"go-kit-nine/agent/src"
	"go-kit-nine/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	"reflect"
	"time"
)

type UserAgent struct {
	instance *etcdv3.Instancer
	logger   log.Logger
	tracer   opentracing.Tracer
}

func NewAgentClient(addr []string, logger log.Logger) (*UserAgent, error) {
	var ttl = time.Second * 5
	opts := etcdv3.ClientOptions{DialTimeout: ttl, DialKeepAlive: ttl}
	tracer, _, err := utils.NewJaegerTracer("user_agent_client")
	if err != nil {
		return nil, err
	}
	client, err := etcdv3.NewClient(context.Background(), addr, opts)
	if err != nil {
		return nil, err
	}
	instance, err := etcdv3.NewInstancer(client, "service.user.agent", logger)
	if err != nil {
		return nil, err
	}
	return &UserAgent{instance: instance, logger: logger, tracer: tracer}, nil
}

func (u *UserAgent) UserAgentClient() (src.Service, error) {
	endpointer := sd.NewEndpointer(u.instance, u.factoryFor(src.MakeLoginEndpoint), u.logger)
	balancer := lb.NewRoundRobin(endpointer)
	return src.EndpointService{
		LoginEndpoint: lb.Retry(3, 5*time.Second, balancer),
	}, nil
}

func (u *UserAgent) factoryFor(makeEndpoint func(service src.Service) endpoint.Endpoint) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		grpc_middle.ChainUnaryClient(
			grpc_tracer.UnaryClientInterceptor(grpc_tracer.WithTracer(u.tracer)),
			grpc_zap.UnaryClientInterceptor(utils.GetLog()),
		)
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
			md.Set(utils.CONTEXT_UID, UID)
			return metadata.NewOutgoingContext(ctx, *md)
		}),
	}
	return src.EndpointService{
		LoginEndpoint: grpc_transport.NewClient(
			conn, "pb.User", "RpcUserLogin",
			u.RequestLogin, u.ResponseLogin, pb.UserLogic_LoginAck{}, opts...,
		).Endpoint(),
	}
}

func (u *UserAgent) RequestLogin(ctx context.Context, req interface{}) (interface{}, error) {
	if in, ok := req.(*pb.UserLogic_Login); ok {
		return in, nil
	}
	return nil, errors.New("type err:" + reflect.TypeOf(req).String())
}

func (u *UserAgent) ResponseLogin(ctx context.Context, resp interface{}) (interface{}, error) {
	if out, ok := resp.(*pb.UserLogic_LoginAck); ok {
		return out, nil
	}
	return nil, errors.New("type err:" + reflect.TypeOf(resp).String())
}
