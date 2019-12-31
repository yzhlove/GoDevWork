package main

import (
	"github.com/unknwon/goconfig"
	"strings"
)

//转化格式

func main() {

	before, err := goconfig.LoadConfigFile("Chats/auth/test10/old.ini")
	if err != nil {
		panic(err)
	}

	after, err := goconfig.LoadConfigFile("Chats/auth/test10/new.ini")
	if err != nil {
		panic(err)
	}

	userMap, err := before.GetSection("users")
	if err != nil {
		panic(err)
	}

	//设置用户
	for user, passwd := range userMap {
		after.SetValue(user, "passwd", passwd)
	}

	//设置权限
	ruleMap, err := before.GetSection("rules")
	if err != nil {
		panic(err)
	}

	userRuleMap := make(map[string][]string, 4)
	for _, rule := range ruleMap {
		tokens := strings.Split(rule, ",")
		for i, token := range tokens {
			tokens[i] = strings.TrimSpace(token)
		}
		if len(tokens) >= 3 {
			userRuleMap[tokens[1]] = append(userRuleMap[tokens[1]], tokens[2])
		}
	}

	for user, rules := range userRuleMap {
		if _, err := after.GetSection(user); err == nil {
			after.SetValue(user, "auth", strings.TrimRight(strings.Join(rules, "|"), "|"))
		}
	}

	_ = goconfig.SaveConfigFile(after, "Chats/auth/test10/new.ini")
}
