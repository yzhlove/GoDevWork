package main

import (
	"log"
	"sync"
	"time"
)

func main() {

	lock := &sync.Mutex{}
	cond := sync.NewCond(lock)
	wg := sync.WaitGroup{}
	var status bool
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			defer cond.L.Unlock() //不要忘记Unlock ，is dangerous
			cond.L.Lock()
			log.Println("before lock:", index)
			for !status {
				cond.Wait()
			}
			log.Println("after lock:", index)
		}(i)
	}

	go func() {
		time.Sleep(time.Second)
		log.Println("active all ...")
		status = true
		cond.Broadcast()
	}()

	wg.Wait()
	log.Println("over ...")
}
