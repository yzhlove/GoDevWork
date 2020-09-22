package handler

import (
	"fmt"
	"micro_agent/misc/packet"
	"micro_agent/sess"
)

func HeartBeatReq(s *sess.Session, reader *packet.Packet) []byte {
	if tbl, err := PacketAutoId(reader); err != nil {
		return failed(s, fmt.Sprintf("heart beat error:%v", err))
	} else {
		return succeed(s, tbl)
	}
}
