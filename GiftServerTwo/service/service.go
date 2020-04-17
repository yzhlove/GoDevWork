package service

import (
	"os"
	"os/signal"
	"syscall"
)

type Interface interface {
	Init() error
	Start() error
	Stop() error
}

func Run(svc Interface, sigs ...os.Signal) (err error) {
	if err = svc.Init(); err != nil {
		return
	}
	if err = svc.Start(); err != nil {
		return
	}
	if len(sigs) == 0 {
		sigs = []os.Signal{syscall.SIGINT, syscall.SIGTERM}
	}
	sChan := make(chan os.Signal, 1)
	signal.Notify(sChan, sigs...)
	<-sChan
	return svc.Stop()
}
