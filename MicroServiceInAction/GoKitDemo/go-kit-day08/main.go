package main

import (
	"context"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-kit/kit/log"
	metrics_prometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/go-kit/kit/sd/etcdv3"
	grpc_transport "github.com/go-kit/kit/transport/grpc"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go-kit-eight/agent/pb"
	"go-kit-eight/agent/src"
	"go-kit-eight/utils"
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

const grpc_addr = "127.0.0.1:1234"
const prometheus_addr = "127.0.0.1:2345"

var quitChan = make(chan error, 1)

func main() {

	etcd_addr := []string{"127.0.0.1:2379"}
	server_name := "server.user.agent"
	ttl := 5 * time.Second

	utils.NewLoggerServer()
	logger := utils.GetLog()

	opts := etcdv3.ClientOptions{DialTimeout: ttl, DialKeepAlive: ttl}
	client, err := etcdv3.NewClient(context.Background(), etcd_addr, opts)
	if err != nil {
		panic(err)
	}

	register := etcdv3.NewRegistrar(client, etcdv3.Service{
		Key:   fmt.Sprintf("%s/%d", server_name, crc32.ChecksumIEEE([]byte(grpc_addr))),
		Value: grpc_addr,
	}, log.NewNopLogger())

	go func() {
		c := metrics_prometheus.NewCounterFrom(prometheus.CounterOpts{
			Subsystem: "user_agent",
			Name:      "request_count",
			Help:      "请求次数统计",
		}, []string{"method"})

		h := metrics_prometheus.NewHistogramFrom(prometheus.HistogramOpts{
			Subsystem: "user_agent",
			Name:      "request_consume",
			Help:      "请求时间统计",
		}, []string{"method"})

		l := rate.NewLimiter(3, 1)
		service := src.NewService(logger, c, h)
		endpoints := src.NewEndpointServer(service, l)
		userServer := src.NewGrpc(endpoints, logger)

		listen, err := net.Listen("tcp", grpc_addr)
		if err != nil {
			quitChan <- err
			return
		}

		register.Register()
		server := grpc.NewServer(grpc.UnaryInterceptor(grpc_transport.Interceptor))
		pb.RegisterUserServer(server, userServer)

		quitChan <- server.Serve(listen)
	}()

	go func() {
		m := http.NewServeMux()
		m.Handle("/metrics", promhttp.Handler())
		quitChan <- http.ListenAndServe(prometheus_addr, m)
	}()

	go func() {
		hs := hystrix.NewStreamHandler()
		hs.Start()
		quitChan <- http.ListenAndServe(net.JoinHostPort("", "9543"), hs)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		quitChan <- fmt.Errorf("%s", <-c)
	}()

	defer register.Deregister()
	if err := <-quitChan; err != nil {
		panic(err)
	}
}
