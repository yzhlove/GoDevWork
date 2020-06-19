package session

import "testing"

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

var (
	user1 = User{"tom", 18}
	user2 = User{"sam", 25}
	user3 = User{"jack", 25}
)

func testRecordInit(t *testing.T) *Session {
	t.Helper()
	s := NewSession().Model(User{})
	if err := s.DropTable(); err != nil {
		t.Fatal(err)
	}
	if err := s.CreateTable(); err != nil {
		t.Fatal(err)
	}
	if _, err := s.Insert(user1, user2); err != nil {
		t.Fatal(err)
	}
	return s
}

func TestSession_Insert(t *testing.T) {
	s := testRecordInit(t)
	affected, err := s.Insert(user3)
	if err != nil || affected != 1 {
		t.Fatal("failed to create record")
	}
	t.Log("ok.")
}

func TestSession_Find(t *testing.T) {
	s := testRecordInit(t)
	var users []User
	if err := s.Find(&users); err != nil {
		t.Fatal(err)
	}
	for _, user := range users {
		t.Log(user)
	}
	t.Log("ok.")
}
