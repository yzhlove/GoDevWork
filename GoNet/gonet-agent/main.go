package main

import (
	"flag"
	"github.com/sirupsen/logrus"
	"micro_agent/config"
	"micro_agent/signal"
	"net/http"
	"strings"
	"time"
)

var cfg *config.Config

func init() {
	listen := flag.String("listen", ":8888", "listen address:port")
	etcdHost := flag.String("etcd-host", "http://localhost:2379", "etcd host")
	etcdRoot := flag.String("etcd-root", "/backends", "etcd root path")
	services := flag.String("services", "snowfalke-10000,game-10000", "auto-discover service")
	readDeadline := flag.Int64("read-deadline", 15, "per connection read timeout")
	txQueuelen := flag.Int("txqueuelen", 128, "per connection output message queue")
	sockBuf := flag.Int("sockbuf", 32767, "per connection tcp socket buffer")
	udpSockBuf := flag.Int("udp-sockbuf", 4194304, "per connection udp send window")
	udpSndwnd := flag.Int("udp-sndwnd", 32, "per connection send window")
	udpRcvwnd := flag.Int("upd-rcvwnd", 32, "per connection  udp recv window")
	udpMtu := flag.Int("udp-mtu", 1280, "MTU of udp packets,without IP(20) + UDP(80)")
	dscp := flag.Int("dscp", 46, "set DSCP(6bit)")
	nodelay := flag.Int("nodelay", 1, "ikcp_nodelay")
	interval := flag.Int("interval", 20, "ikcp_nodelay")
	resend := flag.Int("resend", 1, "ikcp_nodelay")
	nc := flag.Int("nc", 1, "ikcp_nodelay")
	rpm := flag.Int("rpm", 200, "per connection rpc limit")
	cfg = &config.Config{
		Listen:       *listen,
		EtcdRoot:     *etcdRoot,
		EtcdHost:     []string{*etcdHost},
		Services:     strings.Split(*services, ","),
		ReadDeadline: time.Duration(*readDeadline) * time.Second,
		TxQueueLen:   *txQueuelen,
		SockBuf:      *sockBuf,
		UDPSockBuf:   *udpSockBuf,
		Sndwnd:       *udpSndwnd,
		Rcvwnd:       *udpRcvwnd,
		MTU:          *udpMtu,
		Dscp:         *dscp,
		NoDelay:      *nodelay,
		Interval:     *interval,
		Resend:       *resend,
		NC:           *nc,
		RPM:          *rpm,
	}
}

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	go http.ListenAndServe("0.0.0.0:6060", nil)

}

func startUp(cfg *config.Config) {
	go signal.SigHandler()

}
