package main

import (
	"fmt"
	"time"
)

func main() {

	a := make(chan int)
	t := make(chan struct{})
	go func() {
		for i := 0; i < 10; i++ {
			select {
			case a <- i:
			case <-t:
				return
			}
			time.Sleep(time.Second)
		}
		fmt.Println(" exit ....")
	}()
	go func() {
		for n := range a {
			fmt.Println("n =>", n)
		}
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("close chan ...")
	a = nil
	close(t)
	time.Sleep(5 * time.Second)

}
