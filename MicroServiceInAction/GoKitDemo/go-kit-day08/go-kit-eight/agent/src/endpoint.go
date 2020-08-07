package src

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"go-kit-eight/agent/pb"
	"golang.org/x/time/rate"
	"reflect"
)

type EndpointServer struct {
	LoginEndpoint endpoint.Endpoint
}

func NewEndpointServer(s Service, l *rate.Limiter) EndpointServer {
	var login endpoint.Endpoint
	{
		login = MakeLoginEndpoint(s)
		//去除限流组件
		//login = NewGoRateMid(l)(login)
	}
	return EndpointServer{LoginEndpoint: login}
}

func (e EndpointServer) Login(ctx context.Context, in *pb.Login) (*pb.LoginAck, error) {
	if resp, err := e.LoginEndpoint(ctx, in); err != nil {
		return nil, err
	} else if out, ok := resp.(*pb.LoginAck); ok {
		return out, nil
	} else {
		return nil, errors.New("type err:" + reflect.TypeOf(resp).String())
	}
}

func MakeLoginEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		if req, ok := request.(*pb.Login); ok {
			return s.Login(ctx, req)
		}
		return nil, errors.New("type err:" + reflect.TypeOf(request).String())
	}
}
