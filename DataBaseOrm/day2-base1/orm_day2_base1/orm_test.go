package orm_day2_base1

import (
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func OpenDB(t *testing.T) *Engine {
	t.Helper()
	engine, err := NewEngine("sqlite3", "./session/gee.db")
	if err != nil {
		t.Error(err)
		return nil
	}
	return engine
}

func Test_NenEngine(t *testing.T) {
	engine := OpenDB(t)
	if engine != nil {
		defer engine.Close()
		t.Log("ok")
	} else {
		t.Error("engine is nil.")
	}
}
