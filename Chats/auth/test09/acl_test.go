package main

import (
	"WorkSpace/GoDevWork/Chats/auth/test09/apt"
	"WorkSpace/GoDevWork/Chats/auth/test09/conf"
	"WorkSpace/GoDevWork/Chats/auth/test09/storage"
	"testing"
)

func Test_Acl(t *testing.T) {

	conf.StorageFile = "TestStorage/auth.ini"

	s, err := storage.NewStorage()
	if err != nil {
		t.Error(err)
		return
	}

	adapter, err := apt.NewEnforcerContext(s)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(adapter.Enforcer("yzh", "abc"))
	t.Log(adapter.Enforcer("yzh", "login"))
	t.Log(adapter.Enforcer("wuyifan", "adc"))

	//_, _ = adapter.DelAuth("yzh")
	//adapter.Delete("yzh", "login")
	//adapter.DelAuth("yzh")
	//err = adapter.SetAuth("yzh", []string{"hello", "world", "what", "are", "you", "doing"})
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	/*
			=== RUN   Test_Acl
			line ===>  p, yzh, super
			tokens ==>  [p  yzh  super]
			tokens ==>  [p yzh super]
			 key ==>  p  sec ==>  p
			line ===>  p, yzh, login
			tokens ==>  [p  yzh  login]
			tokens ==>  [p yzh login]
			 key ==>  p  sec ==>  p
			line ===>  p, yzh, manager
			tokens ==>  [p  yzh  manager]
			tokens ==>  [p yzh manager]
			 key ==>  p  sec ==>  p
			sec =>  p  ptype =>  p  rules =>  [super super]
			user ==  super  auth ===  [super]
			===============================  222
		sec =>  p  ptype =>  p  rules =>  [yzh hello world what are you doing]
		user ==  yzh  auth ===  [hello world what are you doing]
			sec ==>  p
			ptype ==>  p
			fieldIndex ==>  0
			fieldValues ==>  [yzh]
	*/

}
