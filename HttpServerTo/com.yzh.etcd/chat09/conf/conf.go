package conf

import "time"

var (
	Endpoints   = []string{"127.0.0.1:2379"}
	DialTimeout = 5 * time.Second
	Root        = "discovery"
	MonitorServ = []string{"snowflak", "nats"}
	EventMax    = 128
)
