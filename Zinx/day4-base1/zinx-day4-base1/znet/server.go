package znet

import (
	"fmt"
	"net"
	"time"
	"zinx-day4-base1/config"
	"zinx-day4-base1/ziface"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
	handler   ziface.MessageHandleInterface
}

func NewServer() ziface.ServerInterface {
	config.GlobalConfig.Reload()

	server := &Server{
		Name:      config.GlobalConfig.Name,
		IPVersion: "tcp4",
		IP:        config.GlobalConfig.Host,
		Port:      config.GlobalConfig.TcpPort,
		handler:   NewMessageHandle(),
	}
	return server
}

func (s *Server) Start() {
	fmt.Printf("[START] server listener at running %s:%d .\n", s.IP, s.Port)
	go func() {
		tcpAddr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("[SERVER] resolve tcp addr err:", err)
			return
		}
		listener, err := net.ListenTCP(s.IPVersion, tcpAddr)
		if err != nil {
			fmt.Println("[SERVER] listener err:", err)
			return
		}
		fmt.Printf("[SERVER] start zinx server by name :%s \n", s.Name)
		var cid uint32
		for {
			if conn, err := listener.AcceptTCP(); err != nil {
				fmt.Println("[SERVER] accept err:", err)
			} else {
				handle := NewConn(conn, cid, s.handler)
				cid++
				go handle.Start()
			}
		}
	}()
}

func (s *Server) Stop() {
	fmt.Println("[SERVER] stop zinx server , name by ", s.Name)
}

func (s *Server) Run() {
	s.Start()
	for {
		time.Sleep(time.Second * 10)
	}
}

func (s *Server) RegisterRouter(msgID uint32, router ziface.RouterInterface) {
	s.handler.RegisterRouter(msgID, router)
}
