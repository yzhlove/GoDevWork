package main

import (
	"fmt"
	"time"
)

type content struct {
	msgCh    chan int
	notifyCh chan struct{}
	closeCh  chan struct{}
}

func main() {

	c := &content{
		msgCh:    make(chan int, 15),
		notifyCh: make(chan struct{}),
		closeCh:  make(chan struct{}),
	}

	go func() {
		defer close(c.closeCh)
		for i := 1; i <= 10; i++ {
			select {
			case <-c.notifyCh:
				fmt.Println("break notify ch")
				return
			case c.msgCh <- i:
				fmt.Println("senf msg i=> ", i)
			}
			time.Sleep(time.Second)
		}
		fmt.Println("send message ok ...")
	}()

	time.Sleep(time.Second * 3)
	go func() {
		select {
		case <-c.closeCh:
			fmt.Println("close message ch")
			close(c.msgCh)
		}
	}()
	close(c.notifyCh)

	time.Sleep(time.Second * 5)
}
