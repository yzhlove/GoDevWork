package rules

import (
	"WorkSpace/GoDevWork/Chats/auth/test11/storage"
	"errors"
	"github.com/casbin/casbin/v2/model"
)

var storageError = errors.New("storage load err")

type FileAdapter struct {
	s *storage.Storage
}

func newFileAdapter(s *storage.Storage) *FileAdapter {
	return &FileAdapter{s: s}
}

func (a *FileAdapter) LoadPolicy(model model.Model) error {
	if a.s != nil {
		for user, auths := range a.s.LoadAuth() {
			if len(auths) > 1 {
				key := auths[0]
				sec := key[:1]
				for _, v := range auths[1:] {
					model[sec][key].Policy = append(model[sec][key].Policy, []string{user, v})
				}
			}
		}
	}
	return nil
}

func (a *FileAdapter) SavePolicy(_ model.Model) error {
	return nil
}

func (a *FileAdapter) AddPolicy(_ string, _ string, rule []string) error {
	if len(rule) > 1 {
		return a.s.SaveAuth(rule[0], rule[1:])
	}
	return nil
}

func (a *FileAdapter) RemovePolicy(sec string, ptype string, rule []string) error {
	return errors.New("not implemented")
}

func (a *FileAdapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	return errors.New("not implemented")
}
