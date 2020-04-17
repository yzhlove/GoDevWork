package geefour

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	Resp       http.ResponseWriter
	Req        *http.Request
	Path       string
	Method     string
	Params     map[string]string
	StatusCode int
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Resp:   w,
		Req:    r,
		Path:   r.URL.Path,
		Method: r.Method,
	}
}

func (c *Context) Param(key string) string {
	return c.Params[key]
}

func (c *Context) PostFrom(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) SetCode(code int) {
	c.StatusCode = code
	c.Resp.WriteHeader(code)
}

func (c *Context) SetHead(key, value string) {
	c.Resp.Header().Set(key, value)
}

func (c *Context) String(code int, format string, a ...interface{}) {
	c.SetHead("Content-Type", "text/plain")
	c.SetCode(code)
	c.Resp.Write([]byte(fmt.Sprintf(format, a...)))
}

func (c *Context) JSON(code int, value interface{}) {
	c.SetHead("Content-Type", "application/json")
	c.SetCode(code)
	encoder := json.NewEncoder(c.Resp)
	if err := encoder.Encode(value); err != nil {
		http.Error(c.Resp, err.Error(), http.StatusInternalServerError)
	}
}

func (c *Context) BinaryValue(code int, value []byte) {
	c.SetCode(code)
	c.Resp.Write(value)
}

func (c *Context) HTML(code int, html string) {
	c.SetHead("Content-Type", "text/html")
	c.SetCode(code)
	c.Resp.Write([]byte(html))
}
