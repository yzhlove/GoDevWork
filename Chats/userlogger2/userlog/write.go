package userlog

import (
	"fmt"
	"yo-star.com/nekopara/manager/logger/userlogger/base"
	"yo-star.com/nekopara/manager/logger/userlogger/tiedb"
)

type LogWriterManager struct {
	ChanMessage   chan base.LogMessage
	LogWriterList []base.LogWriter
}

func newLogWriterManager() *LogWriterManager {
	return &LogWriterManager{
		ChanMessage: make(chan base.LogMessage, 128),
	}
}

func (u *LogWriterManager) start() {
	go func() {
		for msg := range u.ChanMessage {
			for _, apt := range u.LogWriterList {
				if err := apt.Write(msg); err != nil {
					fmt.Println(apt.GetLogName(), " write err:", err)
				}
			}
		}
	}()
}

//写入日志
func WriteLog(msg base.LogMessage) {
	if userLogWriterManager != nil {
		if len(msg) > 0 {
			userLogWriterManager.ChanMessage <- msg
		}
	}
}

var userLogWriterManager *LogWriterManager

func registerWriter() error {
	//[console] 测试
	//userLogWriterManager.LogWriterList = []base.LogWriter{console.NewConsoleWrite()}

	if tw, err := tiedb.NewTieWriter(); err != nil {
		return err
	} else {
		userLogWriterManager.LogWriterList = append(userLogWriterManager.LogWriterList, tw)
	}
	return nil
}

//初始化
func LogWriterInit() error {

	userLogWriterManager = newLogWriterManager()
	//注册writer
	if err := registerWriter(); err != nil {
		return err
	}
	//初始化writer
	for _, apt := range userLogWriterManager.LogWriterList {
		if err := apt.Init(); err != nil {
			fmt.Println(apt.GetLogName(), " Init err:", err)
			return err
		}
	}
	//start
	userLogWriterManager.start()
	return nil
}
