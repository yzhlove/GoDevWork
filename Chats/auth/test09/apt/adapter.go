package apt

import (
	"WorkSpace/GoDevWork/Chats/auth/test09/storage"
	"errors"
	"fmt"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"strings"
)

var storageError = errors.New("storage load error")

type FileAdapter struct {
	s *storage.Storage
}

func newAdapter(s *storage.Storage) *FileAdapter {
	return &FileAdapter{s: s}
}

// LoadPolicyLine loads a text line as a policy rule to model.
func localLoadPolicyLine(line string, model model.Model) {
	if line == "" || strings.HasPrefix(line, "#") {
		return
	}

	tokens := strings.Split(line, ",")
	fmt.Println("tokens ==> ", tokens)
	for i := 0; i < len(tokens); i++ {
		tokens[i] = strings.TrimSpace(tokens[i])
	}
	key := tokens[0]
	sec := key[:1]
	fmt.Println(" key ==> ", key, " sec ==> ", sec)
	model[sec][key].Policy = append(model[sec][key].Policy, tokens[1:])
}

func (a *FileAdapter) LoadPolicy(model model.Model) error {
	if a.s == nil {
		return storageError
	}
	for _, rule := range a.s.LoadAuths() {
		fmt.Println("rule => ", rule)
		//localLoadPolicyLine(strings.TrimSpace(rule), model)
		persist.LoadPolicyLine(rule, model)
	}

	for k, v := range model {
		fmt.Println("k ==> ", k)
		for kk, vv := range v {
			fmt.Println("kk => ", kk, " vv => ", vv)
		}
	}
	fmt.Println("=====================================")
	return nil
}

func (a *FileAdapter) SavePolicy(model model.Model) error {
	for k, v := range model {
		fmt.Println("k ==> ", k)
		for kk, vv := range v {
			fmt.Println("kk => ", kk, " vv => ", vv)
		}
	}
	return nil
}

// AddPolicy adds a policy rule to the storage.
func (a *FileAdapter) AddPolicy(sec string, ptype string, rule []string) error {
	fmt.Println("sec => ", sec, " ptype => ", ptype, " rules => ", rule)

	return errors.New("not implemented")
}

// RemovePolicy removes a policy rule from the storage.
func (a *FileAdapter) RemovePolicy(sec string, ptype string, rule []string) error {
	fmt.Println("===============================")
	fmt.Println("delete ===>  sec => ", sec, " ptype => ", ptype, " rule => ", rule)
	return errors.New("not implemented")
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
func (a *FileAdapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	fmt.Println("===============================  222 ")
	fmt.Println("sec ==> ", sec)
	fmt.Println("ptype ==> ", ptype)
	fmt.Println("fieldIndex ==> ", fieldIndex)
	fmt.Println("fieldValues ==> ", fieldValues)
	return errors.New("not implemented")
}
