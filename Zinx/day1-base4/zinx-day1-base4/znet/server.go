package znet

import (
	"errors"
	"fmt"
	"net"
	"time"
	"zinx-day1-base4/ziface"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
}

func handleLogic(conn *net.TCPConn, data []byte, cnt int) error {
	fmt.Println("[Conn Handle] Callback to client ...")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err ", err)
		return errors.New("handleLogic Error " + err.Error())
	}
	return nil
}

func (s *Server) Start() {
	fmt.Printf("[START] Server listenner at IP:%s ,Port %d is starting ... ", s.IP, s.Port)
	go func() {
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err: ", err)
			return
		}
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen ", s.IPVersion, " err ", err)
			return
		}
		fmt.Println("start ZINX server ", s.Name, " succeed ,now listening ...")
		var cid uint32
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err:", err)
				continue
			}
			handleConn := NewConnection(conn, cid, handleLogic)
			cid++
			go handleConn.Start()
		}
	}()
}

func (s *Server) Stop() {
	fmt.Println("[STOP] ZINX server ,name ", s.Name)
}

func (s *Server) Serve() {
	s.Start()
	for {
		time.Sleep(time.Second * 10)
	}
}

func NewServer(name string) ziface.ServerInterface {
	return &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      1234,
	}
}
