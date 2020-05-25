package znet

import (
	"fmt"
	"strconv"
	"zinx-day5-base1/ziface"
)

type MessageHandle struct {
	APIs map[uint32]ziface.RouterInterface
}

func NewMessageHandle() *MessageHandle {
	return &MessageHandle{
		APIs: make(map[uint32]ziface.RouterInterface),
	}
}

func (m *MessageHandle) DoMessageHandle(req ziface.RequestInterface) {
	if handler, ok := m.APIs[req.GetMessageID()]; ok {
		handler.BeforeHandle(req)
		handler.Handle(req)
		handler.AfterHandle(req)
	} else {
		fmt.Println("api msgID = ", req.GetMessageID(), " is not FOUND !")
	}
}

func (m *MessageHandle) RegisterRouter(msgID uint32, router ziface.RouterInterface) {
	if _, ok := m.APIs[msgID]; !ok {
		m.APIs[msgID] = router
		fmt.Println("register router by id => ", msgID)
	} else {
		panic("repeated api msg id => " + strconv.Itoa(int(msgID)))
	}
}
