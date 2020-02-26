package main

import (
	"fmt"
	"time"
)

func main() {

	var a chan int
	a = make(chan int, 10)
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("set i => ", i)
			a <- i
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(time.Second * 3)
	a = nil
	time.Sleep(time.Second)

}
