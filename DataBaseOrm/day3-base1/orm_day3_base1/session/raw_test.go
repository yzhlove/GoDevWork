package session

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"orm_day3_base1/dialect"
	"os"
	"testing"
)

var (
	TestDB      *sql.DB
	TestDial, _ = dialect.GetDialect("sqlite3")
)

func TestMain(m *testing.M) {
	TestDB, _ = sql.Open("sqlite3", "gee.db")
	code := m.Run()
	_ = TestDB.Close()
	os.Exit(code)
}

func NewSession() *Session {
	return New(TestDB, TestDial)
}
