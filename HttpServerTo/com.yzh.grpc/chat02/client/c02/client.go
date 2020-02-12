package main

import (
	"WorkSpace/GoDevWork/HttpServerTo/com.yzh.grpc/chat02/conf"
	"net/rpc"
)

type HelloServiceClient struct {
	*rpc.Client
}

/*
var _ HelloServiceInterface = (*HelloServiceClient)(nil)
确保HelloServiceClient 实现了 HelloServiceInterface
如果没有实现，则编译报错
*/

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	cli, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{cli}, nil
}

func (h *HelloServiceClient) Hello(request string, replay *string) error {
	return h.Call(conf.HelloServiceName+".Hello", request, replay)
}
