package main

import (
	"fmt"
	"sync"
	"time"
)

var sharedRsc = false

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	m := sync.Mutex{}
	c := sync.NewCond(&m)

	go func() {
		defer wg.Done()
		c.L.Lock()
		for sharedRsc == false {
			fmt.Println("go routine wait")
			c.Wait()
		}
		fmt.Println("go routine ", sharedRsc)
		c.L.Unlock()
	}()

	go func() {
		defer wg.Done()
		c.L.Lock()
		for sharedRsc == false {
			fmt.Println("go routine 2 wait")
			c.Wait()
		}
		fmt.Println("go routine 2 ", sharedRsc)
		c.L.Unlock()
	}()

	time.Sleep(2 * time.Second)
	c.L.Lock()
	fmt.Println("main go routine ready")
	sharedRsc = true
	c.Broadcast()
	fmt.Println("main go routine broadcast")
	c.L.Unlock()
	wg.Wait()
}
