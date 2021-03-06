package core

import (
	"github.com/golang/protobuf/proto"
	"log"
	"math/rand"
	"snowflake"
	"zinx-game-example/mmo/pb"
	"zinx/ziface"
	"zinx/zlog"
)

type Player struct {
	PID  int32
	Conn ziface.ConnImp
	X    float32
	Y    float32
	Z    float32
	V    float32 //旋转0-360
}

func NewPlayer(conn ziface.ConnImp) *Player {
	return &Player{
		PID:  int32(snowflake.Get() % 10000),
		Conn: conn,
		X:    float32(160 + rand.Intn(20)),
		Y:    0,
		Z:    float32(134 + rand.Intn(20)),
		V:    0,
	}
}

func (p *Player) Send(msgID uint32, data proto.Message) {
	zlog.Infof("before player message %+v \n", data)
	msg, err := proto.Marshal(data)
	if err != nil {
		zlog.Error("marshal data message err:", err)
		return
	}
	zlog.Infof("after player message %+v \n", msg)
	if p.Conn == nil {
		log.Println("connection in player is nil.")
		return
	}
	if err := p.Conn.Send(msgID, msg); err != nil {
		zlog.Info("player send message error:", err)
		return
	}
}

func (p *Player) SyncPid() {
	p.Send(1, &pb.SyncPid{Pid: p.PID})
}

func (p *Player) BroadCastStartPosition() {
	p.Send(200, &pb.BroadCast{
		Pid: p.PID,
		Tp:  2,
		Data: &pb.BroadCast_P{P: &pb.Position{
			X: p.X,
			Y: p.Y,
			Z: p.Z,
			V: p.V,
		}},
	})
}

func (p *Player) Talk(content string) {
	msg := &pb.BroadCast{Pid: p.PID, Tp: 1,
		Data: &pb.BroadCast_Content{Content: content},
	}
	for _, player := range WorldMgr.GetPlayers() {
		player.Send(200, msg)
	}
}

func (p *Player) SyncRangePlayers() {
	pids := WorldMgr.AoiMgr.GetPlayerIDS(p.X, p.Z)
	players := make([]*Player, 0, len(pids))
	for _, pid := range pids {
		players = append(players, WorldMgr.GetPlayer(int32(pid)))
	}

	msg := &pb.BroadCast{
		Pid: p.PID,
		Tp:  2,
		Data: &pb.BroadCast_P{
			P: &pb.Position{
				X: p.X, Y: p.Y, Z: p.Z, V: p.V}},
	}

	for _, player := range players {
		player.Send(200, msg)
	}

	playersData := make([]*pb.Player, 0, len(players))
	for _, player := range players {
		data := &pb.Player{
			Pid: player.PID,
			P:   &pb.Position{X: player.X, Y: player.Y, Z: player.Z, V: player.V},
		}
		playersData = append(playersData, data)
	}
	p.Send(202, &pb.SyncPlayers{Ps: playersData})
}

func (p *Player) UpdatePos(x, y, z, v float32) {
	p.X, p.Y, p.Z, p.V = x, y, z, v

	msg := &pb.BroadCast{
		Pid: p.PID,
		Tp:  4,
		Data: &pb.BroadCast_P{
			P: &pb.Position{X: p.X, Y: p.Y, Z: p.Z, V: p.V},
		},
	}
	for _, player := range p.GetRandPlayers() {
		player.Send(200, msg)
	}
}

func (p *Player) GetRandPlayers() []*Player {
	pids := WorldMgr.AoiMgr.GetPlayerIDS(p.X, p.Z)
	log.Printf("area players => %v \n", pids)
	players := make([]*Player, 0, len(pids))
	for _, pid := range pids {
		players = append(players, WorldMgr.GetPlayer(int32(pid)))
	}
	return players
}

func (p *Player) LostConn() {
	players := p.GetRandPlayers()
	msg := &pb.SyncPid{Pid: p.PID}
	for _, player := range players {
		player.Send(201, msg)
	}
	WorldMgr.AoiMgr.DelFormGridByPos(int(p.PID), p.X, p.Z)
	WorldMgr.DelPlayer(p.PID)
}
