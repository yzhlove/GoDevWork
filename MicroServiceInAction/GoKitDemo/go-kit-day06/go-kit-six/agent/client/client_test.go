package client

import (
	"context"
	"github.com/go-kit/kit/log"
	"go-kit-six/agent/pb"
	"go-kit-six/agent/src"
	"go-kit-six/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"os"
	"testing"
	"time"
)

func TestNewAgentClient(t *testing.T) {

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	client, err := NewAgentClient([]string{"127.0.0.1:2379"}, logger)
	if err != nil {
		t.Error(err)
		return
	}
	for i := 0; i < 6; i++ {
		time.Sleep(time.Second)
		if agent, err := client.AgentClient(); err != nil {
			t.Error(err)
			return
		} else {
			if ack, err := agent.Login(context.Background(), &pb.Login{Account: "yzh", Passwd: "123456789"}); err != nil {
				t.Error(err)
				return
			} else {
				t.Log("token => ", ack.Token)
			}
		}
	}
}

func TestGrpc(t *testing.T) {
	addr := "localhost:1234"
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.Close()
	client := pb.NewUserClient(conn)
	UID := utils.GetUID()
	md := metadata.Pairs(src.CONTEXT_REQ_UID, UID)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	result, err := client.RpcUserLogin(ctx, &pb.Login{Account: "yzh", Passwd: "123456789"})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(" token => ", result.Token)
}
