package main

import (
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var userMap = map[string]string{"token": "yzh"}

func index(w http.ResponseWriter, r *http.Request, t httprouter.Params) {
	fmt.Println("====> 12345")
}

func login(w http.ResponseWriter, r *http.Request, t httprouter.Params) {

	fmt.Println("login ===> ")
	fmt.Println("login name => ", r.Context().Value("username"))

}

type MyRouter struct {
	router *httprouter.Router
}

func NewRouter() *MyRouter {
	return &MyRouter{router: httprouter.New()}
}

func (m *MyRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("url => ", r.URL)
	_ = r.ParseForm()
	var username string
	if user, ok := userMap[r.FormValue("token")]; !ok {
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte("not found user"))
		return
	} else {
		fmt.Println("user => ", user)
		username = user
	}
	ctx := context.WithValue(r.Context(), "username", username)
	m.router.ServeHTTP(w, r.WithContext(ctx))
}

func main() {
	my := NewRouter()
	my.router.GET("/test/index", index)
	my.router.POST("/test/login", login)
	if err := http.ListenAndServe(":1234", my); err != nil {
		panic(err)
	}
}
