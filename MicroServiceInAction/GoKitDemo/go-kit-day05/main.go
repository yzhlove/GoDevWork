package main

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	ept "go-kit-five/endpoint"
	"go-kit-five/pb"
	"go-kit-five/service"
	"go-kit-five/transport"
	"go-kit-five/utils"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"net"
)

func main() {

	utils.NewLoggerServer()
	logger := utils.GetLog()
	limit := rate.NewLimiter(3, 2)
	s := service.NewService(logger)
	endpoint := ept.NewEndpoint(s, logger, limit)
	grpcHandle := transport.NewGrpcServer(endpoint, logger, limit)
	logger.Info("server run :1234")
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer(grpc.UnaryInterceptor(grpctransport.Interceptor))
	pb.RegisterUserServer(server, grpcHandle)
	if err := server.Serve(listener); err != nil {
		logger.Error("server", zap.Error(err))
		panic(err)
	}
}
