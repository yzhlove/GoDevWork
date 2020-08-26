package ants_worker_pool

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewLoopQueue(t *testing.T) {
	size := 100
	queue := newWorkerLoopQueue(size)
	assert.EqualValues(t, 0, queue.len(), "Len error")
	assert.Equal(t, true, queue.isEmpty(), "IsEmpty error")
	assert.Nil(t, queue.detach(), "Dequeue error")
}

func TestLoopQueue(t *testing.T) {
	size := 10
	queue := newWorkerLoopQueue(size)

	for i := 0; i < 5; i++ {
		if err := queue.insert(&GoWorker{recycleTime: time.Now()}); err != nil {
			t.Error(err)
			break
		}
	}
	assert.EqualValues(t, 5, queue.len(), "Len error")
	v := queue.detach()
	t.Log(v)
	assert.EqualValues(t, 4, queue.len(), "Len error")

	time.Sleep(time.Second)

	for i := 0; i < 6; i++ {
		if err := queue.insert(&GoWorker{recycleTime: time.Now()}); err != nil {
			t.Error(err)
			break
		}
	}
	assert.EqualValues(t, 10, queue.len(), "Len error")

	if err := queue.insert(&GoWorker{recycleTime: time.Now()}); err != nil {
		assert.Error(t, err, "Enqueue, error")
	}

	queue.retrieveExpire(time.Second)
	assert.EqualValuesf(t, 6, queue.len(), "Len error:%d", queue.len())
}

func TestLoopQueueExpire(t *testing.T) {
	nt := time.Now()

	queue := newWorkerLoopQueue(5)

	queue.insert(&GoWorker{recycleTime: nt, _id: 1})
	queue.insert(&GoWorker{recycleTime: nt.Add(time.Second), _id: 2})
	queue.insert(&GoWorker{recycleTime: nt.Add(time.Second * 2), _id: 3})
	queue.insert(&GoWorker{recycleTime: nt.Add(time.Second * 3), _id: 4})
	queue.insert(&GoWorker{recycleTime: nt.Add(time.Second * 4), _id: 5})

	go func() {
		tick := time.NewTicker(time.Second)
		defer tick.Stop()
		time.Sleep(2 * time.Second)
		for range tick.C {
			if expireWorker := queue.retrieveExpire(0); expireWorker != nil {
				for _, w := range expireWorker {
					t.Log("time ", w.recycleTime, " _id ", w._id)
				}
			} else {
				t.Log("expire nil")
			}
		}
	}()

	time.Sleep(time.Second * 10)

}
