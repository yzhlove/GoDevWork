package console

import (
	"fmt"
	"yo-star.com/nekopara/manager/logger/userlogger/base"
)

//测试

type ConsoleWrite struct{}

func NewConsoleWrite() *ConsoleWrite {
	return &ConsoleWrite{}
}

func (ConsoleWrite) GetLogName() string {
	return "[console]"
}

func (c ConsoleWrite) Write(message base.LogMessage) (err error) {
	fmt.Println(c.GetLogName(), " ", message)
	return
}

func (c ConsoleWrite) Init() (err error) {
	fmt.Println(c.GetLogName(), " Init ok.")
	return
}
