package src

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"go-kit-nine/agent/pb"
	"golang.org/x/time/rate"
	"reflect"
)

type EndpointService struct {
	LoginEndpoint endpoint.Endpoint
}

func NewLoginEndpoint(service Service, l *rate.Limiter, tracer opentracing.Tracer) EndpointService {
	var loginEndpoint endpoint.Endpoint
	{
		loginEndpoint = MakeLoginEndpoint(service)
		loginEndpoint = NewGoRateMiddle(l)(loginEndpoint)
		loginEndpoint = NewTracerEndpointMiddle(tracer)(loginEndpoint)
	}
	return EndpointService{LoginEndpoint: loginEndpoint}
}

func (es EndpointService) Login(ctx context.Context, in *pb.UserLogic_Login) (*pb.UserLogic_LoginAck, error) {
	if result, err := es.LoginEndpoint(ctx, in); err != nil {
		return nil, err
	} else if out, ok := result.(*pb.UserLogic_LoginAck); ok {
		return out, nil
	} else {
		return nil, errors.New("type err:" + reflect.TypeOf(result).String())
	}
}

func MakeLoginEndpoint(service Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		if out, ok := request.(*pb.UserLogic_Login); ok {
			return out, nil
		}
		return nil, errors.New("type err:" + reflect.TypeOf(request).String())
	}
}
