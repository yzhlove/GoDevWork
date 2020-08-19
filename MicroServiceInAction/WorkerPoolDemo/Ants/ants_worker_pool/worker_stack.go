package ants_worker_pool

import "time"

type workerStack struct {
	items  []*GoWorker
	expire []*GoWorker
	size   int
}

func newWorkerStack(size int) *workerStack {
	return &workerStack{
		items: make([]*GoWorker, 0, size),
		size:  size,
	}
}

func (wq *workerStack) len() int {
	return len(wq.items)
}

func (wq *workerStack) isEmpty() bool {
	return len(wq.items) == 0
}

func (wq *workerStack) insert(worker *GoWorker) error {
	wq.items = append(wq.items, worker)
	return nil
}

func (wq *workerStack) detach() *GoWorker {
	l := wq.len()
	if l == 0 {
		return nil
	}
	w := wq.items[l-1]
	wq.items = wq.items[:l-1]
	return w
}

func (wq *workerStack) retrieveExpire(duration time.Duration) []*GoWorker {
	n := wq.len()
	if n == 0 {
		return nil
	}
	expireTime := time.Now().Add(-duration)
	index := wq.binarySearch(0, n-1, expireTime)

	wq.expire = wq.expire[:0]
	if index != -1 {
		wq.expire = append(wq.expire, wq.items[:index]...)
		m := copy(wq.items, wq.items[index+1:])
		wq.items = wq.items[:m]
	}
	return wq.expire
}

func (wq *workerStack) binarySearch(l, r int, expireTime time.Time) int {
	var mid int
	for l <= r {
		mid = (l + r) >> 1
		if expireTime.Before(wq.items[mid].recycleTime) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return r
}

func (wq *workerStack) reset() {
	for i := 0; i < wq.len(); i++ {
		wq.items[i].task <- nil
	}
	wq.items = wq.items[:0]
}
