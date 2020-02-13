package service

import (
	"os"
	"os/signal"
	"syscall"
)

var signalNotify = signal.Notify

type ServiceInterface interface {
	Init() error
	Start() error
	Stop() error
}

func Run(service ServiceInterface, sig ...os.Signal) (err error) {
	if err = service.Init(); err != nil {
		return
	}
	if err = service.Start(); err != nil {
		return
	}
	if len(sig) == 0 {
		sig = []os.Signal{syscall.SIGINT, syscall.SIGTERM}
	}
	signalChan := make(chan os.Signal, 1)
	signalNotify(signalChan, sig...)
	<-signalChan

	return service.Stop()
}
