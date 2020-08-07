package src

import (
	"context"
	"errors"
	grpc_transport "github.com/go-kit/kit/transport/grpc"
	"go-kit-eight/agent/pb"
	"go-kit-eight/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
	"reflect"
)

type grpc struct {
	login grpc_transport.Handler
}

func NewGrpc(e EndpointServer, log *zap.Logger) pb.UserServer {
	opts := []grpc_transport.ServerOption{
		grpc_transport.ServerBefore(func(ctx context.Context, md metadata.MD) context.Context {
			return context.WithValue(ctx, CONTEXT_UID, md.Get(CONTEXT_UID))
		}),
		grpc_transport.ServerErrorHandler(utils.NewZapErrorHandle(log)),
	}
	return &grpc{
		login: grpc_transport.NewServer(
			e.LoginEndpoint,
			RequestLogin, ResponseLogin,
			opts...),
	}
}

func (s grpc) RpcUserLogin(ctx context.Context, in *pb.Login) (*pb.LoginAck, error) {
	_, result, err := s.login.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	if out, ok := result.(*pb.LoginAck); ok {
		return out, nil
	}
	return nil, errors.New("type err:" + reflect.TypeOf(result).String())
}

func RequestLogin(ctx context.Context, req interface{}) (interface{}, error) {
	if in, ok := req.(*pb.Login); ok {
		return in, nil
	}
	return nil, errors.New("type err:" + reflect.TypeOf(req).String())

}

func ResponseLogin(ctx context.Context, resp interface{}) (interface{}, error) {
	if out, ok := resp.(*pb.LoginAck); ok {
		return out, nil
	}
	return nil, errors.New("type err:" + reflect.TypeOf(resp).String())
}
