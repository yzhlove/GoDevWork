package geeone

import (
	"fmt"
	"testing"
)

type StdString string

func (str StdString) Len() int {
	return len(str)
}

func Test_Get(t *testing.T) {

	LRU := New(0, nil)
	LRU.Set("key1", StdString("12345"))
	LRU.Show()
	if v, ok := LRU.Get("key1"); !ok || string(v.(StdString)) != "12345" {
		t.Error("value failed ")
		return
	}
	if _, ok := LRU.Get("key2"); ok {
		t.Error("failed key...")
		return
	}
	t.Log("ok.")
}

func Test_Del(t *testing.T) {
	k1, k2, k3 := "key1", "key2", "k3"
	v1, v2, v3 := "value1", "value2", "v3"
	c := len(k1 + k2 + v1 + v2)
	LRU := New(int64(c), nil)
	LRU.Set(k1, StdString(v1))
	LRU.Set(k2, StdString(v2))
	LRU.Set(k3, StdString(v3))

	LRU.Show()
}

func Test_Callback(t *testing.T) {

	keys := make([]string, 0)
	event := func(key string, value Value) {
		keys = append(keys, key)
		keys = append(keys, string(value.(StdString)))
	}

	LRU := New(int64(10), event)
	LRU.Set("key1", StdString("123456"))
	LRU.Set("k2", StdString("v2"))
	LRU.Set("k3", StdString("v3"))
	LRU.Set("k4", StdString("v4"))
	LRU.Set("k5", StdString("v5"))

	fmt.Println(keys)

}

func Test_Set(t *testing.T) {
	LRU := New(0, nil)
	LRU.Set("key", StdString("1"))
	LRU.Set("key", StdString("1111"))
	t.Log(LRU.cntBytes)
	LRU.Set("key1", StdString("11"))
	t.Log(LRU.cntBytes)
}
