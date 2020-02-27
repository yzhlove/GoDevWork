package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int, 10)
	tag := make(chan struct{})

	go func() {
		for i := 1; i <= 10; i++ {
			select {
			case ch <- i:
			case <-tag:
				fmt.Println("tag close .")
				return
			}
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for number := range ch {
			fmt.Println("number => ", number)
		}
		fmt.Println("ch is close .")
	}()

	time.Sleep(3 * time.Second)

	go func() {
		select {
		case <-tag:
			fmt.Println("close ch ...")
			close(ch)
		}
	}()
	fmt.Println("close tag running ...")
	close(tag)
	time.Sleep(5 * time.Second)
}
