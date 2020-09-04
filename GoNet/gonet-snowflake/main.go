package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"micro_snowflake/config"
	"micro_snowflake/etcdclient"
	"micro_snowflake/proto"
	"micro_snowflake/service"
	"net"
	"time"
)

var cfg *config.Config

func init() {
	host := flag.String("host", ":1234", "snowflake host")
	etcdHost := flag.String("etcd", "localhost:2379", "etcd host")
	prefix := flag.String("prefix", "snowflake", "etcd key prefix")
	root := flag.String("root", "go-net", "etcd library")
	id := flag.String("machine_id", "0", "machine id")
	flag.Parse()
	cfg = config.New(*host, *id, *root, *prefix, []string{*etcdHost}, 5*time.Second)
}

func main() {
	e := etcdclient.Etcd{}
	e.Init(cfg)

	s := &service.Server{}
	s.Init(cfg)

	l, err := net.Listen("tcp", cfg.Host)
	if err != nil {
		panic("listent error:" + err.Error())
	}

	rpc := grpc.NewServer()
	proto.RegisterSfServiceServer(rpc, s)
	log.Info("start service by Host:", cfg.Host)
	if err := rpc.Serve(l); err != nil {
		panic(err)
	}
}
