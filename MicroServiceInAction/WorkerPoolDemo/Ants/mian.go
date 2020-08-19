package main

import (
	"ants_worker_pool"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var sum int32

func myFunc(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	fmt.Printf("run with %d\n", n)
}

func demoFunc() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello World")
}

func main() {
	defer ants_worker_pool.Release()

	runTimes := 1000

	var wg sync.WaitGroup
	syncCalcuateSum := func() {
		demoFunc()
		wg.Done()
	}

	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		ants_worker_pool.Submit(syncCalcuateSum)
	}
	wg.Wait()
	fmt.Printf("running go routines:%d\n", ants_worker_pool.Running())
	fmt.Printf("finish all tasks.\n")

	p, _ := ants_worker_pool.NewPoolWithFunc(10, func(i interface{}) {
		myFunc(i)
		wg.Done()
	})

	defer p.Release()
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		if err := p.Invoke(int32(i)); err != nil {
			panic(err)
		}
	}
	wg.Wait()
	fmt.Printf("running go routines:%d \n", p.Running())
	fmt.Printf("finish all tasks ,result is %d \n", sum)
	if sum != 499500 {
		panic("error sum")
	}
}
