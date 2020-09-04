package config

import "time"

type Config struct {
	Listen                        string
	ReadDeadline                  time.Duration
	SockBuf                       int
	UDPSockBuf                    int
	TxQueueLen                    int
	Dscp                          int
	Sndwnd                        int
	Rcvwnd                        int
	MTU                           int
	NoDelay, Interval, Resend, NC int
	Timeout                       time.Duration
	EtcdHost                      []string
	EtcdRoot                      string
	Services                      []string
}
