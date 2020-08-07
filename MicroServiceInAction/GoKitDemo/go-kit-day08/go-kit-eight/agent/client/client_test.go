package client

import (
	"context"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-kit/kit/log"
	"go-kit-eight/agent/pb"
	"go-kit-eight/agent/src"
	"go-kit-eight/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"os"
	"testing"
	"time"
)

func TestNewUserAgentClient(t *testing.T) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	client, err := NewUserAgentClient([]string{"127.0.0.1:2379"}, logger)
	if err != nil {
		t.Error(err)
		return
	}
	hy := utils.NewHystrix("Service Degradation")
	cbs, _, _ := hystrix.GetCircuit("login")
	for i := 0; i < 100; i++ {
		time.Sleep(time.Millisecond)
		userAgent, err := client.UserAgentClient()
		if err != nil {
			t.Error(err)
			return
		}
		err = hy.Run("login", func() error {
			if _, err := userAgent.Login(context.Background(), &pb.Login{
				Account: "yuzihan", Passwd: "12345",
			}); err != nil {
				return err
			}
			return nil
		})
		fmt.Println("Hystrix Status:", cbs.IsOpen(), " Request:", cbs.AllowRequest())
		if err != nil {
			t.Log(err)
		}
	}
}

func TestGrpc(t *testing.T) {
	addr := "127.0.0.1:1234"
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.Close()
	client := pb.NewUserClient(conn)
	UID := utils.GetUID()
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(src.CONTEXT_UID, UID))
	hy := utils.NewHystrix("grpc service")
	cbs, _, _ := hystrix.GetCircuit("login")
	for i := 0; i < 100; i++ {
		if err := hy.Run("login", func() error {
			if resp, err := client.RpcUserLogin(ctx, &pb.Login{
				Account: "yuzihan", Passwd: "12345",
			}); err != nil {
				return err
			} else {
				t.Log("token => ", resp.Token)
			}
			return nil
		}); err != nil {
			t.Log(err)
		}
		fmt.Println("Hystrix Status:", cbs.IsOpen(), " Request:", cbs.AllowRequest())
	}
}
