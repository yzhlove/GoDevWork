package main

import "github.com/casbin/casbin/v2"

func main() {

	e, err := casbinInit()
	if err != nil {
		panic("casbin init err:" + err.Error())
	}

	if _, err = e.AddPolicy("yzh", "manager-list", "btn-list"); err != nil {
		panic("add policy err:" + err.Error())
	}

	if err = e.SavePolicy(); err != nil {
		panic("save err:" + err.Error())
	}

}

func casbinInit() (*casbin.Enforcer, error) {

	m, err := getMode()
	if err != nil {
		panic("mode err:" + err.Error())
	}

	apt, err := NewAdapter("Chats/auth/test06/config.ini")
	if err != nil {
		panic("adapter err:" + err.Error())
	}

	return casbin.NewEnforcer(m, apt)
}
