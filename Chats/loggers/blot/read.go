package blot

import (
	"WorkSpace/GoDevWork/Chats/loggers/adapter"
	"github.com/HouzuoGuo/tiedot/db"
)

type TieReadFile struct {
	col *db.Col
}

func NewTieReadFile() (*TieReadFile, error) {
	if fd, err := db.OpenDB(getUserLogDB()); err != nil {
		return nil, err
	} else {
		return &TieReadFile{col: fd.Use(getUserTable())}, nil
	}
}

func (tf *TieReadFile) GetLogName() string {
	return "[tie-query]"
}

func (tf *TieReadFile) Init() error {
	return nil
}

func (tf *TieReadFile) Query(cond adapter.QueryCond) (interface{}, error) {

	return nil, nil
}
