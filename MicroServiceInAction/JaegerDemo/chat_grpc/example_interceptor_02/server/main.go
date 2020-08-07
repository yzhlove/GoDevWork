package main

import (
	"WorkSpace/GoDevWork/MicroServiceInAction/JaegerDemo/chat_grpc/example_interceptor_02/pb"
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"strconv"
	"time"
)

const addr = ":1234"

type server struct{}

func (s server) SayHello(ctx context.Context, in *pb.Req) (*pb.Resp, error) {
	log.Printf("client In :%+v \n", in.In)
	return &pb.Resp{Out: "hello client -x-"}, nil
}

func (s server) Take(stream pb.Hello_TakeServer) error {
	ctx := stream.Context()
	go func() {
		for {
			if result, err := stream.Recv(); err != nil {
				log.Printf("take recv err:%+v \n", err)
				return
			} else {
				log.Printf("client message: %+v \n", result.In)
			}
		}
	}()
	var count int
	t := time.NewTicker(time.Second)
	for {
		count++
		select {
		case <-ctx.Done():
			return errors.New("client is close ")
		case <-t.C:
			if err := stream.Send(&pb.Resp{Out: "hello client:" + strconv.Itoa(count)}); err != nil {
				return err
			}
		}
	}
}

func main() {

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(GenerateServerIntercept),
		grpc.StreamInterceptor(StreamServerIntercept))
	pb.RegisterHelloServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(listen); err != nil {
		panic(err)
	}
}

func GenerateServerIntercept(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("before generate intercept. info: %+v \n", info)
	resp, err := handler(ctx, req)
	log.Printf("after generate intercept. resp:%+v\n", resp)
	return resp, err
}

func StreamServerIntercept(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Printf("before stream intercept. Info:%+v \n", info)
	err := handler(srv, ss)
	log.Printf("after stream intercept. err:%+v \n", err)
	return err
}
