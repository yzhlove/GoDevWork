package src

import (
	"context"
	"errors"
	grpc_transport "github.com/go-kit/kit/transport/grpc"
	"go-kit-six/agent/pb"
	"go-kit-six/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
	"reflect"
)

type Grpc struct {
	login grpc_transport.Handler
}

func NewGrpc(e LoginEndpoint, log *zap.Logger) pb.UserServer {
	opts := []grpc_transport.ServerOption{
		grpc_transport.ServerBefore(func(ctx context.Context, md metadata.MD) context.Context {
			ctx = context.WithValue(ctx, CONTEXT_REQ_UID, md.Get(CONTEXT_REQ_UID))
			return ctx
		}),
		grpc_transport.ServerErrorHandler(utils.NewZapErrorHandle(log)),
	}
	return &Grpc{login: grpc_transport.NewServer(
		e.LoginPoint, request, response, opts...)}
}

func (s *Grpc) RpcUserLogin(ctx context.Context, req *pb.Login) (*pb.LoginAck, error) {
	_, resp, err := s.login.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	if ack, ok := resp.(*pb.LoginAck); ok {
		return ack, err
	}
	return nil, errors.New("transport.RpcUserLogin type err:" + reflect.TypeOf(resp).Name())
}

func request(ctx context.Context, in interface{}) (interface{}, error) {
	if req, ok := in.(*pb.Login); ok {
		return req, nil
	}
	return nil, errors.New("transport.request type err:" + reflect.TypeOf(in).Name())
}

func response(ctx context.Context, in interface{}) (interface{}, error) {
	if resp, ok := in.(*pb.LoginAck); ok {
		return resp, nil
	}
	return nil, errors.New("transport.response type err:" + reflect.TypeOf(in).Name())
}
