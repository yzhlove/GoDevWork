package app

import (
	"WorkSpace/GoDevWork/GiftServer/config"
	"WorkSpace/GoDevWork/GiftServer/db"
	pb "WorkSpace/GoDevWork/GiftServer/proto"
	"WorkSpace/GoDevWork/GiftServer/pubsub"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"math/rand"
	"net"
	"sync"
	"time"
)

type app struct {
	mu       sync.RWMutex
	listener *net.TCPListener
}

func (p *app) Init() error {

	rand.Seed(time.Now().UnixNano())

	//pubsub
	pubsub.Init()

	//redis
	if err := db.Init(); err != nil {
		return err
	}

	return nil
}

func (p *app) Start() error {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", config.Listen)
	if err != nil {
		return err
	}
	if p.listener, err = net.ListenTCP("tcp", tcpAddr); err != nil {
		return err
	}
	log.Info("listening on:", p.listener.Addr().String())
	server := grpc.NewServer(grpc.MaxConcurrentStreams(config.MaxConcurrentStreams))
	pb.RegisterGiftServiceServer(server, p)

	go server.Serve(p.listener)

	return nil
}

func (p *app) Stop() error {
	log.Debug("app stopping ...")
	p.listener.Close()
	log.Debug("app stop")
	return nil
}

func New() *app {
	return &app{}
}
