package router

import (
	log "github.com/sirupsen/logrus"
	"micro_agent/build"
	"micro_agent/handler"
	"micro_agent/misc/packet"
	"micro_agent/sess"
	"micro_agent/utils"
	"time"
)

func Router(s *sess.Session, p []byte) []byte {
	start := time.Now()
	defer utils.Trace(s, p)

	//解密
	if s.Flag&sess.SESS_ENCRYPT != 0 {
		s.Decoder.XORKeyStream(p, p)
	}

	log.Printf("[read bytes] => %v \n", p)

	//封装reader
	reader := packet.Reader(p)

	//读客户端数据包序列号(1,2,3...)
	//客户端发送的数据包必须包含一个自增的序号，必须严格递增
	//加密后，可避免重放攻击-REPLAY-ATTACK
	seqId, err := reader.ReadU32()
	if err != nil {
		log.Error("read client timestamp failed:", err)
		s.Flag |= sess.SESS_KICKED_OUT
		return nil
	}

	//数据包序列号验证
	if seqId != s.PacketCount {
		log.Errorf("packet sequence id: %v should be:%v size:%v ", seqId, s.PacketCount, len(p)-6)
		s.Flag |= sess.SESS_KICKED_OUT
		return nil
	}

	log.Info("read seqId => ", seqId)

	//读取协议号
	pid, err := reader.ReadS16()
	if err != nil {
		log.Error("read protocol number failed.")
		s.Flag |= sess.SESS_KICKED_OUT
		return nil
	}

	log.Info("read pid => ", pid)

	//设置协议号
	s.LastReqId = pid

	//根据协议号段做服务划分
	//协议号的划分采用分割协议区间，用户可自定义多个区间，用于转发到不同的后端服务
	var ret []byte
	if pid > 1000 {
		if err := build.Forward(s, p[4:]); err != nil {
			log.Errorf("service id:%v execute failed,error:%v", pid, err)
			s.Flag |= sess.SESS_KICKED_OUT
			return nil
		}
	} else {
		if event, ok := handler.Handlers[pid]; ok && event != nil {
			ret = event(s, reader)
		} else {
			log.Errorf("service id:%v not bind ", pid)
			s.Flag |= sess.SESS_KICKED_OUT
			return nil
		}
	}

	elasped := time.Now().Sub(start)
	//排除心跳包日志 (pid == 0)
	if pid != 0 {
		log.WithFields(log.Fields{
			"api":  "======",
			"code": elasped,
		}).Debug("REQ")
	}
	return ret
}
