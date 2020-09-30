package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"micro_geoip/config"
	"micro_geoip/proto"
	"micro_geoip/service"
	"net"
)

var _path = "/Users/yostar/workSpace/gowork/src/GoDevWork/GeoIP2-City.mmdb"

var cfg *config.Config

func init() {
	p := flag.String("path", _path, "city db path")
	host := flag.String("host", ":4388", "service monitor address")
	flag.Parse()
	cfg = &config.Config{Path: *p, Host: *host, Streams: 1024}
}

func main() {
	l, err := net.Listen("tcp", cfg.Host)
	if err != nil {
		log.Fatal(err)
	}
	log.Info(config.SERVICE, " listening to ", cfg.Host)
	server := grpc.NewServer(grpc.MaxConcurrentStreams(cfg.Streams))
	_service := new(service.GeoService)
	_service.Init(cfg)
	proto.RegisterGeoIpServer(server, _service)
	if err := server.Serve(l); err != nil {
		log.Fatal(err)
	}
}
