package ants_worker_pool

import "time"

type goWorker struct {
	pool        *Pool
	task        chan func()
	recycleTime time.Time
}

func (w *goWorker) run() {

}