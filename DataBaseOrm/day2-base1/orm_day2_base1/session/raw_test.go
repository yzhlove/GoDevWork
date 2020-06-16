package session

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"orm_day2_base1/dialect"
	"os"
	"testing"
)

var (
	TestDB      *sql.DB
	TestDial, _ = dialect.GetDialect("sqlite3")
)

func NewSession() *Session {
	return New(TestDB, TestDial)
}

func TestMain(m *testing.M) {
	TestDB, _ = sql.Open("sqlite3", "./gee.db")
	code := m.Run()
	_ = TestDB.Close()
	os.Exit(code)
}

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

func Test_Table(t *testing.T) {
	sess := NewSession().Model(&User{})
	_ = sess.DropTable()
	_ = sess.CreateTable()
	if !sess.HasTable() {
		t.Error("failed to create table user.")
	}
	t.Log("ok.")
}
