package znet

import (
	"fmt"
	"zinx-day8-base1/config"
	"zinx-day8-base1/ziface"
)

type MsgHandle struct {
	APIs           map[uint32]ziface.RouterInterface
	WorkerPoolSize uint32
	TaskQueue      []chan ziface.RequestInterface
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		APIs:           make(map[uint32]ziface.RouterInterface),
		WorkerPoolSize: config.GlobalConfig.WorkerPoolSize,
		TaskQueue:      make([]chan ziface.RequestInterface, config.GlobalConfig.WorkerPoolSize),
	}
}

func (m *MsgHandle) StartOnWorker(workerID int, taskQueue chan ziface.RequestInterface) {
	fmt.Printf("[worker] id:%d is started.\n", workerID)
	for {
		select {
		case req, ok := <-taskQueue:
			if ok {
				m.Do(req)
				break
			}
			return
		}
	}
}

func (m *MsgHandle) StartWorkerPool() {
	for i := 0; i < int(m.WorkerPoolSize); i++ {
		m.TaskQueue[i] = make(chan ziface.RequestInterface, config.GlobalConfig.MaxWorkerTaskSize)
		go m.StartOnWorker(i+1, m.TaskQueue[i])
	}
}

func (m *MsgHandle) Do(request ziface.RequestInterface) {
	if h, ok := m.APIs[request.GetMsgID()]; ok {
		h.BeforeHandle(request)
		h.Handle(request)
		h.AfterHandle(request)
	} else {
		fmt.Println("api msg id:", request.GetMsgID(), " not found")
	}
}

func (m *MsgHandle) RegisterRouter(msgID uint32, router ziface.RouterInterface) {
	if _, ok := m.APIs[msgID]; ok {
		fmt.Println("repeated router id:", msgID)
	} else {
		m.APIs[msgID] = router
		fmt.Println("register router succeed id:", msgID)
	}
}

func (m *MsgHandle) AsyncTaskQueue(req ziface.RequestInterface) {
	workerID := req.GetConn().GetConnID() % m.WorkerPoolSize
	fmt.Printf("AsyncTask conn id:%d worker id:%d \n", req.GetConn().GetConnID(), workerID)
	m.TaskQueue[workerID] <- req
}
