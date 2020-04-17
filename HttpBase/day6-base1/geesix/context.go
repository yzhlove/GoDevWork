package geesix

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

type H map[string]interface{}

type Context struct {
	Resp       http.ResponseWriter
	Req        *http.Request
	Path       string
	Method     string
	Params     map[string]string
	StatusCode int
	handlers   []HandleFunc
	index      int
	engine     *Engine
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Path:   r.URL.Path,
		Method: r.Method,
		Resp:   w,
		Req:    r,
		index:  -1,
	}
}

func (c *Context) Next() {
	c.index++
	for s := len(c.handlers); c.index < s; c.index++ {
		curFunc := c.handlers[c.index]
		fmt.Println("---> ", runtime.FuncForPC(reflect.ValueOf(curFunc).Pointer()).Name())
		c.handlers[c.index](c)
	}
}

func (c *Context) Fail(code int, err string) {
	c.index = len(c.handlers)
	c.JSON(code, H{"message": err})
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

func (c *Context) HTML(code int, name string, data interface{}) {
	c.SetHead("Content-Type", "text/html")
	c.SetCode(code)
	if err := c.engine.htmlTemplates.ExecuteTemplate(c.Resp, name, data); err != nil {
		c.Fail(http.StatusInternalServerError, err.Error())
	}
}
