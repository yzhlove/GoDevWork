package config

import (
	"encoding/json"
	"io/ioutil"
	"zinx-day4-base1/ziface"
)

type GlobalConf struct {
	TcpServer   ziface.MessageInterface //zinx server
	Host        string
	TcpPort     int
	Name        string
	Version     string
	MaxPackSize uint32
	MaxConn     int
}

var GlobalConfig *GlobalConf

func init() {
	GlobalConfig = &GlobalConf{
		Name:        "ZinxServerApp",
		Version:     "1.0",
		TcpPort:     7777,
		Host:        "0.0.0.0",
		MaxConn:     12000,
		MaxPackSize: 4096,
	}
	//GlobalConfig.Reload()
}

func (g *GlobalConf) Reload() {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic("[conf]" + err.Error())
	}
	if err := json.Unmarshal(data, &GlobalConfig); err != nil {
		panic("[conf]" + err.Error())
	}
}
