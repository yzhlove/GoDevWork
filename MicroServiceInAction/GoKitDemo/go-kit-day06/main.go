package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd/etcdv3"
	grpc_transport "github.com/go-kit/kit/transport/grpc"
	"go-kit-six/agent/pb"
	"go-kit-six/agent/src"
	"go-kit-six/utils"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var addr string
var quitChan = make(chan error, 1)

func init() {
	flag.StringVar(&addr, "g", "127.0.0.1:1234", "grpc addr")
	flag.Parse()
}

func main() {
	var (
		etcdAddr = []string{"127.0.0.1:2379"}
		svcName  = "/register/svc.user.agent"
		ttl      = 5 * time.Second
	)
	utils.NewLoggerServer()
	logger := utils.GetLog()

	opts := etcdv3.ClientOptions{
		DialTimeout:   ttl,
		DialKeepAlive: ttl,
	}
	etcdClient, err := etcdv3.NewClient(context.Background(), etcdAddr, opts)
	if err != nil {
		panic(err)
	}
	register := etcdv3.NewRegistrar(etcdClient, etcdv3.Service{
		Key:   fmt.Sprintf("%s/%s", svcName, addr),
		Value: addr,
	}, log.NewNopLogger())

	go func() {
		rateLimiter := rate.NewLimiter(5, 2)
		service := src.NewService(logger)
		endpoints := src.NewLoginEndpoint(service, rateLimiter)
		server := src.NewGrpc(endpoints, logger)
		listener, err := net.Listen("tcp", addr)
		if err != nil {
			logger.Info("[Agent] listen ", zap.Error(err))
			quitChan <- fmt.Errorf("[Agent] listen err:%v", err.Error())
			return
		}
		register.Register()
		s := grpc.NewServer(grpc.UnaryInterceptor(grpc_transport.Interceptor))
		pb.RegisterUserServer(s, server)
		if err := s.Serve(listener); err != nil {
			logger.Info("[Agent] server ", zap.Error(err))
			quitChan <- fmt.Errorf("[Agent] run error:%v", err.Error())
			return
		}
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		<-c
		quitChan <- fmt.Errorf("[Signal] Error")
	}()

	logger.Info("[Agent] run " + addr)
	if err := <-quitChan; err != nil {
		fmt.Printf("[Agent Server ] runinng error: %s\n", err.Error())
	}
	register.Deregister()
	logger.Error("[Agent] run error ... ")
	time.Sleep(time.Second * 2)
}
