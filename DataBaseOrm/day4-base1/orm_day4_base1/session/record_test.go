package session

import "testing"

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

var (
	user1 = User{"Tom", 18}
	user2 = User{"Sam", 25}
	user3 = User{"Jack", 25}
)

func testRecordInit(t *testing.T) *Session {
	t.Helper()
	sess := NewSession().Model(&User{})
	if err := sess.DropTb(); err != nil {
		t.Fatal(err)
	}
	if err := sess.CreateTb(); err != nil {
		t.Fatal(err)
	}
	if _, err := sess.Insert(user1, user2); err != nil {
		t.Fatal(err)
	}
	return sess
}

func TestSession_Insert(t *testing.T) {
	sess := testRecordInit(t)
	affected, err := sess.Insert(user3)
	if err != nil || affected != 1 {
		t.Fatal(err)
	}
	t.Log("insert ok")
}

func TestSession_Find(t *testing.T) {
	sess := testRecordInit(t)
	var users []User
	if err := sess.Find(&users); err != nil || len(users) != 2 {
		t.Fatal(err)
	}
	t.Log("find ok")
}

func TestSession_First(t *testing.T) {
	sess := testRecordInit(t)
	u := &User{}
	if err := sess.First(u); err != nil {
		t.Fatal(err)
	}
	t.Log("first user:", u)
}

func TestSession_Limit(t *testing.T) {
	sess := testRecordInit(t)
	var users []User
	if err := sess.Limit(1).Find(&users);err != nil {
		t.Fatal(err)
	}
	t.Log("users:",users)
}

func TestSession_Where(t *testing.T) {
	//sess := testRecordInit(t)
	//var users []User

}
