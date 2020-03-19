package microkernel

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
)

type State int

const (
	Running State = iota
	Waiting
)

var WrongStateError = errors.New("can not take the operator in the current state")

type CollectorsError struct {
	CollectorErrors []error
}

func (ce CollectorsError) Error() string {
	var strs []string
	for _, err := range ce.CollectorErrors {
		strs = append(strs, err.Error())
	}
	return strings.Join(strs, ";")
}

type Event struct {
	Source  string
	Content string
}

type EventReceiver interface {
	OnEvent(e Event)
}

type Collector interface {
	Init(ev EventReceiver) error
	Start(ctx context.Context) error
	Stop() error
	Destroy() error
}

type Agent struct {
	collectors map[string]Collector
	evtBuf     chan Event
	cancel     context.CancelFunc
	ctx        context.Context
	state      State
}

func (agt *Agent) EventProcess() {
	for {
		select {
		case evt, ok := <-agt.evtBuf:
			if ok {
				fmt.Println("[Agent Event Process] ", evt)
			}
		case <-agt.ctx.Done():
			return
		}
	}
}

func NewAgent(sizeEvtBuf int) *Agent {
	return &Agent{
		collectors: make(map[string]Collector, 4),
		evtBuf:     make(chan Event, sizeEvtBuf),
		state:      Waiting,
	}
}

func (agt *Agent) RegisterCollector(name string, collector Collector) error {
	if agt.state != Waiting {
		return WrongStateError
	}
	agt.collectors[name] = collector
	return nil
}

func (agt *Agent) startCollectors() error {
	var errs CollectorsError
	var err error
	var mutex sync.Mutex
	for name, collector := range agt.collectors {
		go func(name string, collector Collector, ctx context.Context) {
			defer mutex.Unlock()
			err = collector.Start(agt.ctx)
			mutex.Lock()
			if err != nil {
				errs.CollectorErrors = append(errs.CollectorErrors,
					errors.New(name+":"+err.Error()))
			}
		}(name, collector, agt.ctx)
	}
	return errs
}

func (agt *Agent) stopCollectors() error {
	var err error
	var errs CollectorsError
	for name, collector := range agt.collectors {
		if err = collector.Stop(); err != nil {
			errs.CollectorErrors = append(errs.CollectorErrors,
				errors.New(name+":"+err.Error()))
		}
	}
	return errs
}

func (agt *Agent) destroyCollectors() error {
	var err error
	var errs CollectorsError
	for name, collector := range agt.collectors {
		if err = collector.Destroy(); err != nil {
			errs.CollectorErrors = append(errs.CollectorErrors,
				errors.New(name+":"+err.Error()))
		}
	}
	return errs
}

func (agt *Agent) Start() error {
	if agt.state != Waiting {
		return WrongStateError
	}
	agt.state = Running
	agt.ctx, agt.cancel = context.WithCancel(context.Background())
	go agt.EventProcess()
	return agt.startCollectors()
}

func (agt *Agent) Stop() error {
	if agt.state != Running {
		return WrongStateError
	}
	agt.state = Waiting
	agt.cancel()
	return agt.stopCollectors()
}

func (agt *Agent) Destroy() error {
	if agt.state != Waiting {
		return WrongStateError
	}
	return agt.destroyCollectors()
}

func (agt *Agent) OnEvent(evt Event) {
	agt.evtBuf <- evt
}
