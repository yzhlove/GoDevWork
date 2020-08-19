package ants_worker_pool

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewWorkerStack(t *testing.T) {
	size := 100
	stack := newWorkerStack(size)

	assert.EqualValues(t, 0, stack.len(), "Len error")
	assert.Equal(t, true, stack.isEmpty(), "IsEmpty error")
	assert.Nil(t, stack.detach(), "Dequeue error")
}

func TestWorkerStack(t *testing.T) {
	queue := newWorkerArray(arrayType(-1), 0)
	for i := 0; i < 5; i++ {
		if err := queue.insert(&GoWorker{recycleTime: time.Now()}); err != nil {
			t.Error(err)
			break
		}
	}

	assert.EqualValues(t, 5, queue.len(), "Len error")
	expire := time.Now()

	if err := queue.insert(&GoWorker{recycleTime: expire}); err != nil {
		t.Error(err)
		return
	}

	time.Sleep(time.Second)

	for i := 0; i < 6; i++ {
		if err := queue.insert(&GoWorker{recycleTime: time.Now()}); err != nil {
			t.Error(err)
			return
		}
	}

	assert.EqualValues(t, 12, queue.len(), "Len error")
	queue.retrieveExpire(time.Second)
	assert.EqualValues(t, 6, queue.len(), "Len error")
}

func TestSearch(t *testing.T) {
	queue := newWorkerStack(0)

	expire_1 := time.Now()
	queue.insert(&GoWorker{recycleTime: time.Now()})

	assert.EqualValues(t, 0, queue.binarySearch(0, queue.len()-1, time.Now()), "index should be 0")
	assert.EqualValues(t, -1, queue.binarySearch(0, queue.len()-1, expire_1), "index should be -1")

	expire_2 := time.Now()
	queue.insert(&GoWorker{recycleTime: time.Now()})

	assert.EqualValues(t, -1, queue.binarySearch(0, queue.len()-1, expire_1), "index should be -1")
	assert.EqualValues(t, 0, queue.binarySearch(0, queue.len()-1, expire_2), "index should be 0")
	assert.EqualValues(t, 1, queue.binarySearch(0, queue.len()-1, time.Now()), "index should be 1 ")

	for i := 0; i < 5; i++ {
		queue.insert(&GoWorker{recycleTime: time.Now()})
	}

	expire_3 := time.Now()
	queue.insert(&GoWorker{recycleTime: expire_3})

	for i := 0; i < 10; i++ {
		queue.insert(&GoWorker{recycleTime: time.Now()})
	}

	assert.EqualValues(t, 7, queue.binarySearch(0, queue.len()-1, expire_3), "index should be 7")
}
