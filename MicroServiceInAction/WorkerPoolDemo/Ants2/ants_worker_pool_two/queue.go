package ants_worker_pool_two

import (
	"errors"
	"time"
)

var (
	errQueueIsReleased = errors.New("queue is nil")
	errQueueIsFull     = errors.New("queue is full")
)

type WorkerQueue interface {
	len() int
	empty() bool
	push(worker *GoWorker) error
	pop() *GoWorker
	checkExpire(t time.Duration) []*GoWorker
	reset()
}

type QueueType int

const (
	StackQueue QueueType = 1 << iota
	LoopQueue
)

func NewWorkerQueue(tp QueueType, size int) WorkerQueue {
	if tp == LoopQueue {
		return NewLoopQueue(size)
	}
	return NewStackQueue(size)
}
