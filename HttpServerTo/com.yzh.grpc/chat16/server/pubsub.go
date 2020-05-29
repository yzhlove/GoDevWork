package main

import (
	"WorkSpace/GoDevWork/HttpServerTo/com.yzh.grpc/chat16/proto"
	"fmt"
	"sync"
	"time"
)

const Max = 16

type pubsub struct {
	sync.RWMutex
	//queue map[uint32]chan ziface{}
	queue map[chan interface{}]uint32
}

func New() *pubsub {
	return &pubsub{queue: make(map[chan interface{}]uint32)}
}

func (p *pubsub) Sub(z uint32) chan interface{} {

	channel := make(chan interface{}, Max)
	p.Lock()
	p.queue[channel] = z
	p.Unlock()
	go func() {
		for _, msg := range MsgList() {
			if msg.Zone == 0 || msg.Zone == z {
				channel <- msg
				time.Sleep(time.Second * 3)
			}
		}
	}()
	return channel
}

func (p *pubsub) Pub(msg *proto.Manager_Msg) {
	for channel, z := range p.queue {
		if msg.Zone == 0 || msg.Zone == z {
			channel <- msg
		}
	}
}

func (p *pubsub) Close(channel chan interface{}) {
	fmt.Println("close channel ...")
	if _, ok := p.queue[channel]; ok {
		p.Lock()
		delete(p.queue, channel)
		p.Unlock()
	}
	if _, isClose := <-channel; isClose {
		close(channel)
	}
}

func MsgList() []*proto.Manager_Msg {
	return []*proto.Manager_Msg{
		{Zone: 0, Var: "send message 1"},
		{Zone: 0, Var: "send message 2"},
		{Zone: 0, Var: "send message 3"},
		{Zone: 0, Var: "send message 4"},
		{Zone: 0, Var: "send message 5"},
		{Zone: 1, Var: "send message 6"},
		{Zone: 1, Var: "send message 7"},
		{Zone: 1, Var: "send message 8"},
		{Zone: 1, Var: "send message 9"},
		{Zone: 2, Var: "send message 10"},
		{Zone: 2, Var: "send message 11"},
		{Zone: 2, Var: "send message 12"},
		{Zone: 2, Var: "send message 13"},
		{Zone: 3, Var: "send message 14"},
		{Zone: 3, Var: "send message 15"},
	}
}
