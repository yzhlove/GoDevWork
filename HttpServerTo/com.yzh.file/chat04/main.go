package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	fmt.Println("number", uint64(1)<<63)

	fmt.Printf("%b \n", uint64(1)<<63)
	fmt.Printf("%9b \n", uint64(1<<64-1))

	fmt.Printf("%b \n", uint64(1)<<11-1)
	test()
	time.Sleep(time.Second * 5)
}

func test() {

	var mutext sync.Mutex
	var queue []int

	for i := 0; i < 10; i++ {
		go func() {
			defer mutext.Unlock()
			mutext.Lock()
			for jk := 0; jk < 10; jk++ {
				queue = append(queue, jk)
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(queue)

}
