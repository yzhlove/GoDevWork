package main

import (
	"sync"
	"testing"
)

type Data struct {
	strs map[string]string
	sync.Mutex
}

func (c *Data) Write(k, v string) {
	c.Lock()
	defer c.Unlock()
	c.strs[k] = v
}

func (c *Data) Read(k string) string {
	c.Lock()
	defer c.Unlock()
	return c.strs[k]
}

var pairs = []struct {
	k string
	v string
}{
	{"polaris", " 徐新华 "},
	{"studygolang", "Go 语言中文网 "},
	{"stdlib", "Go 语言标准库 "},
	{"polaris1", " 徐新华 1"},
	{"studygolang1", "Go 语言中文网 1"},
	{"stdlib1", "Go 语言标准库 1"},
	{"polaris2", " 徐新华 2"},
	{"studygolang2", "Go 语言中文网 2"},
	{"stdlib2", "Go 语言标准库 2"},
	{"polaris3", " 徐新华 3"},
	{"studygolang3", "Go 语言中文网 3"},
	{"stdlib3", "Go 语言标准库 3"},
	{"polaris4", " 徐新华 4"},
	{"studygolang4", "Go 语言中文网 4"},
	{"stdlib4", "Go 语言标准库 4"},
}

var td = &Data{strs: make(map[string]string)}

func Test_Write(t *testing.T) {
	t.Parallel()
	for _, tt := range pairs {
		td.Write(tt.k, tt.v)
	}
}

func Test_Read(t *testing.T) {
	t.Parallel()
	for _, tt := range pairs {
		if tt.v != td.Read(tt.k) {
			t.Error("rade value is not equal")
		}
	}
}
