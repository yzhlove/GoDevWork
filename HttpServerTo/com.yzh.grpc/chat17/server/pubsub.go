package main

import (
	"WorkSpace/GoDevWork/HttpServerTo/com.yzh.grpc/chat17/proto"
	"fmt"
	"sync"
	"time"
)

const Max = 16

type Manager struct {
	Ch   chan interface{}
	Exit chan struct{}
}

type Queue struct {
	sync.RWMutex
	sub map[*Manager]uint32
}

func NewQueue() *Queue {
	return &Queue{sub: make(map[*Manager]uint32, Max)}
}

func (q *Queue) Pub(msg *proto.Manager_Msg) {
	for manager, z := range q.sub {
		if msg.Zone == 0 || msg.Zone == z {
			manager.Ch <- msg
		}
	}
}

func (q *Queue) Sub(z uint32) *Manager {
	manager := &Manager{
		Ch:   make(chan interface{}, Max),
		Exit: make(chan struct{}),
	}
	ipc := make(chan interface{}, Max)
	q.Lock()
	q.sub[manager] = z
	q.Unlock()
	go sendFirstMsg(ipc, manager)
	return manager
}

func (q *Queue) Close(manager *Manager) {
	close(manager.Exit)
	close(manager.Ch)
	if _, ok := q.sub[manager]; ok {
		q.Lock()
		delete(q.sub, manager)
		q.Unlock()
	}
	fmt.Println("PubSub Close ...")
}

func sendFirstMsg(ipc chan interface{}, manager *Manager) {
	defer close(ipc)
	go func() {
		for {
			select {
			case s, ok := <-ipc:
				if ok {
					manager.Ch <- s
				} else {
					return
				}
			case <-manager.Exit:
				fmt.Println("==> send message close ...")
				return
			}
		}
	}()
	for _, s := range getMsgList() {
		ipc <- s
		time.Sleep(5 * time.Second)
	}
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
