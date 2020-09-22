package handler

import (
	"context"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"micro_agent/misc/packet"
	"micro_agent/proto"
	"micro_agent/sess"
)

//玩家登陆过程
func UserLoginReq(s *sess.Session, reader *packet.Packet) []byte {
	//TODO:登陆鉴权
	//简单的鉴权可以直接在agent直接完成，通常公司都存在一个用户中心服务器用于鉴权
	s.UserID = 1

	//TODO:选择GAME服务器
	//选择策略依据业务进行，小服可以选择固定某台，大服可以采用hash或者一致性hash
	s.GSID = DefaultGSID

	//连接到以选定的服务器
	conn, err := grpc.Dial(":4399", grpc.WithInsecure())
	if err != nil {
		log.Error(err)
		return failed(s, errors.New("grpc dail err:"+err.Error()))
	}

	cli := proto.NewGameServiceClient(conn)

	//开启到游戏服的流
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{
		"userid": fmt.Sprint(s.UserID),
	}))

	if s.Stream, err = cli.Stream(ctx); err != nil {
		log.Error(err)
		return failed(s, errors.New("stream err:"+err.Error()))
	}
	go read_task(s)
	return succeed(s, UserSnapshot{Uid: s.UserID})
}

//读取Game返回的消息
func read_task(s *sess.Session) {
	for {
		if in, err := s.Stream.Recv(); err != nil {
			log.Error(err)
			return
		} else {
			select {
			case s.MQ <- *in:
			case <-s.Die:
			}
		}
	}
}
