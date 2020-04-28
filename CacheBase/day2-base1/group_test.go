package geecachetwo

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func Test_Getter(t *testing.T) {
	var fn = GetterFunc(func(key string) ([]byte, error) {
		return []byte(key), nil
	})
	expect := []byte("key")
	if v, _ := fn.Get("key"); reflect.DeepEqual(v, expect) {
		t.Log("ok")
	}
}

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func Test_Get(t *testing.T) {
	counts := make(map[string]int, len(db))
	gee := NewGroup("scores", 2<<10, GetterFunc(func(key string) ([]byte, error) {
		log.Println("[SlowDB] search key", key)
		if v, ok := db[key]; ok {
			if _, ok := counts[key]; !ok {
				counts[key] = 0
			}
			counts[key]++
			return []byte(v), nil
		}
		return nil, fmt.Errorf("%s not exist", key)
	}))

	for k, v := range db {
		if view, err := gee.Get(k); err != nil || view.String() != v {
			t.Fatalf("failed key %v of value %v ", k, v)
		}
		if _, err := gee.Get(k); err != nil || counts[k] > 1 {
			t.Fatalf("cache miss %v ", k)
		}
	}

	if view, err := gee.Get("unknown"); err == nil {
		t.Fatalf("the value of : %v ", view)
	}

}

func Test_GetGroup(t *testing.T) {
	name := "scores"
	NewGroup(name, 2<<10, GetterFunc(func(key string) (byte []byte, err error) {
		return
	}))
	if group := GetGroup(name); group == nil || group.name != name {
		t.Fatalf("group not exists by value:%v", name)
	}
	if group := GetGroup(name + "aaa"); group != nil {
		t.Fatalf("expect name by : %v", group.name)
	}
	t.Log("ok.")
}
