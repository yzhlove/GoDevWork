package znet

import (
	"fmt"
	"net"
	"zinx-day6-base1/config"
	"zinx-day6-base1/ziface"
)

type TcpServer struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
	handler   ziface.MsgHandleInterface
	stopCh    chan struct{}
}

func NewTcpServer() ziface.ServerInterface {
	server := &TcpServer{
		Name:      config.GlobalConfig.ServerName,
		IPVersion: "tcp4",
		IP:        config.GlobalConfig.Host,
		Port:      config.GlobalConfig.TcpPort,
		handler:   NewMsgHandle(),
		stopCh:    make(chan struct{}),
	}
	return server
}

func (s *TcpServer) Start() {
	fmt.Printf("[START] server lisener at running %s:%d\n", s.IP, s.Port)
	go func() {
		defer s.Stop()
		s.handler.StartWorkerPool()
		tcpAddr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err:", err)
			return
		}
		listener, err := net.ListenTCP(s.IPVersion, tcpAddr)
		if err != nil {
			fmt.Println("listener err:", err)
			return
		}
		fmt.Println("start tcp server name ", s.Name)
		var cid uint32
		for {
			if conn, err := listener.AcceptTCP(); err != nil {
				fmt.Println("accept err:", err)
			} else {
				context := NewConnContext(conn, cid, s.handler)
				cid++
				go context.Start()
			}
		}
	}()
}

func (s *TcpServer) Stop() {
	fmt.Println("server stop running ,name ", s.Name)
	close(s.stopCh)
}

func (s *TcpServer) Run() {
	s.Start()
	<-s.stopCh
}

func (s *TcpServer) RegisterRouter(msgID uint32, router ziface.RouterInterface) {
	s.handler.RegisterRouter(msgID, router)
}
