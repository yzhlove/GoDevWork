package tiedb

import (
	"WorkSpace/GoDevWork/Chats/userlogger2/base"
	"encoding/json"
	"errors"
	"github.com/HouzuoGuo/tiedot/db"

)

type TieReader struct {
	Col *db.Col
	qs  map[string]LogQuery
}

func NewTieReader() (*TieReader, error) {
	if fd, err := db.OpenDB(getDataBase()); err != nil {
		return nil, err
	} else {
		return &TieReader{Col: fd.Use(getTable())}, nil
	}
}

func (t *TieReader) GetLogName() string {
	return "[tie-db]"
}

func (t *TieReader) register() {
	for _, q := range List {
		t.qs[q.GetQueryName()] = q
	}
}

func (t *TieReader) Init() error {
	if t.qs == nil {
		t.qs = make(map[string]LogQuery, 4)
	}
	t.register()
	return nil
}

func (t *TieReader) Read(cond base.LogCondMessage) (interface{}, error) {
	var opt struct{ E string `json:"event"` } //用户行为
	if err := json.Unmarshal([]byte(cond), &opt); err != nil {
		return nil, err
	}
	if opt.E == "" {
		opt.E = "complex"
	}
	if inter, ok := t.qs[opt.E]; ok {
		return inter.Query(t, cond)
	}
	return nil, errors.New("not found event:" + opt.E)
}
