package agent

import (
	"micro_agent/bf"
	"micro_agent/proto"
	"micro_agent/sess"
	"micro_agent/signal"
	"micro_agent/timer"
	"micro_agent/utils"
	"time"
)

func Agent(s *sess.Session, in chan []byte, out *bf.Buffer) {
	defer signal.WaitGroup.Done()
	defer utils.Trace()

	s.MQ = make(chan proto.Game_Frame, 512)
	s.ConnectTime = time.Now()
	s.LastPacketTime = time.Now()

	t := time.NewTicker(time.Minute)
	defer func() {
		close(s.Die)
		if s.Stream != nil {
			s.Stream.CloseSend()
		}
		t.Stop()
	}()

	for {
		select {
		case msg, ok := <-in:
			if !ok {
				return
			}
			s.PacketCount++
			s.PacketCount1Min++
			s.PacketTime = time.Now()
			//handler

			s.LastPacketTime = s.PacketTime
		case frame := <-s.MQ:
			switch frame.Type {
			case proto.Game_Message:
				out.Send(s, frame.Message)
			case proto.Game_Kick:
				s.Flag |= sess.SESS_KICKED_OUT
			}
		case <-t.C:
			timer.MinuteWorker(s, out)
		case <-signal.Die:
			s.Flag |= sess.SESS_KICKED_OUT
		}
		if s.Flag&sess.SESS_KICKED_OUT != 0 {
			return
		}
	}

}
