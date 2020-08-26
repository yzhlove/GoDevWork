package ants_worker_pool_two

import (
	"fmt"
	"time"
)

type GoWorkerWithFunc struct {
	p      *PoolWithFunc
	args   chan interface{}
	expire time.Time
}

func (w *GoWorkerWithFunc) run() {
	w.p.Inc()
	go func() {
		defer func() {
			w.p.Dec()
			w.p.cache.Put(w)
			if x := recover(); x != nil {
				trace(fmt.Sprint(x))
			}
		}()

		for args := range w.args {
			w.p.pFunc(args)
			if ok := w.p.Put(w); !ok {
				panic("put worker to queue error")
			}
		}
	}()
}

func ClosedArgs(c chan interface{}) {
	c <- nil
}

type GoWorkerFuncList []*GoWorkerWithFunc

func (s GoWorkerFuncList) Len() int {
	return len(s)
}

func (s GoWorkerFuncList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s GoWorkerFuncList) Less(i, j int) bool {
	return s[i].expire.Before(s[j].expire)
}
