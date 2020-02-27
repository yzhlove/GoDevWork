package main

import (
	"WorkSpace/GoDevWork/HttpServerTo/com.yzh.grpc/chat18/proto"
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"io"
	"net"
)

type PubSub struct {
	_pubsub *Queue
}

func NewService() *PubSub {
	return &PubSub{_pubsub: NewQueue()}
}

func (p *PubSub) Pub(c context.Context, req *proto.Manager_Msg) (*proto.Manager_Nil, error) {
	p._pubsub.Pub(req)
	return &proto.Manager_Nil{}, nil
}

func (p *PubSub) Sub(req *proto.Manager_Zone, stream proto.PubSubServiceInterface_SubServer) error {
	c := p._pubsub.Sub(req.Var)
	ctx := stream.Context()
	for {
		select {
		case msg := <-c.MsgCh:
			if m, ok := msg.(*proto.Manager_Msg); ok {
				if err := stream.Send(m); err != nil {
					if errors.Is(err, io.EOF) {
						return nil
					}
					return err
				}
			}
		case <-ctx.Done():
			p._pubsub.Close(c)
			return nil
		}
	}
}

func main() {
	serv := grpc.NewServer()
	proto.RegisterPubSubServiceInterfaceServer(serv, NewService())
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	logrus.Info("listen:", l.Addr().String())
	serv.Serve(l)
}
