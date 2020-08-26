package ants_worker_pool_two

import (
	"testing"
	"time"
)

func TestLoopQueue(t *testing.T) {

	nt := time.Now()

	queue := NewLoopQueue(5)

	queue.push(&GoWorker{expire: nt, _id: 1})
	queue.push(&GoWorker{expire: nt.Add(time.Second), _id: 2})
	queue.push(&GoWorker{expire: nt.Add(time.Second * 2), _id: 3})
	queue.push(&GoWorker{expire: nt.Add(time.Second * 3), _id: 4})
	queue.push(&GoWorker{expire: nt.Add(time.Second * 4), _id: 5})

	go func() {
		tick := time.NewTicker(time.Second)
		defer tick.Stop()
		time.Sleep(2 * time.Second)
		for range tick.C {
			if expireWorker := queue.checkExpire(0); expireWorker != nil {
				for _, w := range expireWorker {
					t.Log("time ", w.expire, " _id ", w._id)
				}
			} else {
				t.Log("expire nil")
			}
		}
	}()

	time.Sleep(time.Second * 10)

}
