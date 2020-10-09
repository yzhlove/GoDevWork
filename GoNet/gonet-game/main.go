package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"micro_game/config"
	"micro_game/proto"
	"micro_game/service"
	"net"
	"strings"
	"time"
)

var cfg *config.Config

func init() {
	game := flag.String("id", "game1", "id to service")
	listen := flag.String("listen", ":12138", "listen address and port")
	etcdHost := flag.String("etcd-hosts", "127.0.0.1:2379", "etcd host split(,)")
	etcdRoot := flag.String("etcd-root", "/backends", "etcd root")
	services := flag.String("services", "snowflake-1000", "discovery service")
	file := flag.String("numbers", "number", "table path")
	flag.Parse()
	cfg = &config.Config{
		GameId:      *game,
		Listen:      *listen,
		EtcdHosts:   strings.Split(*etcdHost, ","),
		EtcdRoot:    *etcdRoot,
		Services:    strings.Split(*services, ","),
		PathNumbers: *file,
		Streams:     1024,
		Timeout:     5 * time.Second,
	}
}

func main() {
	l, err := net.Listen("tcp", cfg.Listen)
	if err != nil {
		log.Fatal(err)
	}
	//注册服务
	server := grpc.NewServer(grpc.MaxConcurrentStreams(cfg.Streams))
	_service := new(service.GameService)
	proto.RegisterGameServiceServer(server, _service)
	//初始化其他服务
	service.Init(cfg)
	//numbers.Init(cfg.PathNumbers)
	//启动服务
	log.Info("service running to:", cfg.Listen)
	if err := server.Serve(l); err != nil {
		log.Fatal(err)
	}
}
