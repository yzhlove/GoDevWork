package acl

import "github.com/casbin/casbin/v2"

const (
	mode   = "Chats/auth/test04_auth/acl/mode.conf"
	policy = "Chats/auth/test04_auth/acl/policy.csv"
)

func NewCheckAdapter() (*casbin.Enforcer, error) {
	return casbin.NewEnforcer(mode, policy)
}
