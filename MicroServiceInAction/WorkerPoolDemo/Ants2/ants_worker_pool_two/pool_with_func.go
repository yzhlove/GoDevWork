package ants_worker_pool_two

import (
	"errors"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

type runFunc func(interface{})

type PoolWithFunc struct {
	capacity   int32
	running    int32
	queue      GoWorkerFuncList
	state      int32
	mutex      sync.Locker
	cond       *sync.Cond
	pFunc      runFunc
	cache      sync.Pool
	blockCount int
	opts       *Options
}

func NewPoolWithFunc(size int, f runFunc, options ...ModeOption) (*PoolWithFunc, error) {
	if size <= 0 {
		return nil, errors.New("size must overflow 0")
	}
	if f == nil {
		return nil, errors.New("func no is nil")
	}
	opts := InitOptions(options...)
	if opts.Expire <= 0 {
		return nil, errors.New("expire must overflow 0")
	}

	p := &PoolWithFunc{
		capacity: int32(size),
		pFunc:    f,
		mutex:    NewSpinMutex(),
		opts:     opts,
		queue:    make([]*GoWorkerWithFunc, 0, size),
	}
	p.cache.New = func() interface{} {
		return &GoWorkerWithFunc{p: p, args: make(chan interface{}, 1)}
	}
	p.cond = sync.NewCond(p.mutex)
	go p.monitor()
	return p, nil
}

func (p *PoolWithFunc) monitor() {
	heartbeat := time.NewTicker(p.opts.Expire)
	defer heartbeat.Stop()

	var expireWorkers []*GoWorkerWithFunc
	for range heartbeat.C {
		if p.GetState() == Closed {
			break
		}
		t := time.Now()
		p.mutex.Lock()
		sort.Sort(p.queue)
		index := sort.Search(len(p.queue), func(i int) bool {
			return t.Sub(p.queue[i].expire) > p.opts.Expire
		})
		if index != len(p.queue) {
			expireWorkers = append(expireWorkers[:0], p.queue[index:]...)
			for i := index; i < len(p.queue); i++ {
				p.queue[i] = nil
			}
			p.queue = p.queue[:index]
			p.mutex.Unlock()
			for i, w := range expireWorkers {
				ClosedArgs(w.args)
				expireWorkers[i] = nil
			}
			if p.Running() == 0 {
				p.cond.Broadcast()
			}
		}
	}
}

func (p *PoolWithFunc) Invoke(args interface{}) error {
	if p.GetState() == Closed {
		return errors.New("pool is closed")
	}
	if w := p.Pop(); w != nil {
		w.args <- args
		return nil
	}
	return errors.New("pool over load")
}

func (p *PoolWithFunc) Pop() *GoWorkerWithFunc {
	getWorker := func() *GoWorkerWithFunc {
		if w, ok := p.cache.Get().(*GoWorkerWithFunc); ok {
			w.run()
			return w
		}
		panic("go worker func type err")
	}
	p.mutex.Lock()
	defer p.mutex.Unlock()
	//如果队列中有可用worker
	if p.queue.Len() > 0 {
		w := p.queue[p.queue.Len()-1]
		p.queue = p.queue[:p.queue.Len()-1]
		return w
	}
	//如果没有达到队列上限新创建一个worker
	if p.Running() < p.Cap() {
		return getWorker()
	}
	//如果是非等待状态则直接返回
	if p.opts.NonBlocking {
		return nil
	}
	//等待可用worker
	for {
		if max := p.opts.MaxBlockTask; max != 0 && p.blockCount >= max {
			return nil
		}
		p.blockCount++
		p.cond.Wait()
		p.blockCount--
		if p.Running() == 0 {
			return getWorker()
		}
		if p.queue.Len() == 0 {
			continue
		}
		index := p.queue.Len() - 1
		w := p.queue[index]
		p.queue[index] = nil
		p.queue = p.queue[:index]
		return w
	}
}

func (p *PoolWithFunc) Put(w *GoWorkerWithFunc) bool {
	if p.GetState() == Closed || p.Running() > p.Cap() {
		return false
	}
	w.expire = time.Now()
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.queue = append(p.queue, w)
	p.cond.Signal()
	return true
}

func (p *PoolWithFunc) GetState() int32 {
	return atomic.LoadInt32(&p.state)
}

func (p *PoolWithFunc) Inc() {
	atomic.AddInt32(&p.running, 1)
}

func (p *PoolWithFunc) Dec() {
	atomic.AddInt32(&p.running, -1)
}

func (p *PoolWithFunc) Running() int {
	return int(atomic.LoadInt32(&p.running))
}

func (p *PoolWithFunc) Cap() int {
	return int(atomic.LoadInt32(&p.capacity))
}

func (p *PoolWithFunc) Release() {
	atomic.StoreInt32(&p.state, Closed)
	p.mutex.Lock()
	defer p.mutex.Unlock()
	for _, w := range p.queue {
		ClosedArgs(w.args)
	}
	p.queue = nil
}
