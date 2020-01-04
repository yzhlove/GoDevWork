package main

import (
	"WorkSpace/GoDevWork/Chats/auth/test11/config"
	"WorkSpace/GoDevWork/Chats/auth/test11/rules"
	"WorkSpace/GoDevWork/Chats/auth/test11/storage"
	"testing"
)

func Test_Casbin(t *testing.T) {

	config.ACLFilePath = "test.ini"

	s, err := storage.NewStorage()
	if err != nil {
		t.Error(err)
		return
	}

	e, err := rules.NewEnforcer(s)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("isSuper ", e.IsSuper("liuxiaoyu"))
	t.Log("isSuper ", e.IsSuper("guoxh"))
	t.Log("isSuper ", e.IsSuper("super"))

	err = e.SetAuths("nokia", []string{"aaa", "bbb", "ccc", "ddd", "eee", "fff"})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("auths ", e.GetAuths("super"))
	t.Log("auths ", e.GetAuths("guoxh"))
	t.Log("auths ", e.GetAuths("liuxiaoyu"))
	//

	t.Log("===================================")

	t.Log(e.Check("super", "abc"))
	t.Log(e.Check("guoxh", "abc"))
	t.Log(e.Check("liuxiaoyu", "ban"))
	t.Log(e.Check("liuxiaoyu", "abc"))
	t.Log(e.Check("wuyifan", "abcderf"))
	t.Log(e.Check("wuyifan", "ban"))

}
