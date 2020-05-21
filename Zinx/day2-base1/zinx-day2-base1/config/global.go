package config

import (
	"encoding/json"
	"io/ioutil"
	"zinx-day2-base1/ziface"
)

type GlobalConf struct {
	TcpServer     ziface.ServerInterface //zinx的全局server对象
	Host          string                 //主机IP
	TcpPort       int                    //端口
	Name          string                 //服务器名称
	Version       string                 //服务器版本
	MaxPacketSize uint32                 //数据包的最大长度
	MaxConn       int                    //当前允许的最大连接数
}

// GlobalConfig 全局配置
var GlobalConfig *GlobalConf

func init() {
	//default params
	GlobalConfig := &GlobalConf{
		Name:          "ZinxServerApp",
		Version:       "1.0",
		TcpPort:       7777,
		Host:          "0.0.0.0",
		MaxConn:       12000,
		MaxPacketSize: 4096,
	}
	GlobalConfig.Reload()
}

func (g *GlobalConf) Reload() {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic("[Conf] read file err " + err.Error())
	}
	if err = json.Unmarshal(data, &GlobalConfig); err != nil {
		panic("[Conf] conf err " + err.Error())
	}
}
