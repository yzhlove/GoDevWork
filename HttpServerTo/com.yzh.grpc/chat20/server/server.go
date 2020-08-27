package main

//////////////////////////////////////////////
// grpc  deadlines
//////////////////////////////////////////////

import (
	"WorkSpace/GoDevWork/HttpServerTo/com.yzh.grpc/chat20/pb"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"time"
)

type SearchService struct{}

func (s *SearchService) Search(ctx context.Context, in *pb.SearchRequest) (*pb.SearchResponse, error) {
	for i := 0; i < 5; i++ {
		if ctx.Err() == context.Canceled {
			return nil, status.Error(codes.Canceled, "SearchService.Search canceled")
		}
		time.Sleep(time.Second)
	}
	return &pb.SearchResponse{Response: in.Request + " by Server(--"}, nil
}

func main() {

	server := grpc.NewServer()
	pb.RegisterSearchServiceServer(server, new(SearchService))
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	log.Fatal(server.Serve(lis))
}
