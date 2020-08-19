package ants_worker_pool

import (
	"sync"
	"sync/atomic"
	"time"
)

type PoolWithFunc struct {
	capacity    int32
	running     int32
	workers     []*GoWorkerWithFunc
	state       int32
	lock        sync.Locker
	cond        *sync.Cond
	poolFunc    func(interface{})
	workerCache sync.Pool
	blockingNum int
	options     *Options
}

func NewPoolWithFunc(size int, f func(interface{}), options ...Option) (*PoolWithFunc, error) {
	if size <= 0 {
		return nil, ErrInvalidPoolSize
	}
	if f == nil {
		return nil, ErrLackPoolFunc
	}
	opts := loadOptions(options...)

	if expire := opts.ExpireDuration; expire < 0 {
		return nil, ErrInvalidPoolExpire
	} else if expire == 0 {
		opts.ExpireDuration = DefaultCleanIntervalTime
	}

	if opts.Logger == nil {
		opts.Logger = defaultLogger
	}

	p := &PoolWithFunc{
		capacity: int32(size),
		poolFunc: f,
		lock:     NewSpinLock(),
		options:  opts,
	}

	p.workerCache.New = func() interface{} {
		return &GoWorkerWithFunc{
			pool: p,
			args: make(chan interface{}, workerChanCap()),
		}
	}

	if p.options.PreAlloc {
		p.workers = make([]*GoWorkerWithFunc, 0, size)
	}

	p.cond = sync.NewCond(p.lock)

	go p.purgePeriodically()

	return p, nil
}

func (p *PoolWithFunc) purgePeriodically() {
	heartbeat := time.NewTicker(p.options.ExpireDuration)
	defer heartbeat.Stop()

	var expireWorkers []*GoWorkerWithFunc
	for range heartbeat.C {
		if atomic.LoadInt32(&p.state) == CLOSED {
			break
		}
		currentTime := time.Now()
		p.lock.Lock()
		idleWorkers := p.workers
		n := len(idleWorkers)
		var i int
		for i = 0; i < n && currentTime.Sub(idleWorkers[i].recycleTime) > p.options.ExpireDuration; i++ {
		}

		expireWorkers = append(expireWorkers[:0], idleWorkers[i:]...)
		if i > 0 {
			m := copy(idleWorkers, idleWorkers[i:])
			for i := m; i < n; i++ {
				idleWorkers[i] = nil
			}
			p.workers = idleWorkers[:m]
		}
		p.lock.Unlock()
		for i, w := range expireWorkers {
			w.args <- nil
			expireWorkers[i] = nil
		}
		if p.Running() == 0 {
			p.cond.Broadcast()
		}
	}
}

func (p *PoolWithFunc) Invoke(args interface{}) error {
	if atomic.LoadInt32(&p.state) == CLOSED {
		return ErrPoolClosed
	}
	var w *GoWorkerWithFunc
	if w = p.retrieveWorker(); w == nil {
		return ErrPoolOverLoad
	}
	w.args <- args
	return nil
}

func (p *PoolWithFunc) Running() int {
	return int(atomic.LoadInt32(&p.running))
}

func (p *PoolWithFunc) Free() int {
	return p.Cap() - p.Running()
}

func (p *PoolWithFunc) Cap() int {
	return int(atomic.LoadInt32(&p.capacity))
}

func (p *PoolWithFunc) Tune(size int) {
	if size <= 0 || size == p.Cap() || p.options.PreAlloc {
		return
	}
	atomic.StoreInt32(&p.capacity, int32(size))
}

func (p *PoolWithFunc) Release() {
	atomic.StoreInt32(&p.state, CLOSED)
	p.lock.Lock()
	for _, w := range p.workers {
		w.args <- nil
	}
	p.workers = nil
	p.lock.Unlock()
}

func (p *PoolWithFunc) Reboot() {
	if atomic.CompareAndSwapInt32(&p.state, CLOSED, OPENED) {
		go p.purgePeriodically()
	}
}

func (p *PoolWithFunc) incRunning() {
	atomic.AddInt32(&p.running, 1)
}

func (p *PoolWithFunc) decRunning() {
	atomic.AddInt32(&p.running, -1)
}

func (p *PoolWithFunc) retrieveWorker() (w *GoWorkerWithFunc) {
	spawnWorker := func() {
		var ok bool
		if w, ok = p.workerCache.Get().(*GoWorkerWithFunc); ok {
			w.run()
		}
	}

	p.lock.Lock()
	idleWorkers := p.workers
	n := len(idleWorkers) - 1
	if n >= 0 {
		w = idleWorkers[n]
		idleWorkers[n] = nil
		p.workers = idleWorkers[:n]
		p.lock.Unlock()
	} else if p.Running() < p.Cap() {
		p.lock.Unlock()
		spawnWorker()
	} else {
		if p.options.Nonblocking {
			p.lock.Unlock()
			return
		}
	REENTRY:
		if p.options.MaxBlockingTasks != 0 && p.blockingNum >= p.options.MaxBlockingTasks {
			p.lock.Unlock()
			return
		}
		p.blockingNum++
		p.cond.Wait()
		p.blockingNum--
		if p.Running() == 0 {
			p.lock.Unlock()
			spawnWorker()
			return
		}
		l := len(p.workers) - 1
		if l < 0 {
			goto REENTRY
		}
		w = p.workers[l]
		p.workers[l] = nil
		p.workers = p.workers[:l]
		p.lock.Unlock()
	}
	return
}

func (p *PoolWithFunc) revertWorker(worker *GoWorkerWithFunc) bool {
	if atomic.LoadInt32(&p.state) == CLOSED || p.Running() > p.Cap() {
		return false
	}
	worker.recycleTime = time.Now()
	p.lock.Lock()
	p.workers = append(p.workers, worker)
	p.cond.Signal()
	p.lock.Unlock()
	return true
}
