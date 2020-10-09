package service

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
	"io"
	"micro_game/config"
	handler "micro_game/hanlder"
	"micro_game/misc/packet"
	"micro_game/proto"
	"micro_game/registry"
	"micro_game/sess"
	"micro_game/utils"
	"strconv"
)

var (
	ErrorIncorrectFrameType = errors.New("incorrect frame type")
	ErrorServiceNotBind     = errors.New("service not bind")
)

type GameService struct{}

func (s *GameService) recv(stream proto.GameService_StreamServer, die chan struct{}) chan *proto.Game_Frame {
	ch := make(chan *proto.Game_Frame, 1)
	go func() {
		defer func() { close(ch) }()
		for {
			if in, err := stream.Recv(); err != nil {
				if errors.Is(err, io.EOF) {
					return
				}
				log.Error(err)
				return
			} else {
				select {
				case ch <- in:
				case <-die:
				}
			}
		}
	}()
	return ch
}

func (s *GameService) Stream(stream proto.GameService_StreamServer) error {
	defer utils.Trace()
	var session sess.Session
	die := make(chan struct{})
	agent := s.recv(stream, die)
	ipc := make(chan *proto.Game_Frame, config.DefaultChanIpcSize)

	defer func() {
		registry.UnRegister(session.UserId, ipc)
		close(die)
		log.Debug("stream end:", session.UserId)
	}()

	meta, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		log.Error("cannot read metadata from metadata")
		return ErrorIncorrectFrameType
	}

	if len(meta["user_id"]) == 0 {
		log.Error("cannot read key:userid from metadata")
		return ErrorIncorrectFrameType
	}

	userId, err := strconv.Atoi(meta["user_id"][0])
	if err != nil {
		log.Error(err)
		return ErrorIncorrectFrameType
	}

	session.UserId = int32(userId)
	registry.Register(session.UserId, ipc)
	log.Info("user_id:", session.UserId, " logged in")

	for {
		select {
		case frame, ok := <-agent:
			if !ok {
				return nil
			}
			switch frame.Type {
			case proto.Game_Message:
				reader := packet.Reader(frame.Message)
				c, err := reader.ReadS16()
				if err != nil {
					log.Error(err)
					return err
				}
				var handle handler.HandleFunc
				if handle = handler.Handlers[c]; handle == nil {
					log.Error("service not bind:", c)
					return ErrorServiceNotBind
				}
				if ret := handle(&session, reader); ret != nil {
					if err := stream.Send(&proto.Game_Frame{Type: proto.Game_Message, Message: ret}); err != nil {
						log.Error(err)
						return err
					}
				}
				if session.Flag&sess.SESS_KICKED_OUT != 0 {
					if err := stream.Send(&proto.Game_Frame{Type: proto.Game_Kick}); err != nil {
						log.Error(err)
						return err
					}
					return nil
				}
			case proto.Game_Ping:
				if err := stream.Send(&proto.Game_Frame{Type: proto.Game_Ping, Message: frame.Message}); err != nil {
					log.Error(err)
					return err
				}
			default:
				log.Error("incorrect frame type:", frame.Type)
				return ErrorIncorrectFrameType
			}
		case frame := <-ipc:
			if err := stream.Send(frame); err != nil {
				log.Error(err)
				return err
			}
		}
	}

}
