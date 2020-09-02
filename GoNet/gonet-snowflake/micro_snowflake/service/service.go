package service

import (
	"context"
	"errors"
	"log"
	"micro_snowflake/config"
	"micro_snowflake/proto"
	"time"
)

const (
	MaxQueue = 1024
)

const (
	TsMask        = 0x1FFFFFFFFFF // 41bit
	SnMask        = 0xFFF         // 12bit
	MachineIdMask = 0x3FF         // 10bit
)

var errNoUID = errors.New("generate uid error")

type Server struct {
	queue     chan chan uint64
	machineId uint64 //10 bit
}

func (s *Server) Init(cfg *config.Config) {
	s.queue = make(chan chan uint64, MaxQueue)
	s.machineId = uint64(cfg.MachineID&MachineIdMask) << 12
	go s.produce()
}

// GetUID Grpc获取UID
func (s *Server) GetUID(ctx context.Context, _ *proto.Sf_Nil) (*proto.Sf_UID, error) {
	ret := make(chan uint64, 1)
	s.queue <- ret
	if uid, ok := <-ret; ok {
		close(ret)
		return &proto.Sf_UID{Uid: uid}, nil
	}
	return nil, errNoUID
}

// produce 生产UID
func (s *Server) produce() {
	var sn uint64
	var lt int64
	for {
		if ret, ok := <-s.queue; ok {
			ct := ts()
			if ct < lt {
				log.Println("timestamp is failed")
				ct = s.wait(lt)
			}
			//毫秒级时间戳不一致，计数器归0
			if ct != lt {
				sn = 0
			} else {
				if sn = (sn + 1) & SnMask; sn == 0 {
					ct = s.wait(lt)
				}
			}
			lt = ct
			ret <- pack(ct, s.machineId, sn)
		} else {
			panic("chan is closed")
		}
	}
}

func (s *Server) wait(lt int64) int64 {
	var t int64
	for t = ts(); t < lt; t = ts() {
		time.Sleep(time.Duration(lt-t) * time.Millisecond)
	}
	return t
}

func pack(t int64, id, sn uint64) uint64 {
	var uid uint64
	log.Println(t, id, sn)
	uid |= (uint64(t) & TsMask) << 22
	return uid | id | sn
}

func ts() int64 {
	return time.Now().UnixNano() / 1e6
}
