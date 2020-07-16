package endpoint

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"go-kit-two/service"
	"go-kit-two/utils"
	"go.uber.org/zap"
	"reflect"
)

type EndpointServer struct {
	AddEndpoint endpoint.Endpoint
}

func NewEndpointServer(sev service.Service, log *zap.Logger) EndpointServer {
	var addEndpoint endpoint.Endpoint
	addEndpoint = MakeAddEndpointServer(sev)
	addEndpoint = LoggerMiddleware(log)(addEndpoint)
	return EndpointServer{
		AddEndpoint: addEndpoint,
	}
}

func MakeAddEndpointServer(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		utils.GetLog().Debug("request type => ", zap.String("requestType", reflect.TypeOf(request).String()))
		if req, ok := request.(service.Add); ok {
			response = s.TestAdd(ctx, req)
			return
		}
		err = errors.New("type err")
		utils.GetLog().Debug("transform  err", zap.Error(err),
			zap.Any("response", response))
		return
	}
}
