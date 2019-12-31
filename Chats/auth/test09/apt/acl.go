package apt

import (
	"WorkSpace/GoDevWork/Chats/auth/test09/conf"
	"WorkSpace/GoDevWork/Chats/auth/test09/storage"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
)

// "auth = super" 设置权限为超管
const rule = `
[request_definition]
r = user, auth
[policy_definition]
p = user, auth
[policy_effect]
e = some(where (p.eft == allow))
[matchers]
m = p.user == r.user && p.auth == r.auth || r.auth == "super" || p.auth == "super"
`

//权限验证
type EnforcerContext struct {
	e *casbin.Enforcer
}

func NewEnforcerContext(s *storage.Storage) (*EnforcerContext, error) {
	mode, err := model.NewModelFromString(rule)
	if err != nil {
		return nil, err
	}
	ctx := &EnforcerContext{}
	if ctx.e, err = casbin.NewEnforcer(mode, newAdapter(s)); err != nil {
		return nil, err
	}
	return ctx, ctx.SetAuth(conf.SuperName, []string{conf.SuperAuth})
}

//验证权限
func (ctx *EnforcerContext) Enforcer(username, auth string) (ok bool) {
	var err error
	if ok, err = ctx.e.Enforce(username, auth); err != nil {
		return false
	}
	return
}

//是否是超管
func (ctx *EnforcerContext) IsSuper(username string) bool {
	if username == conf.SuperName {
		return true
	}
	for _, auth := range ctx.GetAuth(username) {
		if auth == conf.SuperAuth {
			return true
		}
	}
	return false
}

//获取权限列表
func (ctx *EnforcerContext) GetAuth(username string) []string {
	ps := ctx.e.GetPermissionsForUser(username)
	if len(ps) == 0 {
		return nil
	}
	auth := make([]string, 0, len(ps))
	for _, rs := range ps {
		if len(rs) > 1 && rs[0] == username {
			auth = append(auth, rs[1])
		}
	}
	return auth
}

func (ctx *EnforcerContext) SetAuth(user string, auths []string) (err error) {
	//for _, auth := range auths {
	//	if _, err = ctx.e.AddPolicy(user, auth); err != nil {
	//		return
	//	}
	//}
	//if _, err = ctx.e.AddPolicy(user, auths); err != nil {
	//	return
	//}
	ctx.e.AddPermissionForUser(user,auths...)
	fmt.Println("user == ", user, " auth === ", auths)
	return ctx.e.SavePolicy()
}

func (ctx *EnforcerContext) DelAuth(user string) (ok bool, err error) {
	if ok, err = ctx.e.DeletePermissionsForUser(user); err != nil {
		fmt.Println("err string => ", err.Error())
		return
	}
	if err = ctx.e.SavePolicy(); err != nil {
		ok = false
	}
	return
}

func (ctx *EnforcerContext) Delete(user, auth string) (ok bool, err error) {
	ok, err = ctx.e.DeletePermissionForUser(user, "login")
	return
}
