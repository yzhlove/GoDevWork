package ants_worker_pool_two

import "time"

type ModeOption func(opts *Options)
type HandlePanic func(x interface{})

type Options struct {
	Expire       time.Duration
	IsAlloc      bool
	MaxBlockTask int
	NonBlocking  bool
}

func InitOptions(opts ...ModeOption) *Options {
	opt := new(Options)
	for _, f := range opts {
		f(opt)
	}
	return opt
}

func WithExpire(t int) ModeOption {
	return func(opts *Options) {
		opts.Expire = time.Duration(t) * time.Second
	}
}

func WithAlloc(v bool) ModeOption {
	return func(opts *Options) {
		opts.IsAlloc = v
	}
}
