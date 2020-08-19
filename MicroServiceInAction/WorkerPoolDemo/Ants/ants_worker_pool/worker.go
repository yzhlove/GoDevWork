package ants_worker_pool

import (
	"runtime"
	"time"
)

type GoWorker struct {
	pool        *Pool
	task        chan func()
	recycleTime time.Time
}

func (w *GoWorker) run() {
	w.pool.incRunning()
	go func() {
		defer func() {
			w.pool.decRunning()
			w.pool.workerCache.Put(w)
			if p := recover(); p != nil {
				if ph := w.pool.options.PanicHandler; ph != nil {
					ph(p)
				} else {
					w.pool.options.Logger.Printf("workers exit from a panic: %v \n", p)
					var buf [4096]byte
					n := runtime.Stack(buf[:], false)
					w.pool.options.Logger.Printf("worker exit from panic:%v \n", string(buf[:n]))
				}
			}
		}()

		for f := range w.task {
			if f == nil {
				return
			}
			f()
			if ok := w.pool.revertWorker(w); !ok {
				return
			}
		}
	}()
}
