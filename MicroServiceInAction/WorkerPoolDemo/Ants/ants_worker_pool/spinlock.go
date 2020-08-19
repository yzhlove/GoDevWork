package ants_worker_pool

import (
	"runtime"
	"sync"
	"sync/atomic"
)

////////////////////////////////////////////
// 自旋锁
////////////////////////////////////////////

type spinLock uint32

func (s *spinLock) Lock() {
	for !atomic.CompareAndSwapUint32((*uint32)(s), 0, 1) {
		runtime.Gosched()
	}
}

func (s *spinLock) Unlock() {
	atomic.StoreUint32((*uint32)(s), 0)
}

func NewSpinLock() sync.Locker {
	return new(spinLock)
}
