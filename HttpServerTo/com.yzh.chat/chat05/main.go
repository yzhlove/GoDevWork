package main

import (
	"fmt"
	"sync"
)

func main() {

	resultChan := make(chan chan string, 1<<4)
	chanList := make([]chan string, 0, 3)
	wg := sync.WaitGroup{}
	go func() {
		for {
			select {
			case ch, ok := <-resultChan:
				if ok {
					msg, tag := <-ch
					fmt.Printf("get chan message: status(%v) message(%v) \n", tag, msg)
					wg.Done()
				}
			}
		}
	}()
	for i := 0; i < 3; i++ {
		result := make(chan string)
		chanList = append(chanList, result)
		resultChan <- result
		wg.Add(1)
	}
	go func() {
		chanList[1] <- "1 like qingqing "
	}()
	go func() {
		chanList[0] <- "0 xjj is befault girl "
	}()
	go func() {
		chanList[2] <- "2 is all good "
	}()
	//stop
	wg.Wait()
	close(resultChan)
}
