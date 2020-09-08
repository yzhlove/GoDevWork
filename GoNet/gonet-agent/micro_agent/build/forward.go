package build

import (
	"errors"
	"micro_agent/proto"
	"micro_agent/sess"
)

var ErrorStreamNotOpen = errors.New("stream not open yet")

func Forward(s *sess.Session, data []byte) error {

	if s.Stream != nil {
		msg := &proto.Game_Frame{Type: proto.Game_Message, Message: data}
		if err := s.Stream.Send(msg); err != nil {
			return err
		}
		return nil
	}
	return ErrorStreamNotOpen
}
