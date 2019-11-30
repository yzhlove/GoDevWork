package log

import (
	"WorkSpace/GoDevWork/Chats/loggers/adapter"
	"WorkSpace/GoDevWork/Chats/loggers/blot"
	"WorkSpace/GoDevWork/Chats/loggers/console"
	"fmt"
)

type UserLogManager struct {
	ChanLogMsg chan adapter.LogMessage
	LogWriter  []adapter.LogWriter
}

func NewUserLogManager() *UserLogManager {
	return &UserLogManager{
		ChanLogMsg: make(chan adapter.LogMessage, 128),
		LogWriter:  adapter.GetAdapter(),
	}
}

func (us *UserLogManager) start() {
	go func() {
		for {
			select {
			case msg, ok := <-us.ChanLogMsg:
				if ok {
					for _, apt := range us.LogWriter {
						if err := apt.Write(msg); err != nil {
							fmt.Println(apt.GetLogName(), " write err ", err)
						}
					}
				}
			}
		}
	}()
}

func WriteLog(msg ...interface{}) {
	if userLogManager != nil {
		userLogManager.ChanLogMsg <- msg
	}
}

var userLogManager *UserLogManager

func registerAdapter() error {
	bt, err := blot.NewTieWriteFile()
	if err != nil {
		return err
	}
	cs := console.NewConsoleLog()
	adapter.Register(bt, cs)
	return nil
}

func UserLogInit() error {
	if err := registerAdapter(); err != nil {
		return err
	}
	userLogManager = NewUserLogManager()
	//初始化所有的适配器
	for _, apt := range userLogManager.LogWriter {
		if err := apt.Init(); err != nil {
			fmt.Println(apt.GetLogName(), " init err ", err)
			return err
		}
	}
	//开启协程
	userLogManager.start()
	return nil
}
