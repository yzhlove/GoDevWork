package ants_worker_pool

import (
	"errors"
	"time"
)

var (
	errQueueIsFull     = errors.New("the queue is full")
	errQueueIsReleased = errors.New("the queue length is zero")
)

type workerArray interface {
	len() int
	isEmpty() bool
	insert(worker *GoWorker) error
	detach() *GoWorker
	retrieveExpire(duration time.Duration) []*GoWorker
	reset()
}

type arrayType int

const (
	stackType arrayType = 1 << iota
	loopQueueType
)

func newWorkerArray(typ arrayType, size int) workerArray {
	if typ == loopQueueType {
		return newWorkerLoopQueue(size)
	}
	return newWorkerStack(size)
}
