package client

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	grpc_transport "github.com/go-kit/kit/transport/grpc"
	ept "go-kit-five/endpoint"
	"go-kit-five/pb"
	"go-kit-five/service"
	"go-kit-five/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func NewGrpcClient(conn *grpc.ClientConn, log *zap.Logger) service.Service {
	options := []grpc_transport.ClientOption{
		grpc_transport.ClientBefore(func(ctx context.Context, md *metadata.MD) context.Context {
			UID := utils.GetUID()
			log.Debug("set UID", zap.Any("UID", UID))
			md.Set(service.ContextReq, UID)
			ctx = metadata.NewOutgoingContext(context.Background(), *md)
			return ctx
		}),
	}
	var loginEndpoint endpoint.Endpoint
	{
		loginEndpoint = grpc_transport.NewClient(
			conn, "pb.User", "RpcUserLogin", ReqLogin, RespLogin, pb.LoginAck{}, options...,
		).Endpoint()
	}
	return ept.EndpointServer{
		LoginEndpoint: loginEndpoint,
	}
}

func ReqLogin(ctx context.Context, req interface{}) (interface{}, error) {
	if in, ok := req.(*pb.Login); ok {
		return in, nil
	}
	return nil, errors.New("type error")
}

func RespLogin(ctx context.Context, resp interface{}) (interface{}, error) {
	if out, ok := resp.(*pb.LoginAck); ok {
		return out, nil
	}
	return nil, errors.New("type error")
}
