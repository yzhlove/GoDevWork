package ants_worker_pool

import "time"

type workerArray interface {
	len() int
	isEmpty() bool
	insert(worker *goWorker) error
	detach() *goWorker
	retrieveExpire(duration time.Duration) []*goWorker
	reset()
}

type arrayType int

const (
	stackType arrayType = 1 << iota
	loopQueueType
)

func newWorkerArray(typ arrayType, size int) workerArray {

	return nil
}
