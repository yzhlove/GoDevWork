package api

import "net/rpc"

const HelloServiceName = "path/to/pkg.HelloService"

type HelloServiceInterface = interface {
	Hello(string, *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

type HelloServiceClient struct {
	*rpc.Client
}

//检查HelloServiceClient是否implement HelloServiceInterface
//如果没有实现，则编译报错
var _ HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	cli, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{cli}, nil
}

func (h *HelloServiceClient) Hello(request string, replay *string) error {
	return h.Call(HelloServiceName+".Hello", request, replay)
}
