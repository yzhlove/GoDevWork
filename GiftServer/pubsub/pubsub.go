package pubsub

import (
	"WorkSpace/GoDevWork/GiftServer/entity"
	"WorkSpace/GoDevWork/GiftServer/manager"
	"sync"
)

/*
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
*/

const QueueMax = 16

type PubSub struct {
	sync.Mutex
	queue map[uint32]chan interface{}
}

var _pubsub PubSub

func Init() error {
	_pubsub.queue = make(map[uint32]chan interface{}, 4)
	return nil
}

func Pub(zs []uint32, msg interface{}) {

	var tag map[uint32]struct{}
	if l := len(zs); l > 0 {
		tag = make(map[uint32]struct{}, l)
		for _, t := range zs {
			tag[t] = struct{}{}
		}
	}
	//for z, channel := range _pubsub.queue {
	//
	//}

}

func Sub(z uint32) (channel chan interface{}) {
	ok := false
	if channel, ok = _pubsub.queue[z]; !ok {
		channel = make(chan interface{}, QueueMax)
		_pubsub.Lock()
		_pubsub.queue[z] = channel
		_pubsub.Unlock()
	}
	go func() {
		_send := func(z uint32, zones []uint32) bool {
			if len(zones) == 0 {
				return true
			}
			for _, t := range zones {
				if z == t {
					return true
				}
			}
			return false
		}
		for _, code := range entity.GetCodesMap() {
			if _send(z, code.ZoneIds) {
				channel <- manager.GeneratePtoCodeMessage(code)
			}
		}
	}()
	return
}
