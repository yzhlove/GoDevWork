package core

import (
	"github.com/golang/protobuf/proto"
	"log"
	"math/rand"
	"sync"
	"zinx-game-example/mmo/pb"
	"zinx/ziface"
)

type Player struct {
	PID  int32
	Conn ziface.ConnImp
	X    float32
	Y    float32
	Z    float32
	V    float32 //旋转0-360
}

var GenPID int32 = 1 //ID生成
var mutex sync.RWMutex

func GeneratePID() int32 {
	mutex.Lock()
	defer mutex.Unlock()
	t := GenPID
	GenPID++
	return t
}

func NewPlayer(conn ziface.ConnImp) *Player {
	return &Player{
		PID:  GeneratePID(),
		Conn: conn,
		X:    float32(160 + rand.Intn(10)),
		Y:    0,
		Z:    float32(134 + rand.Intn(17)),
		V:    0,
	}
}

func (p *Player) Send(msgID uint32, data proto.Message) {
	log.Printf("before player message %+v \n", data)
	msg, err := proto.Marshal(data)
	if err != nil {
		log.Println("marshal data message err:", err)
		return
	}
	log.Printf("after player message %+v \n", msg)
	if p.Conn == nil {
		log.Println("connection in player is nil.")
		return
	}
	if err := p.Conn.Send(msgID, msg); err != nil {
		log.Println("player send message error:", err)
		return
	}
}

func (p *Player) SyncPid() {
	p.Send(1, &pb.SyncPid{Pid: p.PID})
}

func (p *Player) BroadCastStartPostion() {
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
