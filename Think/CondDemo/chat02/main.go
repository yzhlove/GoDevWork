package chat02

import (
	"runtime"
	"sync"
	"sync/atomic"
)

type SpinLocker uint32

func (s *SpinLocker) Lock() {
	for !atomic.CompareAndSwapUint32((*uint32)(s), 0, 1) {
		runtime.Gosched()
	}
}

func (s *SpinLocker) Unlock() {
	atomic.StoreUint32((*uint32)(s), 0)
}

func NewSpinLocker() sync.Locker {
	return new(SpinLocker)
}
