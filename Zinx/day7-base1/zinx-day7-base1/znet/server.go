package znet

import (
	"fmt"
	"net"
	"zinx-day7-base1/config"
	"zinx-day7-base1/ziface"
)

type TcpServer struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
	handler   ziface.MsgHandleInterface
	manager   ziface.ConnManagerInterface
	stopCh    chan struct{}
	//server 连接时 hook
	onConnStart ziface.CallbackConnFunc
	//server 断开时 hook
	onConnStop ziface.CallbackConnFunc
}

func NewTcpServer() ziface.ServerInterface {
	server := &TcpServer{
		Name:      config.GlobalConfig.ServerName,
		IPVersion: "tcp4",
		IP:        config.GlobalConfig.Host,
		Port:      config.GlobalConfig.TcpPort,
		handler:   NewMsgHandle(),
		manager:   NewConnManager(),
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

				if s.manager.Len() >= config.GlobalConfig.MaxConn {
					conn.Close()
					continue
				}

				context := NewConnContext(conn, cid, s)
				cid++
				go context.Start()
			}
		}
	}()
}

func (s *TcpServer) Stop() {
	fmt.Println("server stop running ,name ", s.Name)
	s.manager.ClearConn()
	close(s.stopCh)
}

func (s *TcpServer) Run() {
	s.Start()
	<-s.stopCh
}

func (s *TcpServer) RegisterRouter(msgID uint32, router ziface.RouterInterface) {
	s.handler.RegisterRouter(msgID, router)
}

func (s *TcpServer) GetConnManager() ziface.ConnManagerInterface {
	return s.manager
}

func (s *TcpServer) GetMsgHandle() ziface.MsgHandleInterface {
	return s.handler
}

func (s *TcpServer) SetOnConnStart(fn ziface.CallbackConnFunc) {
	s.onConnStart = fn
}

func (s *TcpServer) SetOnConnStop(fn ziface.CallbackConnFunc) {
	s.onConnStop = fn
}

func (s *TcpServer) CallOnConnStart(conn ziface.ConnectionInterface) {
	if s.onConnStart != nil {
		fmt.Println("--> CallOnConnStart ...")
		s.onConnStart(conn)
	}
}

func (s *TcpServer) CallOnConnStop(conn ziface.ConnectionInterface) {
	if s.onConnStop != nil {
		fmt.Printf("--> CallonConnStop ...")
		s.onConnStop(conn)
	}
}
