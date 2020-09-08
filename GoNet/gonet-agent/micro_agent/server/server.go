package server

import (
	log "github.com/sirupsen/logrus"
	"micro_agent/bf"
	"micro_agent/config"
	"micro_agent/sess"
	"micro_agent/signal"
	"micro_agent/utils"
	"net"
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

		}
	}

}

func KcpServer(cfg *config.Config) {

}

func handleClient(conn net.Conn, cfg *config.Config) {
	defer utils.Trace()
	defer conn.Close()
	head := make([]byte, 2)
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

}
