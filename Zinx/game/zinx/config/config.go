package config

import "time"

//Etcd
var (
	Endpoints   = []string{"127.0.0.1:2379"}
	DialTimeout = 5 * time.Second
	Root        = "zinx-services"
	Monitors    = []string{"snowflake"}
	EventMax    = 128
)
