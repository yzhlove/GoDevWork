package acl

import (
	"WorkSpace/GoDevWork/Chats/auth/test04/config"
	"github.com/casbin/casbin"
	"github.com/casbin/casbin/model"
)

func GetACLMode() model.Model {
	return casbin.NewModel(config.ModeValue)
}
