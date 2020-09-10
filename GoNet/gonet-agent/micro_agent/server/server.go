package server

import (
	"encoding/binary"
	log "github.com/sirupsen/logrus"
	"github.com/xtaci/kcp-go"
	"io"
	"micro_agent/agent"
	"micro_agent/bf"
	"micro_agent/config"
	"micro_agent/sess"
	"micro_agent/signal"
	"micro_agent/utils"
	"net"
	"time"
)

func TcpServer(cfg *config.Config) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", cfg.Listen)
	utils.CheckError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	utils.CheckError(err)

	log.Info("listening on:", listener.Addr())

	for {
		if conn, err := listener.AcceptTCP(); err != nil {
			log.Warning("accept failed:", err)
			continue
		} else {
			utils.CheckError(conn.SetReadBuffer(cfg.SockBuf))
			utils.CheckError(conn.SetWriteBuffer(cfg.SockBuf))
			go handleClient(conn, cfg)
		}
	}

}

func KcpServer(cfg *config.Config) {
	listener, err := kcp.Listen(cfg.Listen)
	utils.CheckError(err)
	log.Info("udp listening on:", listener.Addr())
	lis := listener.(*kcp.Listener)
	utils.CheckError(lis.SetReadBuffer(cfg.SockBuf))
	utils.CheckError(lis.SetWriteBuffer(cfg.SockBuf))
	utils.CheckError(lis.SetDSCP(cfg.Dscp))

	for {
		if conn, err := lis.AcceptKCP(); err != nil {
			log.Warning("accept failed:", err)
			continue
		} else {
			conn.SetWindowSize(cfg.Sndwnd, cfg.Rcvwnd)
			conn.SetNoDelay(cfg.NoDelay, cfg.Interval, cfg.Resend, cfg.NC)
			conn.SetStreamMode(true)
			conn.SetMtu(cfg.MTU)
			go handleClient(conn, cfg)
		}
	}
}

func handleClient(conn net.Conn, cfg *config.Config) {
	defer utils.Trace()
	defer conn.Close()
	in := make(chan []byte)
	defer func() {
		close(in)
	}()

	var s sess.Session
	host, port, err := net.SplitHostPort(conn.RemoteAddr().String())
	if err != nil {
		log.Error("cannot get remote address:", err)
		return
	}
	s.IP = net.ParseIP(host)
	log.Infof("new connection from:%v port:%v \n", host, port)

	s.Die = make(chan struct{})

	out := bf.NewBuffer(conn, s.Die, cfg.TxQueueLen)
	go out.Start()

	signal.WaitGroup.Add(1)
	go agent.Agent(&s, in, out)

	_head_bytes := make([]byte, 2)
	for {
		conn.SetReadDeadline(time.Now().Add(cfg.ReadDeadline))
		if n, err := io.ReadFull(conn, _head_bytes); err != nil {
			log.Warningf("read head failed:ip:%v reason:%v size:%v", s.IP, err, n)
			return
		} else {
			payload_bytes := make([]byte, binary.BigEndian.Uint16(_head_bytes))
			if n, err = io.ReadFull(conn, payload_bytes); err != nil {
				log.Warningf("read payload failed,ip:%v reason:%v size:%v", s.IP, err, n)
				return
			}
			select {
			case in <- payload_bytes:
			case <-s.Die:
				log.Warningf("connection closed by logic,flag:%v ip:%v", s.Flag, s.IP)
				return
			}
		}
	}
}
