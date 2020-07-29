package client

import (
	"context"
	"github.com/go-kit/kit/log"
	"go-kit-seven/agent/pb"
	"go-kit-seven/agent/src"
	"go-kit-seven/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"os"
	"testing"
)

func TestNewAgentClient(t *testing.T) {
	logger := log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	client, err := NewAgentClient([]string{"127.0.0.1:2379"}, logger)
	if err != nil {
		t.Error(err)
		return
	}
	for i := 0; i < 20; i++ {
		userAgent, err := client.UserAgentClient()
		if err != nil {
			t.Error(err)
			return
		}
		ack, err := userAgent.Login(context.Background(), &pb.Login{Account: "yuzihan", Passwd: "123456789"})
		if err != nil {
			t.Error(err)
			continue
		}
		t.Log("token =>", ack.Token)
	}

}

func TestGrpc(t *testing.T) {

	addr := "127.0.0.1:1235"
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	userClient := pb.NewUserClient(conn)
	UID := utils.GetUID()
	md := metadata.Pairs(src.CONTEXT_UID, UID)
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	for i := 0; i < 20; i++ {
		result, err := userClient.RpcUserLogin(ctx, &pb.Login{Account: "yuzihan", Passwd: "123456789"})
		if err != nil {
			t.Error(err)
			return
		}
		t.Log("token =>", result.Token)
	}
}
