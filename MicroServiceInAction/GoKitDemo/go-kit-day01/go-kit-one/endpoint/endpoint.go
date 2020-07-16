package endpoint

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"go-kit-one/service"
)

type EndpointServer struct {
	AddEndpoint endpoint.Endpoint
}

func NewEndpointServer(svc service.Service) EndpointServer {
	return EndpointServer{AddEndpoint: MakeAddEndpoint(svc)}
}

func MakeAddEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		if req, ok := request.(service.Add); ok {
			response = svc.TestAdd(ctx, req)
			return
		}
		err = errors.New("type err")
		return
	}
}
