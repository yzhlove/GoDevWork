package src

import (
	"context"
	"errors"
	grpc_transport "github.com/go-kit/kit/transport/grpc"
	"go-kit-seven/agent/pb"
	"go-kit-seven/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
	"reflect"
)

type GrpcServer struct {
	login grpc_transport.Handler
}

func NewGrpcServer(endpoint EndpointServer, log *zap.Logger) pb.UserServer {
	opts := []grpc_transport.ServerOption{
		grpc_transport.ServerBefore(func(ctx context.Context, md metadata.MD) context.Context {
			ctx = context.WithValue(ctx, CONTEXT_UID, md.Get(CONTEXT_UID))
			return ctx
		}),
		grpc_transport.ServerErrorHandler(utils.NewZapErrorHandle(log)),
	}
	return &GrpcServer{login: grpc_transport.NewServer(
		endpoint.LoginEndpoint,
		Request, Response, opts...)}
}

func (s *GrpcServer) RpcUserLogin(ctx context.Context, in *pb.Login) (*pb.LoginAck, error) {
	if _, resp, err := s.login.ServeGRPC(ctx, in); err != nil {
		return nil, err
	} else if out, ok := resp.(*pb.LoginAck); ok {
		return out, nil
	} else {
		return nil, errors.New("type err:" + reflect.TypeOf(resp).Name())
	}
}

func Request(ctx context.Context, req interface{}) (interface{}, error) {
	if out, ok := req.(*pb.Login); ok {
		return out, nil
	}
	return nil, errors.New("type err:" + reflect.TypeOf(req).Name())
}

func Response(ctx context.Context, req interface{}) (out interface{}, err error) {
	if out, ok := req.(*pb.LoginAck); ok {
		return out, nil
	}
	return nil, errors.New("type err:" + reflect.TypeOf(req).Name())
}
