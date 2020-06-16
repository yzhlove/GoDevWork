package schema

import (
	"orm_day2_base1/dialect"
	"testing"
)

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

func Test_Parse(t *testing.T) {

	testDial, _ := dialect.GetDialect("sqlite3")
	schema := Parse(&User{}, testDial)
	if schema.Name != "User" || len(schema.Fields) != 2 {
		t.Fatal("failed to parse User struct")
	}
	if schema.GetField("Name").Tag != "PRIMARY KEY" {
		t.Fatal("failed to parse primary key")
	}
	t.Log("ok.")
}
