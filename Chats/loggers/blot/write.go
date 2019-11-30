package blot

import (
	"WorkSpace/GoDevWork/Chats/loggers/adapter"
	"WorkSpace/GoDevWork/Chats/loggers/config"
	"fmt"
	"github.com/HouzuoGuo/tiedot/db"
)

//索引列表
var ins = []string{"ts", "operator"}

type TieWriteFile struct {
	fd *db.DB
}

func getUserLogDB() string {
	return config.UserLogPath + "TestUserLogDB"
}

func getUserTable() string {
	return "UserLog"
}

func NewTieWriteFile() (*TieWriteFile, error) {
	if fd, err := db.OpenDB(getUserLogDB()); err != nil {
		return nil, err
	} else {
		return &TieWriteFile{fd: fd}, nil
	}
}

func (f *TieWriteFile) Init() (err error) {
	fmt.Println(f.GetLogName(), " Init ")
	//创建table
	if tb := getUserTable(); !f.fd.ColExists(tb) {
		if err = f.fd.Create(tb); err != nil {
			return
		}
		//创建索引
		err = f.fd.Use(tb).Index(ins)
	}
	return
}

func (f *TieWriteFile) GetLogName() string {
	return "[tie-logger]"
}

func (f *TieWriteFile) Write(message adapter.LogMessage) (err error) {
	var logger *db.Col
	if f.fd != nil {
		logger = f.fd.Use(getUserTable())

		doc := make(map[string]interface{}, len(message)/2)
		for i, msg := range message {
			if i%2 == 0 {
				if key, ok := msg.(string); ok {
					if i+1 < len(message) {
						doc[key] = message[i+1]
					}
				}
			}
		}
		//写入日志
		_, err = logger.Insert(doc)
	}

	return
}
