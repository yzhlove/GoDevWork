package config

import "time"

type Config struct {
	GameId      string
	Listen      string
	EtcdHosts   []string
	EtcdRoot    string
	Services    []string
	PathNumbers string
	Streams     uint32
	Timeout     time.Duration
}
