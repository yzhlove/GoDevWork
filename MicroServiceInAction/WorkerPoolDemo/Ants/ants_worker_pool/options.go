package ants_worker_pool

import "time"

type Option func(opts *Options)

func loadOptions(opts ...Option) *Options {
	o := new(Options)
	for _, opt := range opts {
		opt(o)
	}
	return o
}

type Options struct {
	ExpireDuration   time.Duration
	PreAlloc         bool
	MaxBlockingTasks int
	Nonblocking      bool
	PanicHandler     func(interface{})
	Logger           Logger
}

func WithOptions(options Options) Option {
	return func(opts *Options) {
		*opts = options
	}
}

func WithExpireDuration(expire time.Duration) Option {
	return func(opts *Options) {
		opts.ExpireDuration = expire
	}
}

func WithPreAlloc(preAlloc bool) Option {
	return func(opts *Options) {
		opts.PreAlloc = preAlloc
	}
}

func WithMaxBlockingTasks(max int) Option {
	return func(opts *Options) {
		opts.MaxBlockingTasks = max
	}
}

func WithNonBlocking(block bool) Option {
	return func(opts *Options) {
		opts.Nonblocking = block
	}
}

func WithPanicHandler(h func(interface{})) Option {
	return func(opts *Options) {
		opts.PanicHandler = h
	}
}

func WithLogger(logger Logger) Option {
	return func(opts *Options) {
		opts.Logger = logger
	}
}
