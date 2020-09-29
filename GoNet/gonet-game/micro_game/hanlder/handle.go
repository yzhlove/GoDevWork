package handler

import (
	log "github.com/sirupsen/logrus"
	"micro_game/misc/packet"
	"micro_game/sess"
)

func succeed(s *sess.Session, value interface{}) []byte {
	if r, ok := Result[s.LastReqId]; ok {
		if r[0] > 0 {
			return packet.Pack(r[0], value, nil)
		}
		return packet.Pack(0, "server send succeed invalid message", nil)
	}
	log.Error("succeed send error code:", s.LastReqId)
	return nil
}

func failed(s *sess.Session, value interface{}) []byte {
	if r, ok := Result[s.LastReqId]; ok {
		if r[1] > 0 {
			return packet.Pack(r[0], value, nil)
		}
		return packet.Pack(0, "server send succeed invalid message", nil)
	}
	log.Error("succeed send error code:", s.LastReqId)
	return nil
}
