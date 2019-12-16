package acl

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	adapter "github.com/casbin/casbin/v2/persist/file-adapter"
)

const (
	mode   = "Chats/auth/test04_auth/acl/mode.conf"
	policy = "Chats/auth/test04_auth/acl/policy.csv"
)

func NewCheckAdapter() (*casbin.Enforcer, error) {
	return casbin.NewEnforcer(mode, policy)
}

func NewCheckAdpater2() (*casbin.Enforcer, error) {
	m, err := model.NewModelFromFile(mode)
	if err != nil {
		return nil, err
	}
	apt := adapter.NewAdapter(policy)
	return casbin.NewEnforcer(m, apt)
}
