package main

import (
	"WorkSpace/GoDevWork/Chats/auth/test04_auth/acl"
	"WorkSpace/GoDevWork/Chats/auth/test04_auth/router"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("login succeed")
}

func main() {

	adapter, err := acl.NewCheckAdpater2()
	if err != nil {
		panic(err)
	}
	checkRouter := router.NewCheckRouter(adapter)
	//resource list
	checkRouter.Router.GET("/manager/auth/add/:user/:resource/:auth", checkRouter.AddRules)
	checkRouter.Router.DELETE("/manager/auth/delete", checkRouter.DelRules)
	checkRouter.Router.GET("/manager/auth/login", login)

	//start server
	fmt.Println("start server...")
	if err = http.ListenAndServe(":1234", checkRouter); err != nil {
		panic(err)
	}

}
