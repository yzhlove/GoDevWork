package main

import (
	"WorkSpace/GoDevWork/HttpServerTo/com.yzh.grpc/chat16/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type PubSubService struct {
	_pubsub *pubsub
}

func NewService() *PubSubService {
	return &PubSubService{_pubsub: New()}
}

func (p *PubSubService) Pub(c context.Context, req *proto.Manager_Msg) (*proto.Manager_Nil, error) {
	p._pubsub.Pub(req)
	return &proto.Manager_Nil{}, nil
}

func (p *PubSubService) Sub(req *proto.Manager_Zone, stream proto.PubSubServiceInterface_SubServer) error {
	/*channel := p._pubsub.Sub(req.Var)
	defer p._pubsub.Close(channel)
	for msg := range channel {
		if m, ok := msg.(*proto.Manager_Msg); ok {
			if err := stream.Send(m); err != nil {
				return err
			}
		}
	}*/
	channel := p._pubsub.Sub(req.Var)
	ctx := stream.Context()
	for {
		select {
		case msg := <-channel:
			if m, ok := msg.(*proto.Manager_Msg); ok {
				if err := stream.Send(m); err != nil {
					fmt.Println("close channel by stream")
					p._pubsub.Close(channel)
					return err
				}
			}
		case <-ctx.Done():
			fmt.Println("Done .....")
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
	server.Serve(l)
}
