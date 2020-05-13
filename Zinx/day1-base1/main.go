package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	task := NewTask(func() error {
		time.Sleep(3 * time.Second)
		fmt.Println("time -> ", time.Now().String())
		return nil
	})
	pool := NewPool(3)
	go func() {
		for {
			pool.EntryCh <- task
			time.Sleep(time.Second)
		}
	}()
	pool.Run()
}

type TaskFunc func() error

type Task struct {
	fn TaskFunc
}

func NewTask(fn TaskFunc) *Task {
	return &Task{fn: fn}
}

func (t *Task) Execute() error {
	return t.fn()
}

type Pool struct {
	Max     int
	EntryCh chan *Task
	JobsCh  chan *Task
}

func NewPool(max int) *Pool {
	return &Pool{
		Max:     max,
		EntryCh: make(chan *Task),
		JobsCh:  make(chan *Task),
	}
}

func (p *Pool) worker(workerID int) {
	for task := range p.JobsCh {
		if err := task.Execute(); err != nil {
			log.Printf("[Task]  Err: %v \n", err)
		}
		log.Printf("[Task] worker id:%d is ok.\n", workerID)
	}
}

func (p *Pool) Run() {
	for i := 0; i < p.Max; i++ {
		go p.worker(i)
	}
	for task := range p.EntryCh {
		p.JobsCh <- task
	}
	close(p.JobsCh)
	close(p.EntryCh)
}
