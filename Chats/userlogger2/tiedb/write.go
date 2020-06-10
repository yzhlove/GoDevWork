package tiedb

import (
	"WorkSpace/GoDevWork/Chats/userlogger2/base"
	"WorkSpace/GoDevWork/Chats/userlogger2/config"
	"fmt"
	"github.com/HouzuoGuo/tiedot/db"
	"os"
)

//默认索引
var ins = [][]string{{"operator"}, {"day"}, {"event"}, {"year"}}

type TieWriter struct {
	tw *db.DB
}

func getDataBase() string {
	return config.UserLoggerPath + "LogUserDataBase"
}

func getTable() string {
	return "UserLog"
}

func NewTieWriter() (*TieWriter, error) {
	//创建目录
	if err := os.MkdirAll(config.UserLoggerPath, 0755); err != nil {
		return nil, fmt.Errorf("create user zlog path err:%v", err)
	}
	if tw, err := db.OpenDB(getDataBase()); err != nil {
		return nil, err
	} else {
		return &TieWriter{tw: tw}, nil
	}
}

func (t *TieWriter) GetLogName() string {
	return "[tie-db]"
}

func (t *TieWriter) createIndex() error {
	col := t.tw.Use(getTable())
	for _, index := range ins {
		if err := col.Index(index); err != nil {
			return err
		}
	}
	return nil
}

func (t *TieWriter) Init() (err error) {
	//如果table不存在
	if tb := getTable(); !t.tw.ColExists(tb) {
		if err = t.tw.Create(tb); err != nil {
			return
		}
		//创建索引
		err = t.createIndex()
	}
	return
}

func (t *TieWriter) Write(message base.LogMessage) (err error) {

	log := t.tw.Use(getTable())
	doc := make(map[string]interface{}, len(message)>>1)
	for i, msg := range message {
		if i%2 == 0 {
			if event, ok := msg.(string); ok && i+1 < len(message) {
				doc[event] = message[i+1]
			}
		}
	}
	//写入日志到tie-db
	_, err = log.Insert(doc)
	return
}
