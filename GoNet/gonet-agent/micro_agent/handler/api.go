package handler

import (
	"micro_agent/misc/packet"
	"micro_agent/sess"
)

type HandleFunc func(s *sess.Session, reader *packet.Packet) []byte

var Handlers map[int16]HandleFunc

func init() {
	Handlers = map[int16]HandleFunc{

	}
}
