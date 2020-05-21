package znet

import (
	"fmt"
	"net"
	"time"
	"zinx-day2-base1/ziface"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
	Router    ziface.RouterInterface
}

func NewServer(name string) *Server {
	return &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      7777,
	}
}

func (s *Server) Start() {
	fmt.Printf("[START] server listenner at IP %s:%d ,is starting ... \n", s.IP, s.Port)
	go func() {
		tcpAddr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("[Server] resolve tcp add err :", err)
			return
		}
		listener, err := net.ListenTCP(s.IPVersion, tcpAddr)
		if err != nil {
			fmt.Printf("[Server] listener IPVersion %s err : %s \n", s.IPVersion, err.Error())
			return
		}
		fmt.Println("[Server] start zinx ", s.Name, " succeed is running ...")
		var cid uint32
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("[Server] accept err", err)
				continue
			}
			handleConn := NewConn(conn, cid, s.Router)
			cid++
			go handleConn.Start()
		}
	}()
}

func (s *Server) Stop() {
	fmt.Println("[Server] stop zinx ,name ", s.Name)
}

func (s *Server) Run() {
	s.Start()
	for {
		time.Sleep(time.Second * 10)
	}
}

func (s *Server) RegisterRouter(router ziface.RouterInterface) {
	if s.Router == nil {
		s.Router = router
	} else {
		panic("[Server] register router err ")
	}
	fmt.Println("[Server] register router succeed!")
}
