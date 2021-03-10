package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

//sync.Cond

var done bool

func read(name string, c *sync.Cond) {
	c.L.Lock()
	if !done {
		c.Wait()
	}
	fmt.Println("read name => ", name)
	c.L.Unlock()
}

func write(name string, c *sync.Cond) {
	fmt.Println(name, " start write data ...")
	time.Sleep(time.Second * 3)
	c.L.Lock()
	done = true
	c.L.Unlock()
	c.Broadcast()
	fmt.Println(name, " write name ")
}

func main() {

	cond := sync.NewCond(&sync.Mutex{})

	for i := 0; i < 5; i++ {
		go read(strconv.Itoa(i+1)+"-->", cond)
	}
	write("hello world", cond)

	time.Sleep(time.Second * 5)

}
