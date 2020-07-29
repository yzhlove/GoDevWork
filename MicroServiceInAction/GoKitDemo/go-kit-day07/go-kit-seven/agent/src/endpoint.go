package src

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"go-kit-seven/agent/pb"
	"golang.org/x/time/rate"
	"reflect"
)

type EndpointServer struct {
	LoginEndpoint endpoint.Endpoint
}

func NewEndpointServer(s Service, l *rate.Limiter) EndpointServer {
	var login endpoint.Endpoint
	{
		login = MakeLoginPoint(s)
		login = GolangRateMiddle(l)(login)
	}
	return EndpointServer{LoginEndpoint: login}
}

func (e EndpointServer) Login(ctx context.Context, in *pb.Login) (*pb.LoginAck, error) {
	if result, err := e.LoginEndpoint(ctx, in); err != nil {
		return nil, err
	} else if out, ok := result.(*pb.LoginAck); ok {
		return out, nil
	} else {
		return nil, errors.New("type err:" + reflect.TypeOf(result).Name())
	}
}

func MakeLoginPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		if in, ok := request.(*pb.Login); ok {
			return s.Login(ctx, in)
		}
		return nil, errors.New("type err:" + reflect.TypeOf(request).Name())
	}
}
