package pubsub

import (
	"sync"
)

const Max = 16

type PubSubQueue struct {
	sync.RWMutex
	zoneQueue map[uint32]chan interface{}
}

var _pubsub PubSubQueue

func Init() (err error) {
	_pubsub.zoneQueue = make(map[uint32]chan interface{}, 16)
	return
}

func Pub(zones []uint32, msg interface{}) {
	var status map[uint32]struct{}
	if size := len(zones); size > 0 {
		status = make(map[uint32]struct{}, size)
		for _, zone := range zones {
			status[zone] = struct{}{}
		}
	}
	for zone, ch := range _pubsub.zoneQueue {
		if status != nil {
			if _, ok := status[zone]; !ok {
				continue
			}
		}
		ch <- msg
	}
}

func Sub(zone uint32) (ch chan interface{}, ok bool) {
	if ch, ok = _pubsub.zoneQueue[zone]; !ok {
		ch = make(chan interface{}, Max)
		_pubsub.Lock()
		_pubsub.zoneQueue[zone] = ch
		_pubsub.Unlock()
	}
	return
}

func CloseChan(zone uint32) {
	if ch, ok := _pubsub.zoneQueue[zone]; ok {
		if _, isClose := <-ch; isClose {
			close(ch)
		}
		_pubsub.Lock()
		delete(_pubsub.zoneQueue, zone)
		_pubsub.Unlock()
	}
}
