package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"zinx/ziface"
)

type Conf struct {
	TcpServer         ziface.ServerImp
	ServerName        string
	Host              string
	TcpPort           int
	Version           string
	MaxPackageSize    uint32
	MaxConnections    uint32
	WorkerPoolSize    uint32
	MaxWorkerTaskSize uint32
	MaxMsgChanSize    uint32
}

var (
	GlobalConfig *Conf
	filePath     string
)

func init() {
	flag.StringVar(&filePath, "f", "config.json", "config file path")
	flag.Parse()

	if err := GlobalConfig.Reload(); err != nil {
		log.Println("config reload err:", err)
		GlobalConfig = DefaultConf()
	}

}

func DefaultConf() *Conf {
	return &Conf{
		ServerName:        "zinxTcpServer",
		Version:           "1.0",
		TcpPort:           7777,
		Host:              "0.0.0.0",
		MaxConnections:    1024,
		MaxPackageSize:    4096,
		WorkerPoolSize:    16,
		MaxWorkerTaskSize: 4,
	}
}

func (c *Conf) Reload() error {
	if res, err := ioutil.ReadFile(filePath); err != nil {
		return err
	} else {
		return json.Unmarshal(res, &GlobalConfig)
	}
}
