package main

import (
	"fmt"
	"time"
)

func main() {

	numberCh := make(chan int, 10)
	tagCh := make(chan struct{})

	go func() {
		for i := 0; i < 100000; i++ {
			numberCh <- i
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			select {
			case number, ok := <-numberCh:
				if ok {
					fmt.Println("number ==> ", number)
				} else {
					fmt.Println("close number")
					return
				}
			case <-tagCh:
				fmt.Println("sleep 3 times ...")
				time.Sleep(time.Second * 3)
				return
			}
		}
	}()
	time.Sleep(time.Second * 5)
	fmt.Println("close chan running ...")
	close(tagCh)
	//close(numberCh)
	time.Sleep(5 * time.Second)
	fmt.Println("exit ...")
}
