package utils

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/pkg/errors"
	"sync"
)

var config = hystrix.CommandConfig{
	Timeout:                5000,
	MaxConcurrentRequests:  8,
	SleepWindow:            1000,
	ErrorPercentThreshold:  30,
	RequestVolumeThreshold: 5,
}

type runFunc func() error

type Hystrix struct {
	sync.Map
	callback string
}

func NewHystrix(back string) *Hystrix {
	return &Hystrix{callback: back}
}

func (s *Hystrix) Run(name string, run runFunc) error {
	if _, ok := s.Load(name); !ok {
		hystrix.ConfigureCommand(name, config)
		s.Store(name, name)
	}
	return hystrix.Do(name, func() error {
		return run()
	}, func(err error) error {
		return errors.New(s.callback)
	})
}
