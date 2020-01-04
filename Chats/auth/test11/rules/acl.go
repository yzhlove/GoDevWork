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
m = p.user == r.user && p.auth == r.auth || p.user == r.user && p.auth == "super" || r.user == "super"
`

type EnforcerManager struct {
	e       *casbin.Enforcer
	storage *storage.Storage
}

func NewEnforcer(s *storage.Storage) (c *EnforcerManager, err error) {
	m, err := model.NewModelFromString(acl)
	if err != nil {
		return nil, err
	}
	c = &EnforcerManager{storage: s}
	if c.e, err = casbin.NewEnforcer(m, newFileAdapter(s)); err != nil {
		return nil, err
	}
	return
}

func (c *EnforcerManager) Check(user, auth string) (ok bool) {
	ok, err := c.e.Enforce(user, auth)
	if err != nil {
		return false
	}
	return
}

func (c *EnforcerManager) IsSuper(user string) bool {
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

func (c *EnforcerManager) GetAuths(user string) []string {
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

func (c *EnforcerManager) SetAuths(user string, auths []string) (err error) {
	for _, auth := range auths {
		if _, err = c.e.AddPermissionForUser(user, auth); err != nil {
			return err
		}
	}
	return c.storage.SaveAuth(user, auths)
}
