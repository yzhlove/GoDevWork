package rules

import (
	"WorkSpace/GoDevWork/Chats/auth/test11/config"
	"WorkSpace/GoDevWork/Chats/auth/test11/storage"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
)

const acl = `
[request_definition]
r = user, auth
[policy_definition]
p = user, auth
[policy_effect]
e = some(where (p.eft == allow))
[matchers]
m = p.user == r.user && p.auth == r.auth || r.auth == "super" || p.auth == "super"
`

type Enforcer struct {
	e *casbin.Enforcer
}

func NewEnforcer(s *storage.Storage) (c *Enforcer, err error) {
	var m model.Model
	if m, err = model.NewModelFromString(acl); err != nil {
		return
	}
	c = &Enforcer{}
	if c.e, err = casbin.NewEnforcer(m, newFileAdapter(s)); err == nil {
		err = s.SaveAuth(config.DefaultSuperName, []string{config.DefaultSuperAuth})
	}
	return
}

func (c *Enforcer) ECheck(user, auth string) (ok bool) {
	ok, err := c.e.Enforce(user, auth)
	if err != nil {
		return false
	}
	return
}

func (c *Enforcer) IsSuper(user string) bool {
	if user == config.DefaultSuperName {
		return true
	}
	for _, auth := range c.GetAuths(user) {
		if auth == config.DefaultSuperAuth {
			return true
		}
	}
	return false
}

func (c *Enforcer) GetAuths(user string) []string {
	as := c.e.GetPermissionsForUser(user)
	if len(as) > 0 {
		auths := make([]string, 0, len(as))
		for _, v := range as {
			if len(v) > 1 {
				auths = append(auths, v[1:]...)
			}
		}
		return auths
	}
	return nil
}

func (c *Enforcer) SetAuths(user string, auth []string) (err error) {
	_, err = c.e.AddPermissionForUser(user, auth...)
	return
}
