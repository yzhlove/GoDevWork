package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)
	tag := make(chan struct{})

	go func() {
		for i := 1; i <= 10; i++ {
			select {
			case <-tag:
				fmt.Println("tag is close...")
				return
			default:
				select {
				case <-tag:
					fmt.Println("tag is close...")
					return
				case ch <- i:
					fmt.Println("write data => ", i)
				}
			}
		}
	}()

	//go func() {
	//	for number := range ch {
	//		fmt.Println("number => ", number)
	//	}
	//	fmt.Println("ch is close ...")
	//}()

	fmt.Println("sleep 3 times...")
	time.Sleep(time.Second * 3)
	close(tag)
	close(ch)
	fmt.Println("waiting ...")
	time.Sleep(time.Second * 5)

}
