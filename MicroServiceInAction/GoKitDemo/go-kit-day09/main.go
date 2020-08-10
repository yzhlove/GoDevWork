package main

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	metrics_prometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/go-kit/kit/sd/etcdv3"
	grpc_transport "github.com/go-kit/kit/transport/grpc"
	grpc_middle "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go-kit-nine/agent/pb"
	"go-kit-nine/agent/src"
	"go-kit-nine/utils"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"hash/crc32"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	prometheus_addr = "127.0.0.1:2345"
	grpc_addr       = "127.0.0.1:1234"
	etcd_addr       = []string{"127.0.0.1:2379"}
)

var overChan = make(chan error, 1)

func main() {

	serviceName := "service.user.agent"
	ttl := 5 * time.Second

	utils.NewLogger()
	logger := utils.GetLog()

	opts := etcdv3.ClientOptions{DialKeepAlive: ttl, DialTimeout: ttl}
	client, err := etcdv3.NewClient(context.Background(), etcd_addr, opts)
	if err != nil {
		panic(err)
	}
	register := etcdv3.NewRegistrar(client,
		etcdv3.Service{
			Key:   fmt.Sprintf("%s/%s", serviceName, crc32.ChecksumIEEE([]byte(grpc_addr))),
			Value: grpc_addr,
		}, log.NewNopLogger())

	go func() {
		tracer, _, err := utils.NewJaegerTracer("user_agent_server")
		if err != nil {
			overChan <- err
			return
		}

		c := metrics_prometheus.NewCounterFrom(prometheus.CounterOpts{
			Subsystem: "user_agent",
			Name:      "request_count",
			Help:      "请求次数",
		}, []string{"method"})

		h := metrics_prometheus.NewHistogramFrom(prometheus.HistogramOpts{
			Subsystem: "user_agent",
			Name:      "request_consume",
			Help:      "请求消耗时间",
		}, []string{"method"})

		l := rate.NewLimiter(10, 5)

		service := src.NewService(logger, c, h, tracer)
		endpointer := src.NewLoginEndpoint(service, l, tracer)
		grpcInter := src.NewGrpc(endpointer, logger)
		listen, err := net.Listen("tcp", grpc_addr)
		if err != nil {
			overChan <- err
			return
		}
		register.Register()

		intercept := grpc_middle.ChainUnaryServer(
			grpc_transport.Interceptor,
			grpc_opentracing.UnaryServerInterceptor(grpc_opentracing.WithTracer(tracer)),
			grpc_zap.UnaryServerInterceptor(logger),
		)

		s := grpc.NewServer(grpc.UnaryInterceptor(intercept))
		pb.RegisterUserServer(s, grpcInter)
		overChan <- s.Serve(listen)
	}()

	go func() {
		m := http.NewServeMux()
		m.Handle("/metrics", promhttp.Handler())
		overChan <- http.ListenAndServe(prometheus_addr, m)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		overChan <- fmt.Errorf("%s", <-c)
	}()

	defer register.Deregister()

	if err := <-overChan; err != nil {
		panic(err)
	}
}
