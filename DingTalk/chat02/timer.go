package main

import (
	"fmt"
	"sync"
	"time"
)

type tasker interface {
	task() error
}

var (
	timer  = time.NewTicker(time.Minute * 3)
	tQueue = &taskQueue{}
)

type taskQueue struct {
	sync.RWMutex
	ch []tasker
}

func (t *taskQueue) set(task tasker) {
	t.Lock()
	defer t.Unlock()
	t.ch = append(t.ch, task)
}

func (t *taskQueue) task() {
	t.RLock()
	defer t.RUnlock()
	for _, c := range t.ch {
		if err := c.task(); err != nil {
			fmt.Println("task err:", err)
		}
	}
}

func (t *taskQueue) run() {
	for {
		select {
		case <-timer.C:
			t.task()
		}
	}
}
