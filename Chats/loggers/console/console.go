package console

import (
	"WorkSpace/GoDevWork/Chats/loggers/adapter"
	"fmt"
)

type Log struct{}

func NewConsoleLog() *Log {
	return &Log{}
}

func (Log) GetLogName() string {
	return "[console-logger]"
}

func (log Log) Write(message adapter.LogMessage) (err error) {
	fmt.Println(log.GetLogName(), " ", message)
	return
}

func (log Log) Init() (err error) {
	fmt.Println(log.GetLogName(), " Init ")
	return
}
