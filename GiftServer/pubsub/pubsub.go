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
	CloseCh  chan struct{}
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
		CloseCh:  make(chan struct{}),
	}
	_pubsub.Lock()
	_pubsub.queue[c] = z
	_pubsub.Unlock()
	go postCodeMessage(c)
	return c
}

func Close(c *Content) {
	if _, ok := _pubsub.queue[c]; ok {
		_pubsub.Lock()
		delete(_pubsub.queue, c)
		_pubsub.Unlock()
	}
	go func() {
		select {
		case <-c.CloseCh:
			close(c.MsgChan)
		}
	}()
	close(c.StopChan)
}

func postCodeMessage(c *Content) {
	defer close(c.CloseCh)
	for _, code := range entity.GetCodesMap() {
		select {
		case <-c.StopChan:
			return
		case c.MsgChan <- manager.GeneratePtoCodeMessage(code):
		}
	}
}
