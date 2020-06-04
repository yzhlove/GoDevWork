package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("cpu number => ", runtime.NumCPU())
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go process(i)
	}
	wg.Wait()
	log.Println("ok.")
}

func process(id int) {
	fmt.Println("id => ", id)
	for {
		time.Sleep(time.Second * 10)
	}
}
