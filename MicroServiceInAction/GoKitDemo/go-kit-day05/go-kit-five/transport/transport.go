package transport

import (
	"context"
	"errors"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	ept "go-kit-five/endpoint"
	"go-kit-five/pb"
	"go-kit-five/service"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"google.golang.org/grpc/metadata"
)

type server struct {
	login grpctransport.Handler
}

func NewGrpcServer(e ept.EndpointServer, log *zap.Logger, limit *rate.Limiter) pb.UserServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerBefore(func(ctx context.Context, md metadata.MD) context.Context {
			ctx = context.WithValue(ctx, service.ContextReq, md.Get(service.ContextReq))
			return ctx
		}),
		grpctransport.ServerErrorHandler(NewZapHandle(log)),
	}
	return &server{login: grpctransport.NewServer(
		e.LoginEndpoint,
		DecodeReqFunc,
		EncodeRespFunc,
		options...,
	)}
}

func (s server) RpcUserLogin(ctx context.Context, req *pb.Login) (*pb.LoginAck, error) {
	_, resp, err := s.login.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	if result, ok := resp.(*pb.LoginAck); ok {
		return result, nil
	}
	return nil, errors.New("rpc_user_login type error")
}

func DecodeReqFunc(ctx context.Context, in interface{}) (interface{}, error) {
	if req, ok := in.(*pb.Login); ok {
		return req, nil
	}
	return nil, errors.New("decode req func type error")
}

func EncodeRespFunc(ctx context.Context, in interface{}) (interface{}, error) {
	if resp, ok := in.(*pb.LoginAck); ok {
		return resp, nil
	}
	return nil, errors.New("encode resp func type error")
}
