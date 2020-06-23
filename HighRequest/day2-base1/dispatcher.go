package main

type Dispatcher struct {
	WorkerPool chan chan Job
	MaxWorkers int
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	return &Dispatcher{MaxWorkers: maxWorkers, WorkerPool: make(chan chan Job, maxWorkers)}
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}
	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			go func(job Job) {
				jobChannel := <-d.WorkerPool
				jobChannel <- job
			}(job)
		}
	}
}
