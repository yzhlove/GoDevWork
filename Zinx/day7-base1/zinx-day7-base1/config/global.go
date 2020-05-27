package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"zinx-day7-base1/ziface"
)

type Conf struct {
	TcpServer         ziface.ServerInterface
	ServerName        string
	Host              string
	TcpPort           int
	Version           string
	MaxPackSize       uint32
	MaxConn           int
	WorkerPoolSize    uint32
	MaxWorkerTaskSize uint32
	MaxMsgChanSize    uint32
}

var GlobalConfig *Conf

func init() {

	if err := GlobalConfig.Reload(); err != nil {
		log.Println("loading config file err:", err)
		GlobalConfig = LoadDefaultConf()
	}
}

func LoadDefaultConf() *Conf {
	return &Conf{
		ServerName:        "ZinxServerApp",
		Version:           "1.0",
		TcpPort:           7777,
		Host:              "0.0.0.0",
		MaxConn:           6,
		MaxPackSize:       4096,
		WorkerPoolSize:    4,
		MaxWorkerTaskSize: 10,
	}
}

func (g *Conf) Reload() error {
	if data, err := ioutil.ReadFile("config.json"); err != nil {
		return err
	} else {
		return json.Unmarshal(data, &GlobalConfig)
	}
}
