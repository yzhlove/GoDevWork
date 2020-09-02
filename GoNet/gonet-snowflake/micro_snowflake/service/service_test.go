package service

import (
	"context"
	"google.golang.org/grpc"
	"micro_snowflake/proto"
	"testing"
	"time"
)

func TestGetUID(t *testing.T) {
	conn, err := grpc.Dial(":1234", grpc.WithInsecure())
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.Close()
	client := proto.NewSfServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	if resp, err := client.GetUID(ctx, &proto.Sf_Nil{}); err != nil {
		cancel()
		t.Error(err)
		return
	} else {
		cancel()
		t.Log("GetUID => ", resp.Uid)
	}
}

func BenchmarkGetUID(b *testing.B) {
	conn, err := grpc.Dial(":1234", grpc.WithInsecure())
	if err != nil {
		b.Error(err)
		return
	}
	defer conn.Close()
	client := proto.NewSfServiceClient(conn)

	for i := 0; i < b.N; i++ {
		if _, err := client.GetUID(context.Background(), &proto.Sf_Nil{}); err != nil {
			b.Fatal(err)
		}
	}
}
