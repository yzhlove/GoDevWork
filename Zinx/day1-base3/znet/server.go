package znet

import (
	"WorkSpace/GoDevWork/Zinx/day1-base3/ziface"
	"fmt"
	"net"
	"time"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
}

func (s *Server) Start() {
	fmt.Printf("[START] Server listener at IP: %s,Port %d,is starting \n", s.IP, s.Port)
	go func() {
		tcpAddr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err:", err)
			return
		}
		listener, err := net.ListenTCP(s.IPVersion, tcpAddr)
		if err != nil {
			fmt.Println("listener ", s.IPVersion, " err", err)
			return
		}
		fmt.Println("\nstart ZINX server ", s.Name, " succeed,now listening...")
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err ", err)
				continue
			}
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("recv buf err ", err)
						continue
					}
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write back buf err ", err)
						continue
					}
				}
			}()
		}
	}()
}

func (s *Server) Stop() {
	fmt.Println("[STOP] ZINX server ,name ", s.Name)
}

func (s *Server) Serve() {
	s.Start()

	//阻塞
	for {
		time.Sleep(10)
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
