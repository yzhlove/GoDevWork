package znet

import (
	"fmt"
	"log"
	"zinx/config"
	"zinx/utils"
	"zinx/ziface"
)

type MsgHandle struct {
	apis           map[uint32]ziface.RouterImp
	workerPoolSize uint32
	taskQueue      []chan ziface.ReqImp
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		apis:           make(map[uint32]ziface.RouterImp),
		workerPoolSize: config.GlobalConfig.WorkerPoolSize,
		taskQueue:      make([]chan ziface.ReqImp, config.GlobalConfig.WorkerPoolSize),
	}
}

func (m *MsgHandle) worker(id int, reqQueue chan ziface.ReqImp) {
	log.Println("[worker] id:", id, " is start.")
	for req := range reqQueue {
		m.Do(req)
	}
	log.Println("[worker] id:", id, " is stop.")
}

func (m *MsgHandle) RunWorkerPool() {
	for i := 0; i < int(m.workerPoolSize); i++ {
		m.taskQueue[i] = make(chan ziface.ReqImp, config.GlobalConfig.MaxWorkerTaskSize)
		go m.worker(i, m.taskQueue[i])
	}
}

func (m *MsgHandle) Do(req ziface.ReqImp) {
	defer func() {
		if err := recover(); err != nil {
			errMsg := utils.Trace(fmt.Sprint(err))
			log.Println(errMsg)
		}
	}()
	if router, ok := m.apis[req.GetMsgID()]; ok {
		router.BeforeDo(req)
		router.Handle(req)
		router.AfterDo(req)
	} else {
		log.Printf("apis msg id:%v not found", req.GetMsgID())
	}
}

func (m *MsgHandle) Register(msgID uint32, router ziface.RouterImp) {
	if _, ok := m.apis[msgID]; ok {
		log.Println("repeated register router id:", msgID)
	} else {
		m.apis[msgID] = router
	}
}

func (m *MsgHandle) AsyncTaskQueue(req ziface.ReqImp) {
	workerId := req.GetConn().GetConnID() % m.workerPoolSize
	log.Printf("[AsyncTask] conn id:{%d} woker id{%d} \n", req.GetConn().GetConnID(), workerId)
	m.taskQueue[workerId] <- req
}