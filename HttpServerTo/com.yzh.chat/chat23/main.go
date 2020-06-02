package main

import (
	"errors"
	"log"
	"sync"
	"time"
)

const (
	numberBits      uint8 = 12                      //id序号占的bits数
	workerBits      uint8 = 10                      //工作机器所占的bits数
	numberMax       int64 = -1 ^ (-1 << numberBits) //id序号的最大值(4096)
	workerIdMax     int64 = -1 ^ (-1 << workerBits) //工作机器的ID最大值(1024)
	timeShift       uint8 = workerBits + numberBits //时间戳所需的偏移量
	workerShift     uint8 = numberBits              //工作机器ID所需的偏移量
	sub             int64 = 1525705533000
	defaultWorkerId       = 1 //默认工作机器ID
)

var _work *Worker
var once sync.Once

type Worker struct {
	sync.RWMutex
	lastTimestamp int64 //上一次生成的时间戳
	workerID      int64 //节点ID
	number        int64 //已经生成的ID数
}

func newWorker(workerID int64) (*Worker, error) {
	if workerID < 0 || workerID > workerIdMax {
		return nil, errors.New("out of range")
	}
	return &Worker{workerID: workerID}, nil
}

func (w *Worker) genID() int64 {
	w.Lock()
	defer w.Unlock()
	//当前毫秒级的时间戳
	now := time.Now().UnixNano() / 1e6
	if w.lastTimestamp == now {
		w.number++
		if w.number >= numberMax {
			for now <= w.lastTimestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		w.number = 0
		w.lastTimestamp = now
	}
	id := (now-sub)<<timeShift | w.workerID<<workerShift | w.number
	return id
}

func GetID() int64 {
	once.Do(func() {
		log.Println("==============")
		if w, err := newWorker(defaultWorkerId); err != nil {
			panic(err)
		} else {
			_work = w
		}
	})
	return _work.genID()
}
