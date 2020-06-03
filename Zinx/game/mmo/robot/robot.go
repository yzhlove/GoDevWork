package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"time"
	"zinx-game-example/mmo/pb"
)

type Msg struct {
	Len  uint32
	ID   uint32
	Data []byte
}

type TcpClient struct {
	conn       net.Conn
	X, Y, Z, V float32
	PID        int32
	isOnline   chan struct{}
}

func (tcp *TcpClient) Unpack(data []byte) (*Msg, error) {
	buf := bytes.NewReader(data)
	msg := &Msg{}
	if err := binary.Read(buf, binary.LittleEndian, &msg.Len); err != nil {
		return nil, err
	}
	if err := binary.Read(buf, binary.LittleEndian, &msg.ID); err != nil {
		return nil, err
	}
	return msg, nil
}

func (tcp *TcpClient) Pack(msgID uint32, data []byte) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	if err := binary.Write(buf, binary.LittleEndian, uint32(len(data))); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, msgID); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (tcp *TcpClient) send(msgID uint32, data proto.Message) {
	ptoData, err := proto.Marshal(data)
	if err != nil {
		log.Println("marshal err:", err)
		return
	}
	pack, err := tcp.Pack(msgID, ptoData)
	if err != nil {
		log.Println(err)
		return
	}
	tcp.conn.Write(pack)
}

func (tcp *TcpClient) RobotAction() {
	switch rand.Intn(2) {
	case 0: //聊天
		tcp.send(2, &pb.Talk{Content: fmt.Sprintf("hello player:%d who are you.", tcp.PID)})
	case 1: //移动
		x, z, v := tcp.X, tcp.Z, tcp.V
		switch rand.Intn(2) {
		case 0:
			x, z = x-float32(rand.Intn(10)), z-float32(rand.Intn(10))
		case 1:
			x, z = x+float32(rand.Intn(10)), z+float32(rand.Intn(10))
		}
		switch {
		case x > 400:
			x = 400
		case x < 85:
			x = 85
		case z > 400:
			z = 400
		case z < 75:
			z = 75
		}
		switch rand.Intn(2) {
		case 0:
			v = 25
		case 1:
			v = 0
		}
		log.Printf("player id:%d is walking...\n", tcp.PID)
		tcp.send(3, &pb.Position{X: x, Y: tcp.Y, Z: z, V: v})
	}
}

func (tcp *TcpClient) Do(msg *Msg) {
	switch msg.ID {
	case 1:
		syncPID := &pb.SyncPid{}
		if err := proto.Unmarshal(msg.Data, syncPID); err != nil {
			log.Println("unmarshal err:", err)
		} else {
			tcp.PID = syncPID.Pid
		}
	case 200:
		data := &pb.BroadCast{}
		if err := proto.Unmarshal(msg.Data, data); err == nil {
			switch data.Tp {
			case 1:
				log.Printf("talk %d content: %v \n", data.Pid, data.GetContent())
			case 2:
				if data.Pid == tcp.PID {
					pos := data.GetP()
					tcp.X, tcp.Y, tcp.Z, tcp.V = pos.X, pos.Y, pos.Z, pos.V
					log.Printf("player id:%d online at(%f %f %f %f) \n",
						data.Pid, pos.X, pos.Y, pos.Z, pos.V)
					tcp.isOnline <- struct{}{}
				}
			}
		}
	}
}

func (tcp *TcpClient) Start() {
	go func() {
		for {
			head := make([]byte, 8)
			if _, err := io.ReadFull(tcp.conn, head); err != nil {
				log.Println("read head err:", err)
				return
			}
			pkgHead, err := tcp.Unpack(head)
			if err != nil {
				log.Println("unpack err:", err)
				return
			}
			if pkgHead.Len > 0 {
				pkgHead.Data = make([]byte, pkgHead.Len)
				if _, err := io.ReadFull(tcp.conn, pkgHead.Data); err != nil {
					log.Println("read pack data err:", err)
					return
				}
			}
			log.Println("read message ID => ", pkgHead.ID)
			tcp.Do(pkgHead)
		}
	}()
	select {
	case <-tcp.isOnline:
		go func() {
			tcp.RobotAction()
			time.Sleep(time.Second * 3)
		}()
	}
}

func NewTcpClient(ip string, port int) *TcpClient {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		panic(err)
	}
	return &TcpClient{conn: conn, isOnline: make(chan struct{})}
}

func main() {
	for i := 0; i < 500; i++ {
		cli := NewTcpClient("127.0.0.1", 7777)
		cli.Start()
		time.Sleep(time.Second * 2)
	}
	sign := make(chan os.Signal, 1)
	signal.Notify(sign, os.Interrupt, os.Kill)
	sig := <-sign
	log.Println("signal ==> ", sig)
}
