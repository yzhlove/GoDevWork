package main

import (
	"fmt"
	"strings"
	"sync"
)

const Max = 16

type pubsubQueue struct {
	sync.RWMutex
	suber map[string]chan interface{}
}

func (p *pubsubQueue) Sub(tag string) (ch chan interface{}, ok bool) {
	fmt.Println("SubMessage => Tag:", tag)
	if ch, ok = p.suber[tag]; !ok {
		ch = make(chan interface{}, Max)
		p.Lock()
		p.suber[tag] = ch
		p.Unlock()
	}
	return
}

func (p *pubsubQueue) Pub(msg string) {
	for tag, ch := range p.suber {
		if tag == "all" || strings.HasPrefix(msg, tag) {
			fmt.Println("PubMessage=> Tag:", tag, " Message:", msg)
			ch <- msg
		}
	}
}

func (p *pubsubQueue) CloseChan(tag string) {
	fmt.Println("Close Tag:", tag)
	if ch, ok := p.suber[tag]; ok {
		if _, isClose := <-ch; isClose {
			close(ch)
		}
		p.Lock()
		delete(p.suber, tag)
		p.Unlock()
	}
}

func New() *pubsubQueue {
	return &pubsubQueue{suber: make(map[string]chan interface{}, Max)}
}
