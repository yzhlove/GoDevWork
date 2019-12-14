package router

import (
	"WorkSpace/GoDevWork/Chats/auth/test04_auth/conf"
	"context"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
)

type CheckRouter struct {
	Router  *httprouter.Router
	adapter *casbin.Enforcer
}

func NewCheckRouter(apt *casbin.Enforcer) *CheckRouter {
	return &CheckRouter{
		Router:  httprouter.New(),
		adapter: apt,
	}
}

func Err(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusForbidden)
	_, _ = w.Write([]byte(msg))
}

func Crash(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte(msg))
}

func Succeed(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(msg))
}

func getResource(url string) string {
	s := strings.Split(url, "/")
	if len(s) < 3 {
		return ""
	}
	return s[1] + "-" + s[2]
}

func (c *CheckRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	fmt.Println("params ==> ", params)
	var username string
	if u, ok := conf.UserList[params.Get("token")]; !ok {
		Err(w, "not found user")
		return
	} else {
		username = u
	}
	target := params.Get("target")
	if username != "super" && target == "" {
		Err(w, "target is null")
		return
	}
	fmt.Printf("check:%s %s %s \n", username, getResource(r.URL.Path), target)
	if ok, err := c.adapter.Enforce(username, getResource(r.URL.Path), target); err != nil {
		Err(w, err.Error())
		return
	} else if !ok {
		Err(w, "check err")
		return
	}
	ctx := context.WithValue(r.Context(), "username", username)
	c.Router.ServeHTTP(w, r.WithContext(ctx))
}

func (c *CheckRouter) AddRules(w http.ResponseWriter, r *http.Request, t httprouter.Params) {
	fmt.Println(r.Context().Value("username"), " add rules ...")
	user := t.ByName("user")
	resource := t.ByName("resource")
	auth := t.ByName("auth")
	fmt.Println("| user => ", user, " resource => ", resource, " auth => ", auth)

	if _, err := c.adapter.AddPolicy(user, resource, auth); err != nil {
		Crash(w, "add policy => "+err.Error())
		return
	}
	if err := c.adapter.SavePolicy(); err != nil {
		Crash(w, "save policy => "+err.Error())
		return
	}
	//if err := c.adapter.LoadPolicy(); err != nil {
	//	Crash(w, "load policy => "+err.Error())
	//	return
	//}
	Succeed(w, "ok")
}

func (c *CheckRouter) DelRules(w http.ResponseWriter, r *http.Request, t httprouter.Params) {
	fmt.Println("delete rules")
}
