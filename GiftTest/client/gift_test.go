package client

import (
	"WorkSpace/GoDevWork/GiftTest/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func Test_Generate(t *testing.T) {

	client := GetClient()
	if client == nil {
		t.Error("client is nil")
		return
	}

	/*
		string FixCode = 1;
		uint32 Num = 2;
		int64 StartTime = 3;
		int64 EndTime = 4;
		uint32 TimesPerCode = 5;
		uint32 TimesPerUser = 6;
		repeated uint32 ZoneIds = 7;
		repeated Item Items = 8;
	*/

	req := &pb.Manager_GenReq{
		//FixCode: "yzhhhhhhaaabbc",
		Num:   1,
		Items: []*pb.Manager_Item{{Id: 1, Num: 1}, {Id: 2, Num: 2}},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	if _, err := client.Generate(ctx, req); err != nil {
		t.Error(err)
		return
	}
	cancel()
	t.Log("ok")
}

func Test_List(t *testing.T) {

	client := GetClient()
	if client == nil {
		t.Error("client is nil")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	resp, err := client.List(ctx, &pb.Manager_Nil{})
	if err != nil {
		t.Error(err.Error())
		return
	}
	cancel()

	for _, result := range resp.Details {
		t.Log(result)
	}

}

func Test_Export(t *testing.T) {
	client := GetClient()
	if client == nil {
		t.Error("client is nil")
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	resp, err := client.Export(ctx, &pb.Manager_ExportReq{Id: 1})
	if err != nil {
		t.Error(err)
		return
	}
	cancel()
	fmt.Println("export max ==> ", len(resp.Details))
	for _, result := range resp.Details {
		t.Log(result)
	}
}

func Test_Sync(t *testing.T) {

	client := GetClient()
	if client == nil {
		t.Error("client is nil")
		return
	}
	stream, err := client.Sync(context.Background(), &pb.SyncReq{Zone: 1})
	if err != nil {
		t.Error(err)
		return
	}
	status := make(chan struct{})
	go func() {
		for {
			if result, err := stream.Recv(); err != nil {
				fmt.Printf("stream recv Err:%s \n", err.Error())
				close(status)
				return
			} else {
				fmt.Println("stream result --> ", result)
			}
		}
	}()
	<-status
}

func Test_CodeVerify(t *testing.T) {

	client := GetClient()
	if client == nil {
		t.Error("client is nil")
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := client.CodeVerify(ctx, &pb.VerifyReq{Code: "vsYPoD5o4O1", Zone: 1, UserId: 123456789})
	if err != nil {
		t.Error(err)
		return
	}
	cancel()

	t.Log(fmt.Sprintf("CodeVerify => %d ", resp.Status))

}

func GetClient() pb.GiftServiceClient {

	conn, err := grpc.Dial(":51000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return pb.NewGiftServiceClient(conn)
}
