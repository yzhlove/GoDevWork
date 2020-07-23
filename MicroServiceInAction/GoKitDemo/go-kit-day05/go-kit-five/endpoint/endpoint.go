package endpoint

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"go-kit-five/pb"
	"go-kit-five/service"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

type EndpointServer struct {
	LoginEndpoint endpoint.Endpoint
}

func NewEndpoint(s service.Service, log *zap.Logger, limit *rate.Limiter) EndpointServer {
	var e endpoint.Endpoint
	{
		e = MakeLoginEndpoint(s)
		e = LogMiddle(log)(e)
		e = GolangRateMiddle(limit)(e)
	}
	//e = GolangRateMiddle(limit)(LogMiddle(log)(MakeLoginEndpoint(s)))
	return EndpointServer{LoginEndpoint: e}
}

func (e EndpointServer) Login(ctx context.Context, in *pb.Login) (*pb.LoginAck, error) {
	resp, err := e.LoginEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	if ack, ok := resp.(*pb.LoginAck); ok {
		return ack, nil
	}
	return nil, errors.New("login ack type error")
}

func MakeLoginEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		if in, ok := request.(*pb.Login); ok {
			return s.Login(ctx, in)
		}
		return nil, errors.New("endpoint login type error")
	}
}
