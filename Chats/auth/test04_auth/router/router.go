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

func getResource(url string) string {
	s := strings.Split(url, "/")
	if len(s) < 3 {
		return ""
	}
	return s[1] + s[2]
}

func (c *CheckRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	var username string
	if u, ok := conf.UserList[r.FormValue("token")]; !ok {
		Err(w, "not found user")
		return
	} else {
		username = u
	}
	target := r.FormValue("target")
	if username != "super" && target == "" {
		Err(w, "target is null")
		return
	}
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
	fmt.Println(r.Context().Value("username")," add rules ...")
	
}

func (c *CheckRouter) DelRules(w http.ResponseWriter, r *http.Request, t httprouter.Params) {
	fmt.Println("delete rules")
}
