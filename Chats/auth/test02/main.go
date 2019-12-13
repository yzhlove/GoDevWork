package main

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
)

//自定义mode文件

func main() {

	e, err := casbin.NewEnforcer("./Chats/auth/tmode.conf", "./Chats/auth/tpolicy.csv")
	if err != nil {
		panic(err)
	}

	m := model.NewModel()
	m.AddDef("r", "request", "sub, obj, act")
	m.AddDef("p", "policy", "sub, obj, act")
	m.AddDef("e", "effect", "some(where (p.eft == allow))")
	m.AddDef("m", "match", "m = r.sub == p.sub && r.obj == p.obj && r.act == p.act")

	//e.SetModel(m)

	e.AddNamedPolicy("p", "alice", "book", "read")
	e.AddNamedPolicy("p", "alice", "book", "write")
	e.AddNamedPolicy("p", "alice", "book", "yes")

	//if ok, err := e.Enforce("alice", "data1", "read"); err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println("====> ", ok)
	//}

	if err = e.SavePolicy(); err != nil {
		panic(err)
	}

}
