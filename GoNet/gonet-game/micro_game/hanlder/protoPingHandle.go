package handler

import (
	"micro_game/misc/packet"
	"micro_game/sess"
)

func ProtoPingReq(s *sess.Session, reader *packet.Packet) []byte {
	if tbl, err := PacketAutoId(reader); err != nil {
		return failed(s, err)
	} else {
		return succeed(s, tbl)
	}
}
