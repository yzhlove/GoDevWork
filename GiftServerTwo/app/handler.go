package app

import (
	"WorkSpace/GoDevWork/GiftServerTwo/manager"
	"WorkSpace/GoDevWork/GiftServerTwo/obj"
	"WorkSpace/GoDevWork/GiftServerTwo/pb"
	log "github.com/sirupsen/logrus"
	"sync"
)

type ServerStream = pb.GiftService_SyncServer

type handler struct {
	streams map[uint32]ServerStream
	reqChan chan uint32
	entity  *obj.Entity
	mutex   sync.RWMutex
}

func (h *handler) Init() (err error) {
	h.streams = make(map[uint32]ServerStream, 4)
	h.reqChan = make(chan uint32, 128)

	if h.entity, err = manager.EntityInit(); err != nil {
		return err
	}

	go h.sync()
	return
}

func (h *handler) setConn(z uint32, s ServerStream) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	h.streams[z] = s
}

func (h *handler) getConn(z uint32) (s ServerStream, ok bool) {
	h.mutex.RLock()
	defer h.mutex.RUnlock()
	s, ok = h.streams[z]
	return
}

func (h *handler) sync() {

	for {
		select {
		case z, ok := <-h.reqChan:
			if !ok {
				break
			}
			if s, ok := h.getConn(z); ok {
				for _, code := range h.entity.Infos {
					if code.ZoneCheck(z) {
						if err := s.Send(manager.GeneratePtoCodeInfo(code)); err != nil {
							log.Error("send err:", err)
						}
					}
				}
			}
		}
	}
}

func (h *handler) Sync(code *obj.Code) {

	msg := manager.GeneratePtoCodeInfo(code)
	if len(code.ZoneIds) == 0 {
		for _, s := range h.streams {
			if err := s.Send(msg); err != nil {
				log.Error("send err:", err)
			}
		}
	} else {
		for _, z := range code.ZoneIds {
			if s, ok := h.getConn(z); ok {
				if err := s.Send(msg); err != nil {
					log.Error("send err:", err)
				}
			}
		}
	}
}

func (h *handler) UpdateEntity(code *obj.Code) {
	h.entity.Update(code)
}
