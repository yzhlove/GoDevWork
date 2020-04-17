package app

import (
	"WorkSpace/GoDevWork/GiftServerTwo/config"
	"WorkSpace/GoDevWork/GiftServerTwo/db"
	"WorkSpace/GoDevWork/GiftServerTwo/pb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"math/rand"
	"net"
	"time"
)

type app struct {
	listener *net.TCPListener
	h        handler
	m        MutexType
	ms       *MutexTypeList
}

func (p *app) Init() (err error) {
	rand.Seed(time.Now().Unix())

	p.ms = NewMutexTypeList()

	if err = db.Init(); err != nil {
		return
	}

	if err = p.h.Init(); err != nil {
		return
	}

	return
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
	sev := grpc.NewServer(grpc.MaxConcurrentStreams(config.MaxConcurrentStreams))
	pb.RegisterGiftServiceServer(sev, p)

	go sev.Serve(p.listener)
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
