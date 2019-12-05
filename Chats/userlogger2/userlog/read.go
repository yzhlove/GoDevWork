package userlog

import (
	"yo-star.com/nekopara/manager/logger/userlogger/base"
	"yo-star.com/nekopara/manager/logger/userlogger/tiedb"
)

type LogReaderManager struct {
	LogReader base.LogReader
}

var userLogReaderManager *LogReaderManager

func newLogReaderManager() *LogReaderManager {
	return &LogReaderManager{}
}

func registerReader() error {
	if tr, err := tiedb.NewTieReader(); err != nil {
		return err
	} else {
		userLogReaderManager.LogReader = tr
	}
	return nil
}

func LogReaderInit() error {
	userLogReaderManager = newLogReaderManager()
	//注册
	if err := registerReader(); err != nil {
		return err
	}
	//初始化
	if err := userLogReaderManager.LogReader.Init(); err != nil {
		return err
	}

	return nil
}

//查询
func QueryResult(cond string) (result interface{}, err error) {
	if userLogReaderManager != nil {
		result, err = userLogReaderManager.LogReader.Read(base.LogCondMessage(cond))
	}
	return
}
