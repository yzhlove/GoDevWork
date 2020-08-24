package ants_worker_pool_two

import (
	"runtime"
	"sync"
	"sync/atomic"
)

type spinMutex uint32

func (s *spinMutex) Lock() {
	for !atomic.CompareAndSwapUint32(getType(s), 0, 1) {
		runtime.Gosched()
	}
}

func (s *spinMutex) Unlock() {
	atomic.StoreUint32(getType(s), 0)
}

func getType(s *spinMutex) *uint32 {
	return (*uint32)(s)
}

func NewSpinMutex() sync.Locker {
	return new(spinMutex)
}
