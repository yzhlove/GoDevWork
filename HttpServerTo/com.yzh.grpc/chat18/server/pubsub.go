package main

import (
	"WorkSpace/GoDevWork/HttpServerTo/com.yzh.grpc/chat18/proto"
	"fmt"
	"sync"
	"time"
)

const Max = 16

type Content struct {
	MsgCh   chan interface{}
	StopCh  chan struct{}
	CloseCh chan struct{}
}

type Queue struct {
	sync.RWMutex
	SubQueue map[*Content]uint32
}

func NewQueue() *Queue {
	return &Queue{SubQueue: make(map[*Content]uint32, 4)}
}

func (p *Queue) Pub(msg *proto.Manager_Msg) {
	for c, z := range p.SubQueue {
		if msg.Zone == 0 || msg.Zone == z {
			c.MsgCh <- msg
		}
	}
}

func (p *Queue) Sub(z uint32) *Content {
	c := &Content{
		MsgCh:   make(chan interface{}, Max),
		StopCh:  make(chan struct{}),
		CloseCh: make(chan struct{}),
	}
	p.Lock()
	p.SubQueue[c] = z
	p.Unlock()
	go toFirstMsg(c)
	return c
}

func (p *Queue) Close(c *Content) {
	if _, ok := p.SubQueue[c]; ok {
		p.Lock()
		delete(p.SubQueue, c)
		p.Unlock()
	}
	go func() {
		select {
		case <-c.CloseCh:
			fmt.Println("close message chan ... ")
			close(c.MsgCh)
		}
	}()
	fmt.Println("close stop chan running ...")
	close(c.StopCh)
}

func toFirstMsg(c *Content) {
	defer close(c.CloseCh)
	for _, msg := range getMsgList() {
		select {
		case <-c.StopCh:
			fmt.Println("stop chan exit")
			return
		case c.MsgCh <- msg:
		}
		time.Sleep(time.Second * 3)
	}
	fmt.Println("send message exit ...")
}

func getMsgList() []*proto.Manager_Msg {
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
