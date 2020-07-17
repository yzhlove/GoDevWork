package endpoint

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"go-kit-three/service"
	"go.uber.org/zap"
	"reflect"
)

type Server struct {
	Add   endpoint.Endpoint
	Login endpoint.Endpoint
}

func NewEndpointServer(s service.Service, log *zap.Logger) Server {
	var add endpoint.Endpoint
	{
		add = MakeAddPoint(s)
		add = LoggerDecorator(log)(add)
		add = AuthDecorator(log)(add)
	}
	var login endpoint.Endpoint
	{
		login = MakeLoginPoint(s)
		login = LoggerDecorator(log)(login)
	}
	return Server{Add: add, Login: login}
}

func MakeAddPoint(s service.Service) endpoint.Endpoint {
	point := func(ctx context.Context, req interface{}) (resp interface{}, err error) {
		if in, ok := req.(service.Add); ok {
			return s.TestAdd(ctx, in), nil
		}
		return nil, errors.New("add point type err:" + reflect.TypeOf(req).Name())
	}
	return point
}

func MakeLoginPoint(s service.Service) endpoint.Endpoint {
	point := func(ctx context.Context, req interface{}) (resp interface{}, err error) {
		if in, ok := req.(service.Login); ok {
			ack, err := s.Login(ctx, in)
			return ack, err
		}
		return nil, errors.New("login point type err:" + reflect.TypeOf(req).Name())
	}
	return point
}
