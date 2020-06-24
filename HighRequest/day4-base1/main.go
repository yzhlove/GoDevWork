package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var MaxQueue = 128
var MaxWorker = 10
var MaxBuffer int64 = 4096
var JobQueue chan Job

func init() {
	JobQueue = make(chan Job, MaxQueue)
}

type Job struct {
	Pay Payload
}

type Worker struct {
	WorkerPool chan chan Job
	JobChan    chan Job
	die        chan struct{}
}

type Dispatcher struct {
	WorkerPool chan chan Job
	MaxWorkers int
}

func NewWorker(workerPool chan chan Job) *Worker {
	return &Worker{
		WorkerPool: workerPool,
		JobChan:    make(chan Job, MaxQueue),
		die:        make(chan struct{}, 1),
	}
}

func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobChan
			select {
			case job, ok := <-w.JobChan:
				if ok {
					job.Pay.Update()
				}
			case <-w.die:
				log.Println("worker is stop.")
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	close(w.die)
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	dispatcher := &Dispatcher{
		MaxWorkers: maxWorkers,
		WorkerPool: make(chan chan Job, MaxWorker),
	}
	return dispatcher
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.MaxWorkers; i++ {
		w := NewWorker(d.WorkerPool)
		w.Start()
	}
	go d.putJob()
}

func (d *Dispatcher) putJob() {
	for {
		select {
		case job, ok := <-JobQueue:
			if ok {
				go func(job Job) {
					jobChan := <-d.WorkerPool
					jobChan <- job
				}(job)
			}
		}
	}
}

type Payload struct {
	Path string `json:"path"`
}

type PayloadCollection struct {
	Version  string    `json:"version"`
	Token    string    `json:"token"`
	Payloads []Payload `json:"data"`
}

func (p *Payload) String() string {
	return fmt.Sprintf("[Payload] path = %s", p.Path)
}

func (p *Payload) Update() {
	time.Sleep(100 * time.Millisecond)
	log.Println(p)
}

func payloadH(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var content = &PayloadCollection{}
	if err := json.NewDecoder(io.LimitReader(r.Body, MaxBuffer)).Decode(&content); err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for _, pay := range content.Payloads {
		JobQueue <- Job{Pay: pay}
	}
	w.WriteHeader(http.StatusOK)
}

func main() {
	d := NewDispatcher(MaxWorker)
	d.Run()
	http.HandleFunc("/payload", payloadH)
	log.Println("start server listen by port -> 0.0.0.0:1234")
	if err := http.ListenAndServe(":1234", nil); err != nil {
		panic("[server] start server err:" + err.Error())
	}
}
