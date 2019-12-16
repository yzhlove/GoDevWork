package main

import "github.com/casbin/casbin/v2/model"

func getMode() (model.Model, error) {
	text := `[request_definition]
r = user, resource, auth
[policy_definition]
p = user, resource, auth
[policy_effect]
e = some(where (p.eft == allow))
[matchers]
m = r.user == p.user && r.resource == p.resource && r.auth == p.auth || r.user == "super"
`
	return model.NewModelFromString(text)
}
