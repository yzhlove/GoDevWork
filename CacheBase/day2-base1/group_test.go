package geecachetwo

import (
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
		if v ,ok := db[key];ok {
			if _ ,ok := counts[key];!ok {
				counts[key] = 0
			}
			counts[key]++
			return []byte(v) , nil
		}
		return nil, nil
	}))
}
