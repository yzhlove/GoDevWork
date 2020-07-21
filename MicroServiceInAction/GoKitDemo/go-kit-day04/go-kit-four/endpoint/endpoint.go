package endpoint

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"go-kit-four/service"
	"go.uber.org/ratelimit"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"reflect"
)

type Endpoint struct {
	Add   endpoint.Endpoint
	Login endpoint.Endpoint
}

func NewEndpoint(s service.Service, log *zap.Logger, stdLimit *rate.Limiter, uberLimit ratelimit.Limiter) Endpoint {
	var add endpoint.Endpoint
	{
		add = MakeAddEndpoint(s)
		add = LoginMiddle(log)(add)
		add = AuthMiddle(log)(add)
		add = UberRateMiddle(uberLimit)(add)
	}

	var login endpoint.Endpoint
	{
		login = MakeLoginEndpoint(s)
		login = LoginMiddle(log)(login)
		login = GolangRateAllowMiddle(stdLimit)(login)
	}

	return Endpoint{Add: add, Login: login}
}

func MakeAddEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		if in, ok := request.(service.Add); ok {
			return s.TestAdd(ctx, in), nil
		}
		return nil, errors.New("make add endpoint type error:" + reflect.TypeOf(request).Name())
	}
}

func MakeLoginEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		if in, ok := request.(service.Login); ok {
			return s.Login(ctx, in)
		}
		return nil, errors.New("make login endpoint type error:" + reflect.TypeOf(request).Name())
	}
}
