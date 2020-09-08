package signal

import (
	log "github.com/sirupsen/logrus"
	"micro_agent/utils"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	WaitGroup sync.WaitGroup
	Die       = make(chan struct{})
)

func sigHandler() {
	defer utils.Trace()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM)
	for {
		if msg := <-ch; msg == syscall.SIGTERM {
			//关闭Agent
			close(Die)
			log.Info("sigterm received")
			WaitGroup.Wait()
			log.Info("agent shutdown")
			os.Exit(0)
		}
	}
}
