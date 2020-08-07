package client

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-kit/kit/log"
	"go-kit-nine/agent/pb"
	"go-kit-nine/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"os"
	"testing"
	"time"
)

func TestNewAgentClient(t *testing.T) {
	var _logger log.Logger
	{
		_logger = log.NewLogfmtLogger(os.Stderr)
		_logger = log.With(_logger, "ts", log.DefaultTimestampUTC)
		_logger = log.With(_logger, "func", log.DefaultCaller)
	}

	utils.NewLogger()
	client, err := NewAgentClient([]string{"127.0.0.1:2379"}, _logger)
	if err != nil {
		t.Error(err)
		return
	}
	hx := utils.NewHystrix("drop service")
	cbs, _, _ := hystrix.GetCircuit("login")
	for i := 0; i < 100; i++ {
		time.Sleep(time.Millisecond * 10)
		userAgent, err := client.UserAgentClient()
		if err != nil {
			t.Error(err)
			return
		}
		err = hx.Run("login", func() error {
			if ack, err := userAgent.Login(context.Background(), &pb.UserLogic_Login{
				Account:  "yuzihan",
				Password: "12345",
			}); err != nil {
				t.Error(err)
				return err
			} else {
				t.Log("token -> ", ack.Token)
			}
			return nil
		})
		t.Log("hystrix status => ", cbs.IsOpen(), cbs.AllowRequest())
		if err != nil {
			t.Error(err)
			return
		}
	}
}

func TestGrpc(t *testing.T) {

	addr := "127.0.0.1:1234"
	tracer, closer, err := utils.NewJaegerTracer("user_agent_client")
	if err != nil {
		t.Error(err)
		return
	}
	defer closer.Close()
	conn, err := grpc.Dial(addr, grpc.WithUnaryInterceptor(utils.JaegerClientInterceptor(tracer)))
	if err != nil {
		t.Error(err)
		return
	}
	userClient := pb.NewUserClient(conn)
	UID := utils.GetUID()
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(utils.CONTEXT_UID, UID))
	for i := 0; i < 100; i++ {
		if resp, err := userClient.RpcUserLogin(ctx, &pb.UserLogic_Login{
			Account:  "yuizihan",
			Password: "12345",
		}); err != nil {
			t.Error(err)
		} else {
			t.Log("token -> ", resp.Token)
		}
	}
}
