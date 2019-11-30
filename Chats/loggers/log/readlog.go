package log

import (
	"WorkSpace/GoDevWork/Chats/loggers/adapter"
	"WorkSpace/GoDevWork/Chats/loggers/blot"
	"encoding/json"
)

type QueryManager struct {
	q adapter.LogQuery
}

func NewQueryManager(q adapter.LogQuery) *QueryManager {
	return &QueryManager{q: q}
}

var queryManager *QueryManager

func UserQueryInit() error {
	if q, err := blot.NewTieReadFile(); err != nil {
		return err
	} else {
		queryManager = NewQueryManager(q)
	}
	return nil
}

func QueryResult(cond adapter.QueryCond) (data []byte, err error) {
	if queryManager != nil {
		var result interface{}
		if result, err = queryManager.q.Query(cond); err != nil {
			return
		}
		if result != nil {
			data, err = json.Marshal(result)
		}
	}
	return
}

func QueryByTs(cond adapter.QueryCond) (data []byte, err error) {
	return
}

func QueryByOperator(cond adapter.QueryCond) (data []byte, err error) {
	return
}
