package main

import (
	"errors"
	"fmt"
	"github.com/casbin/casbin/util"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"github.com/unknwon/goconfig"
	"strings"
)

type Adapter struct {
	path string
	conf *goconfig.ConfigFile
}

func NewAdapter(path string) (apt *Adapter, err error) {
	apt = new(Adapter)
	apt.path = path
	if apt.conf, err = goconfig.LoadConfigFile(path); err != nil {
		return
	}
	return
}

func (a *Adapter) LoadPolicy(model model.Model) error {
	if a.path == "" {
		return errors.New("invalid file path, file path cannot be empty")
	}

	return a.loadPolicyFile(model, persist.LoadPolicyLine)
}

func (a *Adapter) loadPolicyFile(model model.Model, handler func(string, model.Model)) error {

	rules, _ := a.conf.GetSection("rules")

	fmt.Println("load file => ", rules)
	for _, rule := range rules {
		line := strings.TrimSpace(rule)
		handler(line, model)
	}
	return nil
}

func (a *Adapter) SavePolicy(model model.Model) error {

	if a.path == "" {
		return errors.New("invalid file path, file path cannot be empty")
	}

	rules := make([]string, 0, 8)
	var buf strings.Builder
	for ptype, ast := range model["p"] {
		for _, rule := range ast.Policy {
			buf.WriteString(ptype + ", ")
			buf.WriteString(util.ArrayToString(rule))
			rules = append(rules, buf.String())
			buf.Reset()
		}
	}

	return a.savePolicyFile(rules)
}

func (a *Adapter) savePolicyFile(rules []string) error {

	a.conf.DeleteSection("rules")

	for _, rule := range rules {
		fmt.Println("set str => ", rule)
		a.conf.SetValue("rules", "-", strings.Trim(rule, "\n"))
	}
	return goconfig.SaveConfigFile(a.conf, a.path)
}

// AddPolicy adds a policy rule to the storage.
func (a *Adapter) AddPolicy(sec string, ptype string, rule []string) error {
	return errors.New("not implemented")
}

// RemovePolicy removes a policy rule from the storage.
func (a *Adapter) RemovePolicy(sec string, ptype string, rule []string) error {
	return errors.New("not implemented")
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
func (a *Adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	return errors.New("not implemented")
}
