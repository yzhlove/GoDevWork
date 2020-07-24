package src

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"go-kit-six/agent/pb"
	"golang.org/x/time/rate"
	"reflect"
)

type LoginEndpoint struct {
	LoginPoint endpoint.Endpoint
}

func NewLoginEndpoint(s Service, limit *rate.Limiter) LoginEndpoint {
	//var login endpoint.Endpoint
	//{
	//	login = MakeLoginPoint(s)
	//	login = GolangRateMiddle(limit)(login)
	//}
	return LoginEndpoint{LoginPoint: GolangRateMiddle(limit)(MakeLoginPoint(s))}
}

func (l LoginEndpoint) Login(ctx context.Context, in *pb.Login) (*pb.LoginAck, error) {
	result, err := l.LoginPoint(ctx, in)
	if err != nil {
		return nil, err
	}
	if out, ok := result.(*pb.LoginAck); ok {
		return out, nil
	}
	return nil, errors.New("endpoint.log.login type err:" + reflect.TypeOf(result).Name())
}

func MakeLoginPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		if req, ok := request.(*pb.Login); ok {
			return s.Login(ctx, req)
		}
		err = errors.New("endpoint.log.make type err:" + reflect.TypeOf(request).Name())
		return
	}
}
