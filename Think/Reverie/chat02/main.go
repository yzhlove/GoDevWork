package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {

	task := NewTaskQueue(100)
	task.Run()
	for i := 0; i < 1000; i++ {
		if err := task.AsyncTask(&Logic{Content: fmt.Sprintf("logic id: %d ", i)}); err != nil {
			log.Println("set task err:", i, " - ", err)
		}
	}
	time.Sleep(time.Second * 10)
}

type LogicInterface interface {
	Do() error
}

type TaskInterface interface {
	AsyncTask(logic LogicInterface) error
}

type TaskQueue struct {
	mutex     sync.Mutex
	taskQueue chan LogicInterface
}

func NewTaskQueue(size int) *TaskQueue {
	return &TaskQueue{
		taskQueue: make(chan LogicInterface, size),
	}
}

func (task *TaskQueue) Run() {
	go func() {
		for logic := range task.taskQueue {
			if err := logic.Do(); err != nil {
				log.Println("logic err:", err)
			}
		}
	}()
}

func (task *TaskQueue) AsyncTask(logic LogicInterface) error {
	select {
	case task.taskQueue <- logic:
		return nil
	default:
		return errors.New("queue overflow")
	}
}

type Logic struct {
	Content string
}

func (l *Logic) Do() error {
	log.Println("logic content -> ", l.Content)
	return nil
}
