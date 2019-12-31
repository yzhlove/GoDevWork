package storage

import (
	"WorkSpace/GoDevWork/Chats/auth/test11/config"
	"github.com/unknwon/goconfig"
	"os"
	"path"
	"strings"
)

type Storage struct {
	db *goconfig.ConfigFile
}

func NewStorage() (s *Storage, err error) {
	if _, err = os.Stat(config.ACLFilePath); err != nil {
		if err = os.MkdirAll(path.Dir(config.ACLFilePath), 0755); err != nil {
			return
		}
		f, err := os.Create(config.ACLFilePath)
		if err != nil {
			return nil, err
		}
		f.Close()
	}
	s = &Storage{}
	if s.db, err = goconfig.LoadConfigFile(config.ACLFilePath); err == nil {
		//create super user
		err = s.Submit(config.DefaultSuperName, config.DefaultSuperPasswd)
	}
	return
}

func (s *Storage) LoadAuth() map[string][]string {
	auths := make(map[string][]string, 8)
	t := []string{"p"}
	for _, user := range s.db.GetSectionList() {
		if value, err := s.db.GetValue(user, "auth"); err == nil {
			as := strings.Split(strings.TrimSpace(value), "|")
			auths[user] = append(auths[user], append(t, as...)...)
		}
	}
	return auths
}

func (s *Storage) SaveAuth(user string, auth []string) error {
	authString := strings.TrimRight(strings.Join(auth, "|"), "|")
	s.db.SetValue(user, "auth", authString)
	return s.save()
}

func (s *Storage) Submit(user, passwd string) error {
	s.db.SetValue(user, "passwd", passwd)
	return s.save()
}

func (s *Storage) Delete(user string) error {
	s.db.DeleteSection(user)
	return s.save()
}

func (s *Storage) Exists(user string) bool {
	if _, err := s.db.GetSection(user); err != nil {
		return false
	}
	return true
}

func (s *Storage) GetUserList() []string {

	return nil
}

func (s *Storage) GetPasswd() string {

	return ""
}

func (s *Storage) save() error {
	return goconfig.SaveConfigFile(s.db, config.ACLFilePath)
}
