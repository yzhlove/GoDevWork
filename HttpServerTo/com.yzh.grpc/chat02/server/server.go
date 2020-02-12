package main

import (
	"WorkSpace/GoDevWork/HttpServerTo/com.yzh.grpc/chat02/conf"
	"net/rpc"
)

type HelloServiceInterface = interface {
	Hello(request string, replay *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(conf.HelloServiceName, svc)
}

