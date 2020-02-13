package main

import (
	"WorkSpace/GoDevWork/HttpServerTo/com.yzh.grpc/chat14/proto"
	"context"
	"github.com/moby/moby/pkg/pubsub"
	"google.golang.org/grpc"
	"net"
	"strings"
	"time"
)

type PubSubService struct {
	pub *pubsub.Publisher
}

func NewPubSubService() *PubSubService {
	return &PubSubService{pub: pubsub.NewPublisher(100*time.Millisecond, 10),}
}

/*
Pub(context.Context, *String) (*String, error)
Sub(*String, PubSubServiceInterface_SubServer) error
*/

func (ps *PubSubService) Pub(c context.Context, request *proto.String) (*proto.String, error) {
	ps.pub.Publish(request.Var)
	return &proto.String{}, nil
}

func (ps *PubSubService) Sub(request *proto.String, stream proto.PubSubServiceInterface_SubServer) error {
	var channel chan interface{}
	if strings.TrimSpace(request.Var) == "" {
		channel = ps.pub.Subscribe()
	} else {
		channel = ps.pub.SubscribeTopic(func(v interface{}) bool {
			if key, ok := v.(string); ok {
				if strings.HasPrefix(key, request.Var) {
					return true
				}
			}
			return false
		})
	}
	//send resp
	for resp := range channel {
		if err := stream.Send(&proto.String{Var: resp.(string)}); err != nil {
			return err
		}
	}
	return nil
}

func main() {

	grpcServer := grpc.NewServer()
	proto.RegisterPubSubServiceInterfaceServer(grpcServer, NewPubSubService())

	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	grpcServer.Serve(l)

}
