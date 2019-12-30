package storage

import (
	"WorkSpace/GoDevWork/Chats/auth/test09/conf"
	"github.com/unknwon/goconfig"
	"os"
	"path"
)

/*
配置文件格式
==============================

[super]
passwd=123456789
auth=super

[yzh]
passwd=123456789
auth=login|useradd|user_authset

*/

const (
	password = "passwd"
	auth     = "auth"
)

type Storage struct {
	db *goconfig.ConfigFile
}

func NewStorage() (s *Storage, err error) {
	if _, err = os.Stat(conf.StorageFile); err != nil {
		if err = os.MkdirAll(path.Dir(conf.StorageFile), 0755); err != nil {
			return
		}
		if f, err := os.Create(conf.StorageFile); err != nil {
			return nil, err
		} else {
			f.Close()
		}
	}
	s = &Storage{}
	if s.db, err = goconfig.LoadConfigFile(conf.StorageFile); err != nil {
		return
	}
	//创建超管账号
	s.createSuper()
	return
}

func (s *Storage) createSuper() {
	if _, err := s.db.GetSection(conf.SuperName); err != nil {
		s.db.SetValue(conf.SuperName, password, conf.SuperPasswd)
	}
}

func (s *Storage) LoadAuths() []string {
	return []string{"p, yzh, super", "p, yzh, login", "p, yzh, manager"}
}

func (s *Storage) SaveAuth(auth []string) error {

	return nil
}

//数据存储
func (s *Storage) save() error {
	return goconfig.SaveConfigFile(s.db, conf.StorageFile)
}
