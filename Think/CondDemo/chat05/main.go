package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// once

type once struct {
	done  uint32 //hot path
	mutex sync.Mutex
}

func (o *once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 0 {
		o.fo(f)
	}
}

func (o *once) fo(f func()) {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}

func main() {

	oc := &once{}

	oc.Do(func() {
		fmt.Println("hello")
	})

	oc.Do(func() {
		fmt.Println("world")
	})

}
