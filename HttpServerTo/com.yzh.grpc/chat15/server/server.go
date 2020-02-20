package main

import (
	"WorkSpace/GoDevWork/HttpServerTo/com.yzh.grpc/chat15/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type PubSubService struct {
	pubsub *pubsubQueue
}

func NewPubSubService() *PubSubService {
	return &PubSubService{pubsub: New()}
}

func (p *PubSubService) Pub(_ context.Context, req *proto.String) (*proto.String, error) {
	p.pubsub.Pub(req.Var)
	return &proto.String{Var: "ok"}, nil
}

func (p *PubSubService) Sub(req *proto.String, stream proto.PubSubServiceInterface_SubServer) error {
	channel, ok := p.pubsub.Sub(req.Var)
	if !ok {
		channel <- "Welcome Tag+>" + req.Var
	}
	defer p.pubsub.CloseChan(req.Var)
	for msg := range channel {
		if str, ok := msg.(string); ok {
			if err := stream.Send(&proto.String{Var: str}); err != nil {
				fmt.Println("stream send Err:", err)
				return err
			}
		}
	}
	return nil
}

func main() {
	server := grpc.NewServer()
	proto.RegisterPubSubServiceInterfaceServer(server, NewPubSubService())
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	server.Serve(l)
}
