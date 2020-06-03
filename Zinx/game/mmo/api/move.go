package api

import (
	"github.com/golang/protobuf/proto"
	"log"
	"zinx-game-example/mmo/core"
	"zinx-game-example/mmo/pb"
	"zinx/ziface"
	"zinx/znet"
)

type MoveApi struct {
	znet.AbstractRouter
}

func (*MoveApi) Handle(req ziface.ReqImp) {
	msg := &pb.Position{}
	if err := proto.Unmarshal(req.GetMsgData(), msg); err != nil {
		log.Println("Move:position unmarshal err:", err)
		return
	}
	pid, ok := req.GetConn().GetAttr("pid")
	if !ok {
		log.Println("get player pid error")
		return
	}
	log.Printf("player pid:%d move(%f %f %f %f)\n", pid, msg.X, msg.Y, msg.Z, msg.V)
	player := core.WorldMgr.GetPlayer(pid.(int32))
	player.UpdatePos(msg.X, msg.Y, msg.Z, msg.V)
}
