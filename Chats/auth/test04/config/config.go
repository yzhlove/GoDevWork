package config

const ModeValue = `
[request_definition]
r = user, resource, auth

[policy_definition]
p = user, resource, auth

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.user == p.user && r.resource == p.resource && r.auth == p.auth || r.user == "super" 
`
