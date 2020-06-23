package main

import (
	"fmt"
	"sync"
)

type Job struct {
	Pay Payload
}

var JobQueue chan Job

type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan struct{}
	once       sync.Once
}

func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job, MAxQueue),
		quit:       make(chan struct{}, 1),
	}
}

func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobChannel
			select {
			case job := <-w.JobChannel:
				if err := job.Pay.UploadS3(); err != nil {
					fmt.Println("err upload to s3:", err)
				}
			case <-w.quit:
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	w.once.Do(func() {
		go func() { close(w.quit) }()
	})
}
