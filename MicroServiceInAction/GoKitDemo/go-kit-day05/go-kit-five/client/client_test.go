package client

import (
	"context"
	"go-kit-five/pb"
	"go-kit-five/service"
	"go-kit-five/tool"
	"go-kit-five/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"testing"
)

func TestNewGrpcClient(t *testing.T) {
	logger := tool.NewLogger(
		tool.SetLogPrefix("go-kit"),
		tool.SetIsDevelopment(true),
		tool.SetLevel(zap.DebugLevel))

	conn, err := grpc.Dial(":1234", grpc.WithInsecure())
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.Close()
	server := NewGrpcClient(conn, logger)
	ack, err := server.Login(context.Background(), &pb.Login{Account: "yzh", Passwd: "123456789"})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("ack.Token -> ", ack.Token)
}

func TestNewGrpc(t *testing.T) {

	conn, err := grpc.Dial(":1234", grpc.WithInsecure())
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.Close()
	client := pb.NewUserClient(conn)
	UID := utils.GetUID()
	md := metadata.Pairs(service.ContextReq, UID)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	ack, err := client.RpcUserLogin(ctx, &pb.Login{Account: "yzh", Passwd: "123456789"})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("ack.Token => ", ack.Token)
}
