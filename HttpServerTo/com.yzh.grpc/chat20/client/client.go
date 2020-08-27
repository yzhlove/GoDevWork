package main

import (
	"WorkSpace/GoDevWork/HttpServerTo/com.yzh.grpc/chat20/pb"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

//////////////////////////////////////////////
// grpc  deadlines
//////////////////////////////////////////////

func main() {

	conn, err := grpc.Dial(":1234", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()

	client := pb.NewSearchServiceClient(conn)
	resp, err := client.Search(ctx, &pb.SearchRequest{Request: "gRPC"})
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				log.Fatalln("client.Search err:deadline")
			}
		}
		log.Fatalf("client.Search err: %v \n", err)
	}
	log.Printf("result resp:%s \n", resp.Response)

}
