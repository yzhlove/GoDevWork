package main

import (
	"WorkSpace/GoDevWork/MicroServiceInAction/JaegerDemo/chat_grpc/example_interceptor_02/pb"
	"context"
	"google.golang.org/grpc"
	"log"
	"strconv"
	"time"
)

const addr = ":1234"

func main() {

	conn, err := grpc.Dial(addr, grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(GenerateClientInterceptor),
		grpc.WithStreamInterceptor(StreamClientInterceptor))

	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewHelloClient(conn)
	if resp, err := client.SayHello(context.Background(), &pb.Req{In: "in client"}); err != nil {
		panic(err)
	} else {
		log.Printf("resp %v \n", resp.Out)
	}

	stream, err := client.Take(context.Background())
	if err != nil {
		panic(err)
	}

	go func() {
		var count int
		for range time.NewTicker(time.Second * 2).C {
			count++
			if err := stream.Send(&pb.Req{In: "client -x- + " + strconv.Itoa(count)}); err != nil {
				log.Printf("client send err: %+v \n", err)
				return
			}
		}
	}()

	for {
		if resp, err := stream.Recv(); err != nil {
			panic(err)
		} else {
			log.Printf("stream recv :%+v \n", resp.Out)
		}
	}
}

func GenerateClientInterceptor(ctx context.Context, method string, req, replay interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	log.Printf("before generate . method:%+v ,request:%+v \n", method, req)
	err := invoker(ctx, method, req, replay, cc, opts...)
	log.Printf("after generate. replay:%+v err:%+v\n", replay, err)
	return err
}

func StreamClientInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn,
	method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	log.Printf("before stream. method:%+v , desc: %+v \n", method, desc)
	clientStream, err := streamer(ctx, desc, cc, method, opts...)
	log.Printf("after stream. method:%+v err:%+v \n", method, err)
	return clientStream, err
}
