package gee

import "testing"

var paths = []string{
	"/",
	"/hello/:name",
	"/hello/b/c",
	"/hi/:name",
	"/assert/*filepath",
}

func newTestRouter() *router {
	r := newRouter()
	for _, path := range paths {
		r.addRouter("GET", path, nil)
	}
	return r
}

func TestParseUlr(t *testing.T) {
	for _, path := range paths {
		t.Log(parseUrl(path))
	}
}
