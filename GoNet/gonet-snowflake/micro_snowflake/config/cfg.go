package config

import (
	"strconv"
	"time"
)

type Config struct {
	Host      string        //grpc地址
	EtcdHost  []string      //etcd地址
	MachineID int           //机器ID
	Root      string        //etcd服务发现目录
	Prefix    string        //etcd服务发现文件夹
	TimeOut   time.Duration //etcd 连接超时
}

// New 创建一个新的Config
func New(host, id, root, prefix string, etcdHost []string, timeout time.Duration) *Config {
	sid, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	return &Config{
		Host:      host,
		EtcdHost:  etcdHost,
		MachineID: sid,
		Root:      root,
		TimeOut:   timeout,
		Prefix:    prefix,
	}
}
