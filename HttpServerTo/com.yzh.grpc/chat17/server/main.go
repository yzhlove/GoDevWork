package main

import (
	"WorkSpace/GoDevWork/HttpServerTo/com.yzh.grpc/chat17/proto"
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type PubSub struct {
	_pubsub *Queue
}

func NewService() *PubSub {
	return &PubSub{NewQueue()}
}

func (p *PubSub) Pub(c context.Context, req *proto.Manager_Msg) (*proto.Manager_Nil, error) {
	p._pubsub.Pub(req)
	return &proto.Manager_Nil{}, nil
}

func (p *PubSub) Sub(req *proto.Manager_Zone, stream proto.PubSubServiceInterface_SubServer) error {
	manager := p._pubsub.Sub(req.Var)
	ctx := stream.Context()
	for {
		select {
		case msg := <-manager.Ch:
			if m, ok := msg.(*proto.Manager_Msg); ok {
				if err := stream.Send(m); err != nil {
					return err
				}
			}
		case <-ctx.Done():
			p._pubsub.Close(manager)
			return nil
		}
	}
}

func main() {

	server := grpc.NewServer()
	proto.RegisterPubSubServiceInterfaceServer(server, NewService())
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	logrus.Info("start server :", l.Addr().String())
	server.Serve(l)
}
