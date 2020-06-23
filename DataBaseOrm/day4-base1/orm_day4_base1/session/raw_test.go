package session

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"orm_day4_base1/dialect"
	"os"
	"testing"
)

const (
	SQLite = "sqlite3"
)

var (
	testDb      *sql.DB
	testDial, _ = dialect.Get(SQLite)
)

func TestMain(m *testing.M) {
	testDb, _ = sql.Open(SQLite, "gee.db")
	code := m.Run()
	_ = testDb.Close()
	os.Exit(code)
}

func NewSession() *Session {
	return New(testDb, testDial)
}

func TestSession_Exec(t *testing.T) {
	sess := NewSession()
	sess.BuildSQL("DROP TABLE IF EXISTS User;").Exec()
	sess.BuildSQL("CREATE TABLE User(Name text);").Exec()
	result, _ := sess.BuildSQL("INSERT INTO User(`Name`) values(?),(?)", "Tom", "Sam").Exec()
	if count, err := result.RowsAffected(); err != nil {
		t.Fatal(fmt.Sprintf("count %d err:%v", count, err))
		return
	}
	t.Log("ok.")
}

func TestSession_QueryRows(t *testing.T) {
	sess := NewSession()
	sess.BuildSQL("DROP TABLE IF EXISTS User;").Exec()
	sess.BuildSQL("CREATE TABLE User(Name text);").Exec()
	row := sess.BuildSQL("SELECT count(*) FROM User").QueryRow()
	var count int
	if err := row.Scan(&count); err != nil || count != 0 {
		t.Fatal("failed to query db", err)
	}
}
