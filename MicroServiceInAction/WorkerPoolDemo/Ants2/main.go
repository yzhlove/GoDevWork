package main

import (
	"ants_worker_pool_two"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	TestPoolWithFunc()

}

func TestPool() {
	pool, err := ants_worker_pool_two.NewPool(5,
		ants_worker_pool_two.WithExpire(1),
		//ants_worker_pool_two.WithAlloc(true),
	)
	if err != nil {
		panic(err)
	}

	defer pool.Release()
	var wg sync.WaitGroup
	run := func() {
		time.Sleep(time.Second)
		fmt.Println("Hello World")
		wg.Done()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		if err := pool.Submit(run); err != nil {
			panic(err)
		}
	}

	wg.Wait()
	log.Println("running count:", pool.Running())
	log.Println("finish all task.")
}

func TestPoolWithFunc() {

	type args struct {
		i   int32
		_id int
	}

	var count int32
	var wg sync.WaitGroup
	runFunc := func(i interface{}) {
		if a, ok := i.(*args); ok {
			atomic.AddInt32(&count, a.i)
			log.Println("args: ", a.i, " id:", a._id)
			wg.Done()
			return
		}
		wg.Done()
		panic("args type err")
	}

	p, err := ants_worker_pool_two.NewPoolWithFunc(5, runFunc,
		ants_worker_pool_two.WithExpire(1))
	if err != nil {
		panic(err)
	}

	defer p.Release()

	for i := 0; i < 100; i++ {
		wg.Add(1)
		if err := p.Invoke(&args{_id: i, i: 1}); err != nil {
			panic(err)
		}
	}
	wg.Wait()
	fmt.Println("running go routines: ", p.Running())
	fmt.Println("count value :", atomic.LoadInt32(&count))
}
