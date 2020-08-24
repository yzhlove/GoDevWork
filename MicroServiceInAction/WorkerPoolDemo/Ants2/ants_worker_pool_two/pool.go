package ants_worker_pool_two

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

type Pool struct {
	capacity   int32
	running    int32
	queue      WorkerQueue
	state      int32
	mutex      sync.Locker
	cond       *sync.Cond
	cache      sync.Pool
	blockCount int
	opts       *Options
}

func NewPool(size int, options ...ModeOption) (*Pool, error) {
	opts := InitOptions(options...)

	if size == 0 {
		size = -1 //不限worker制增长容量
	}

	if expire := opts.Expire; expire <= 0 {
		return nil, errors.New("expire must overflow 0")
	}

	p := &Pool{
		capacity: int32(size),
		mutex:    NewSpinMutex(),
		opts:     opts,
	}

	p.cache.New = func() interface{} {
		return &GoWorker{p: p, taskChan: make(chan func(), 1)}
	}

	if p.opts.IsAlloc {
		p.queue = NewWorkerQueue(LoopQueue, size)
	} else {
		p.queue = NewWorkerQueue(StackQueue, 0)
	}

	p.cond = sync.NewCond(p.mutex)
	go p.monitor()
	return p, nil
}

func (p *Pool) monitor() {
	heartbeat := time.NewTicker(p.opts.Expire)
	defer heartbeat.Stop()

	for range heartbeat.C {
		if p.GetState() == Closed {
			break
		}
		p.mutex.Lock()
		expireWorkers := p.queue.checkExpire(p.opts.Expire)
		p.mutex.Unlock()
		for i := range expireWorkers {
			ClosedChan(expireWorkers[i].taskChan)
		}
		if p.Running() == 0 {
			p.cond.Broadcast()
		}
	}
}

func (p *Pool) Submit(task func()) error {
	if p.GetState() == Closed {
		return errors.New("pool is closed")
	}
	if worker := p.Pop(); worker != nil {
		worker.taskChan <- task
		return nil
	}
	return errors.New("pop worker error")
}

func (p *Pool) Pop() *GoWorker {

	getWorker := func() *GoWorker {
		if w, ok := p.cache.Get().(*GoWorker); ok {
			w.run()
			return w
		}
		panic("pool cache type error")
	}

	p.mutex.Lock()
	defer p.mutex.Unlock()
	w := p.queue.pop()

	if w == nil {
		//可以无限增长
		if p.Cap() == -1 {
			return getWorker()
		}
		//是否达到worker上限
		if p.Running() < p.Cap() {
			return getWorker()
		}
		//不等待，直接返回
		if p.opts.NonBlocking {
			return nil
		}
		//等待，直到获取可执行的worker为止
		for {
			//如果超出最大等待队列，直接退出
			if p.opts.MaxBlockTask != 0 && p.blockCount >= p.opts.MaxBlockTask {
				return nil
			}
			p.blockCount++
			p.cond.Wait() //等待可用worker
			p.blockCount--
			//检查可用worker
			if p.Running() == 0 {
				return getWorker()
			}
			//从队列里面拿到一个worker
			if w := p.queue.pop(); w == nil {
				continue
			} else {
				return w
			}
		}
	}
	return w
}

func (p *Pool) Put(w *GoWorker) bool {
	if (p.Cap() != -1 && p.Running() >= p.Cap()) || p.GetState() == Closed {
		return false
	}
	w.expire = time.Now()
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if err := p.queue.push(w); err != nil {
		return false
	}
	//唤醒block queue
	p.cond.Signal()
	return true
}

func (p *Pool) GetState() int32 {
	return atomic.LoadInt32(&p.state)
}

func (p *Pool) Cap() int {
	return int(atomic.LoadInt32(&p.capacity))
}

func (p *Pool) Running() int {
	return int(atomic.LoadInt32(&p.running))
}

func (p *Pool) Inc() {
	atomic.AddInt32(&p.running, 1)
}

func (p *Pool) Dec() {
	atomic.AddInt32(&p.running, -1)
}

func (p *Pool) Release() {
	atomic.StoreInt32(&p.state, Closed)
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.queue.reset()
}
