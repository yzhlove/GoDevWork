package ants_worker_pool_two

import "time"

type Loop struct {
	entries          []*GoWorker
	expireEntries    []*GoWorker
	head, tail, size int
	isFull           bool
}

func NewLoopQueue(size int) *Loop {
	return &Loop{
		size:    size,
		entries: make([]*GoWorker, size),
	}
}

func (l *Loop) len() int {
	if l.size == 0 {
		return 0
	}
	if l.head == l.tail {
		if l.isFull {
			return l.size
		}
		return 0
	}
	if l.tail > l.head {
		return l.tail - l.head
	}
	return l.size - l.head + l.tail
}

func (l *Loop) empty() bool {
	return l.head == l.tail && !l.isFull
}

func (l *Loop) push(w *GoWorker) error {
	if l.size == 0 {
		return errQueueIsReleased
	}
	if l.isFull {
		return errQueueIsFull
	}
	l.entries[l.tail] = w
	l.tail++
	if l.tail == l.size {
		l.tail = 0
	}
	if l.tail == l.head {
		l.isFull = true
	}
	return nil
}

func (l *Loop) pop() *GoWorker {
	if l.empty() {
		return nil
	}
	w := l.entries[l.head]
	l.head++
	if l.head == l.size {
		l.head = 0
	}
	l.isFull = false
	return w
}

func (l *Loop) checkExpire(t time.Duration) []*GoWorker {
	if l.empty() {
		return nil
	}
	l.expireEntries = l.expireEntries[:0]
	timeout := time.Now().Add(-t)
	for !l.empty() {
		//如果worker最后一次运行的时间在给定的时间之前，则清理掉worker
		if timeout.Before(l.entries[l.head].expire) {
			break
		}
		l.expireEntries = append(l.expireEntries, l.entries[l.head])
		l.head++
		if l.head == l.size {
			l.head = 0
		}
		l.isFull = false
	}
	return l.expireEntries
}

func (l *Loop) reset() {
	if l.empty() {
		return
	}
	for {
		if w := l.pop(); w != nil {
			ClosedChan(w.taskChan)
			continue
		}
		l.entries = l.entries[:0]
		l.size, l.head, l.tail = 0, 0, 0
	}
}
