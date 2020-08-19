package ants_worker_pool

import (
	"runtime"
	"time"
)

type GoWorkerWithFunc struct {
	pool        *PoolWithFunc
	args        chan interface{}
	recycleTime time.Time
}

func (w *GoWorkerWithFunc) run() {
	w.pool.incRunning()
	go func() {
		defer func() {
			w.pool.decRunning()
			w.pool.workerCache.Put(w)
			if p := recover(); p != nil {
				if ph := w.pool.options.PanicHandler; ph != nil {
					ph(p)
				} else {
					w.pool.options.Logger.Printf("worker with func exists from a panic: %v \n", p)
					var buf [4096]byte
					n := runtime.Stack(buf[:], false)
					w.pool.options.Logger.Printf("worker with func exists from panic:%s \n", string(buf[:n]))
				}
			}
		}()

		for args := range w.args {
			if args == nil {
				return
			}
			w.pool.poolFunc(args)
			if ok := w.pool.revertWorker(w); !ok {
				return
			}
		}
	}()
}
