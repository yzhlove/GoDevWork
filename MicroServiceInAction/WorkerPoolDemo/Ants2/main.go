package main

import (
	"ants_worker_pool_two"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {

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
