package ants_worker_pool

import "time"

type loopQueue struct {
	items  []*GoWorker
	expire []*GoWorker
	head   int
	tail   int
	size   int
	isFull bool
}

func newWorkerLoopQueue(size int) *loopQueue {
	return &loopQueue{
		items: make([]*GoWorker, size),
		size:  size,
	}
}

func (wq *loopQueue) len() int {
	if wq.size == 0 {
		return 0
	}
	if wq.head == wq.tail {
		if wq.isFull {
			return wq.size
		}
		return 0
	}
	if wq.tail > wq.head {
		return wq.tail - wq.head
	}
	return wq.size - wq.head + wq.tail
}

func (wq *loopQueue) isEmpty() bool {
	return wq.head == wq.tail && !wq.isFull
}

func (wq *loopQueue) insert(worker *GoWorker) error {
	if wq.size == 0 {
		return errQueueIsReleased
	}

	if wq.isFull {
		return errQueueIsFull
	}

	wq.items[wq.tail] = worker
	wq.tail++

	if wq.tail == wq.size {
		wq.tail = 0
	}

	if wq.tail == wq.head {
		wq.isFull = true
	}

	return nil
}

func (wq *loopQueue) detach() *GoWorker {
	if wq.isEmpty() {
		return nil
	}

	w := wq.items[wq.head]
	wq.head++
	if wq.head == wq.size {
		wq.head = 0
	}
	wq.isFull = false
	return w
}

func (wq *loopQueue) retrieveExpire(duration time.Duration) []*GoWorker {
	if wq.isEmpty() {
		return nil
	}

	wq.expire = wq.expire[:0]
	expireTime := time.Now().Add(-duration)

	for !wq.isEmpty() {
		if expireTime.Before(wq.items[wq.head].recycleTime) {
			break
		}
		wq.expire = append(wq.expire, wq.items[wq.head])
		wq.head++
		if wq.head == wq.size {
			wq.head = 0
		}
		wq.isFull = false
	}
	return wq.expire
}

func (wq *loopQueue) reset() {
	if wq.isEmpty() {
		return
	}
RELEASEING:
	if w := wq.detach(); w != nil {
		w.task <- nil
		goto RELEASEING
	}
	wq.items = wq.items[:0]
	wq.size = 0
	wq.head = 0
	wq.tail = 0
}
