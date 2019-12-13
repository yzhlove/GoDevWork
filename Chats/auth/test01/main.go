package main

//acl
import (
	"fmt"
	"github.com/casbin/casbin/v2"
)

func main() {

	e, err := casbin.NewEnforcer("./Chats/auth/mode.conf", "./Chats/auth/policy.csv")
	if err != nil {
		panic(err)
	}
	ok, err := e.Enforce("super", "b", "c")
	if ok {
		fmt.Println("ok")
	} else {
		fmt.Println("not ok")
	}

	fmt.Println(e.GetAllActions())
	fmt.Println(e.GetAllObjects())
	fmt.Println(e.GetAllSubjects())
	fmt.Println(e.GetAllNamedActions("p"))
	fmt.Println(e.GetAllNamedObjects("p"))
	fmt.Println(e.GetAllNamedSubjects("p"))

	ok, err = e.AddPolicy("love", "manager-default", "default")
	if err != nil {
		panic(err)
	}

	ok, err = e.AddNamedPolicy("p", "like", "manager-default", "yes")

	fmt.Println("ok =======> ", ok)

	fmt.Println(e.GetAllActions())
	fmt.Println(e.GetAllObjects())
	fmt.Println(e.GetAllSubjects())

	e.GetModel().AddDef("m", "match", "r.user == admin")

	if err = e.SavePolicy(); err != nil {
		panic(err)
	}

}
