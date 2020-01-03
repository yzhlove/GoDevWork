package storage

import (
	"WorkSpace/GoDevWork/Chats/auth/test11/config"
	"fmt"
	"testing"
)

func TestStorage_LoadAuth(t *testing.T) {

	auths := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}

	for i, v := range auths[3:] {
		fmt.Println(i, " - ", v)
	}

}

func TestStorage_Option(t *testing.T) {

	config.ACLFilePath = "test.ini"

	s, err := NewStorage()
	if err != nil {
		t.Error(err)
		return
	}

	for user, auth := range s.LoadAuth() {
		t.Log(user, " - ", auth)
	}

	auths := []string{"A", "B", "C", "D", "E"}
	_ = s.Submit("kfc", "123456789")
	_ = s.SaveAuth("kfc", auths)

	t.Log("kfc exists => ", s.Exists("kfc"))
	t.Log("kfd exists => ", s.Exists("kfd"))

	t.Log(s.GetUserList())

	t.Log(s.GetPasswd("kfc"))

	//_ = s.Delete("kfc")

}
