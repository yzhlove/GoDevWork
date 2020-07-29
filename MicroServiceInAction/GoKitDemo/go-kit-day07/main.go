package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	metrics_prometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/go-kit/kit/sd/etcdv3"
	grpc_transport "github.com/go-kit/kit/transport/grpc"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go-kit-seven/agent/pb"
	"go-kit-seven/agent/src"
	"go-kit-seven/utils"
	"go.uber.org/zap"
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

var addr string
var prometheus_addr string
var quitChan = make(chan error, 1)

func init() {
	flag.StringVar(&addr, "g", "127.0.0.1:1235", "grpc addr")
	flag.StringVar(&prometheus_addr, "p", ":1234", "prometheus_addr addr")
	flag.Parse()
}

func main() {

	var (
		etcd_addr = []string{"127.0.0.1:2379"}
		svcName   = "svc.user.agent"
		ttl       = 5 * time.Second
	)
	utils.NewLoggerServer()
	logger := utils.GetLog()
	options := etcdv3.ClientOptions{DialTimeout: ttl, DialKeepAlive: ttl}
	client, err := etcdv3.NewClient(context.Background(), etcd_addr, options)
	if err != nil {
		panic(err)
	}
	register := etcdv3.NewRegistrar(client, etcdv3.Service{
		Key:   fmt.Sprintf("%s/%d", svcName, crc32.ChecksumIEEE([]byte(addr))),
		Value: addr,
	}, log.NewNopLogger())

	go func() {
		count := metrics_prometheus.NewCounterFrom(prometheus.CounterOpts{
			Subsystem: "user_agent",
			Name:      "request_count",
			Help:      "请求数量统计",
		}, []string{"method"})
		histogram := metrics_prometheus.NewHistogramFrom(prometheus.HistogramOpts{
			Subsystem: "user_agent",
			Name:      "request_consume",
			Help:      "请求时长统计",
		}, []string{"method"})

		l := rate.NewLimiter(3, 1)
		service := src.NewService(logger, count, histogram)
		endpoints := src.NewEndpointServer(service, l)
		userServer := src.NewGrpcServer(endpoints, logger)

		listener, err := net.Listen("tcp", addr)
		if err != nil {
			logger.Error("[UserAgent] listener ", zap.Error(err))
			quitChan <- err
			return
		}
		register.Register()
		grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpc_transport.Interceptor))
		pb.RegisterUserServer(grpcServer, userServer)
		quitChan <- grpcServer.Serve(listener)
	}()

	go func() {
		logger.Info("monitor metrics prometheus addr:", zap.String("prometheus_addr", prometheus_addr))
		http.Handle("/metrics", promhttp.Handler())
		quitChan <- http.ListenAndServe(prometheus_addr, nil)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		quitChan <- fmt.Errorf("%s", <-c)
	}()

	if err := <-quitChan; err != nil {
		register.Deregister()
		panic(err)
	}
	logger.Info("[UserAgent] running.")
}
