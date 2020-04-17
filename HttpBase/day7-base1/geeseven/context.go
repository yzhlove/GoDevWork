package geeseven

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"runtime"
)

type Context struct {
	Resp       http.ResponseWriter
	Req        *http.Request
	Path       string
	Method     string
	Params     map[string]string
	StatusCode int
	handles    []HandleFunc
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
	for s := len(c.handles); c.index < s; c.index++ {
		handFunc := c.handles[c.index]
		log.Println("func --->", runtime.FuncForPC(reflect.ValueOf(handFunc).Pointer()).Name())
		handFunc(c)
	}
}

func (c *Context) Error(code int, err string) {
	c.index = len(c.handles)

}

func (c *Context) GetParamValue(key string) string {
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
	c.Resp.WriteHeader(c.StatusCode)
}

func (c *Context) SetHeader(key, value string) {
	c.Resp.Header().Set(key, value)
}

func (c *Context) String(code int, format string, a ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.SetCode(code)
	c.Resp.Write([]byte(fmt.Sprintf(format, a...)))
}

func (c *Context) JSON(code int, v interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.SetCode(code)
	encoder := json.NewEncoder(c.Resp)
	if err := encoder.Encode(v); err != nil {
		http.Error(c.Resp, err.Error(), http.StatusInternalServerError)
	}
}

func (c *Context) Text(code int, v []byte) {
	c.SetCode(code)
	c.Resp.Write(v)
}

func (c *Context) For404Func() HandleFunc {
	return func(ctx *Context) {
		ctx.String(http.StatusNotFound, "404 Not Found %s", ctx.Path)
	}
}

func (c *Context) HTML(code int, name string, v interface{}) {
	c.SetHeader("Content-Type", "text/html")
	c.SetCode(code)

}
