package src

import (
	"context"
	"errors"
	grpc_transport "github.com/go-kit/kit/transport/grpc"
	"go-kit-nine/agent/pb"
	"go-kit-nine/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
	"reflect"
)

type Grpc struct {
	grpc_transport.Handler
}

func NewGrpc(ept EndpointService, log *zap.Logger) pb.UserServer {
	opts := []grpc_transport.ServerOption{
		grpc_transport.ServerBefore(func(ctx context.Context, md metadata.MD) context.Context {
			return context.WithValue(ctx, utils.CONTEXT_UID, md.Get(utils.CONTEXT_UID))
		}),
		grpc_transport.ServerErrorHandler(utils.NewZapLogErrHandle(log)),
	}
	return &Grpc{grpc_transport.NewServer(ept.LoginEndpoint,
		RequestLogin, ResponseLogin, opts...)}
}

func (s *Grpc) RpcUserLogin(ctx context.Context, req *pb.UserLogic_Login) (*pb.UserLogic_LoginAck, error) {
	_, result, err := s.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	if out, ok := result.(*pb.UserLogic_LoginAck); ok {
		return out, nil
	}
	return nil, errors.New("type err:" + reflect.TypeOf(result).String())
}

func RequestLogin(ctx context.Context, req interface{}) (interface{}, error) {
	if in, ok := req.(*pb.UserLogic_Login); ok {
		return in, nil
	}
	return nil, errors.New("type err:" + reflect.TypeOf(req).String())
}

func ResponseLogin(ctx context.Context, resp interface{}) (interface{}, error) {
	if out, ok := resp.(*pb.UserLogic_LoginAck); ok {
		return out, nil
	}
	return nil, errors.New("type err:" + reflect.TypeOf(resp).String())
}
