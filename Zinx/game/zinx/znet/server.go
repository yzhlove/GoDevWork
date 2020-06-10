package znet

import (
	"fmt"
	"net"
	"zinx/config"
	"zinx/discovery"
	"zinx/ziface"
	"zinx/zlog"
)

type TcpServer struct {
	Name        string
	IPVersion   string
	IP          string
	Port        int
	msgHandle   ziface.MsgHandleImp
	connMgr     ziface.ConnMgrImp
	die         chan struct{}
	onConnStart ziface.EventFunc
	onConnStop  ziface.EventFunc
}

func NewTcpServer() ziface.ServerImp {
	return &TcpServer{
		Name:      config.GlobalConfig.ServerName,
		IPVersion: "tcp4",
		IP:        config.GlobalConfig.Host,
		Port:      config.GlobalConfig.TcpPort,
		msgHandle: NewMsgHandle(),
		connMgr:   NewConnMgr(),
		die:       make(chan struct{}),
	}
}

func (s *TcpServer) Start() {
	zlog.Infof("server start on listener at running %s:%d \n", s.IP, s.Port)
	go func() {
		defer s.Stop()
		s.msgHandle.RunWorkerPool()
		tcpAddr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			zlog.Error("resolve tcp addr err:", err)
			return
		}
		listen, err := net.ListenTCP(s.IPVersion, tcpAddr)
		if err != nil {
			zlog.Error("listen tcp addr err:", err)
			return
		}
		zlog.Infof("start tcp server name:{%s}\n", s.Name)
		var connID uint32
		for {
			conn, err := listen.AcceptTCP()
			if err != nil {
				zlog.Error("accept err:", err)
				continue
			}
			if s.connMgr.Size() >= config.GlobalConfig.MaxConnections {
				zlog.Error("conn overflow limit")
				conn.Close()
			}
			c := NewConn(conn, connID, s)
			connID++
			go c.Start()
		}
	}()
}

func (s *TcpServer) Stop() {
	zlog.Info("server stop .")
	s.connMgr.Clear()
	close(s.die)
}

func (s *TcpServer) Run() {
	s.Start()
	if err := s.registerService(); err != nil {
		zlog.Error("register service err:", err)
	}
	<-s.die
}

func (s *TcpServer) Register(msgID uint32, router ziface.RouterImp) {
	s.msgHandle.Register(msgID, router)
}

func (s *TcpServer) GetConnMgr() ziface.ConnMgrImp {
	return s.connMgr
}

func (s *TcpServer) GetMsgHandle() ziface.MsgHandleImp {
	return s.msgHandle
}

func (s *TcpServer) ConnStartEvent(event ziface.EventFunc) {
	s.onConnStart = event
}

func (s *TcpServer) ConnStopEvent(event ziface.EventFunc) {
	s.onConnStop = event
}

func (s *TcpServer) CallbackConnStart(conn ziface.ConnImp) {
	if s.onConnStart != nil {
		s.onConnStart(conn)
	}
}

func (s *TcpServer) CallbackConnStop(conn ziface.ConnImp) {
	if s.onConnStop != nil {
		s.onConnStop(conn)
	}
}

func (s *TcpServer) registerService() error {
	etcd, err := discovery.NewEtcd()
	if err != nil {
		return err
	}
	etcd.Register(s.Name, fmt.Sprintf("%s:%d", s.IP, s.Port))
	return nil
}
