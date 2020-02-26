package pubsub

import (
	"WorkSpace/GoDevWork/GiftServer/entity"
	"WorkSpace/GoDevWork/GiftServer/manager"
	"sync"
)

const Max = 16

type Content struct {
	MsgChan  chan interface{}
	StopChan chan struct{}
}

type PubSub struct {
	sync.RWMutex
	queue map[*Content]uint32
}

var _pubsub PubSub

func Init() error {
	_pubsub.queue = make(map[*Content]uint32, Max)
	return nil
}

func Pub(zones []uint32, msg interface{}) {
	var t map[uint32]struct{}
	if len(zones) > 0 {
		t = make(map[uint32]struct{}, len(zones))
		for _, z := range zones {
			t[z] = struct{}{}
		}
	}

	for c, z := range _pubsub.queue {
		if t == nil {
			c.MsgChan <- msg
		} else if _, ok := t[z]; ok {
			c.MsgChan <- msg
		}
	}
}

func Sub(z uint32) *Content {
	c := &Content{
		MsgChan:  make(chan interface{}, Max),
		StopChan: make(chan struct{}),
	}
	post := make(chan interface{}, Max)
	_pubsub.Lock()
	_pubsub.queue[c] = z
	_pubsub.Unlock()
	go postCodeMessage(post, c)
	return c
}

func Close(c *Content) {
	close(c.StopChan)
	close(c.MsgChan)
	if _, ok := _pubsub.queue[c]; ok {
		_pubsub.Lock()
		delete(_pubsub.queue, c)
		_pubsub.Unlock()
	}
}

func postCodeMessage(post chan interface{}, c *Content) {
	defer close(post)
	go func() {
		for {
			select {
			case inter, ok := <-post:
				if ok {
					c.MsgChan <- inter
				} else {
					return
				}
			case <-c.StopChan:
				return
			}
		}
	}()
	for _, code := range entity.GetCodesMap() {
		post <- manager.GeneratePtoCodeMessage(code)
	}
}
