package timer

import (
	log "github.com/sirupsen/logrus"
	"micro_agent/bf"
	"micro_agent/sess"
)

var _rpm int

func InitRPM(rpm int) {
	_rpm = rpm
}

//玩家一分钟定时器
func MinuteWorker(s *sess.Session, out *bf.Buffer) {
	defer func() {
		s.PacketCount1Min = 0
	}()
	if s.PacketCount1Min > _rpm {
		s.Flag |= sess.SESS_KICKED_OUT
		log.WithFields(log.Fields{
			"userid":  s.UserID,
			"count1m": s.PacketCount1Min,
			"total":   s.PacketCount,
		}).Error("RPM")
		return
	}
}
