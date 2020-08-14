package ants_worker_pool

import (
	"errors"
	"log"
	"math"
	"os"
	"runtime"
	"time"
)

const (
	DefaultAntsPoolSize      = math.MaxInt32
	DefaultCleanIntervalTime = time.Second
)

const (
	OPENED = iota
	CLOSE
)

var (
	ErrInvalidPoolSize     = errors.New("invalid size for pool")
	ErrLackPoolFunc        = errors.New("must provide function for pool")
	ErrInvalidPoolExpire   = errors.New("invalid expiry for pool")
	ErrPoolClosed          = errors.New("this pool has been closed")
	ErrPoolOverLoad        = errors.New("too many go routines blocked on submit or Nonblocking is set")
	ErrInvalidPreAllocSize = errors.New("can not set up a negative capacity under PreAlloc mode")

	workerChanCap = func() int {
		if runtime.GOMAXPROCS(0) == 1 {
			return 0
		}
		return 1
	}
	defaultLogger = Logger(log.New(os.Stderr, "ants", log.LstdFlags))
)

type Logger interface {
	Printf(format string, args ...interface{})
}
