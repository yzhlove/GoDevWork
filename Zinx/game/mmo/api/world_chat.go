package api

import (
	"github.com/golang/protobuf/proto"
	"zinx-game-example/mmo/core"
	"zinx-game-example/mmo/pb"
	"zinx/ziface"
	"zinx/zlog"
	"zinx/znet"
)

type WorldChatApi struct {
	znet.AbstractRouter
}

func (*WorldChatApi) Handle(req ziface.ReqImp) {
	msg := &pb.Talk{}
	if err := proto.Unmarshal(req.GetMsgData(), msg); err != nil {
		zlog.Info("talk unmarshal err:", err)
		return
	}
	pid, ok := req.GetConn().GetAttr("pid")
	if !ok {
		zlog.Info("get attr err not found pid")
		req.GetConn().Stop()
		return
	}
	if player := core.WorldMgr.GetPlayer(pid.(int32)); player != nil {
		player.Talk(msg.Content)
	}
}
