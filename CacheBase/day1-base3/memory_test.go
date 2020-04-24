package main

import (
	"log"
	"testing"
)

func Test_Set(t *testing.T) {
	cache := NewMemoryCache(0)
	values := []string{"test1", "test2", "test3", "test4"}
	key := "key1"
	for _, v := range values {
		cache.Set(key, v)
		if value, ok := cache.Get(key); !ok {
			t.Fatalf("except key:%v ,value:%v ", key, value)
		} else if ok && value != v {
			t.Fatalf("%v except value :%v ,get value:%v ", key, v, value)
		} else {
			t.Logf("value:%v ", value)
		}
	}
	t.Logf("list:%v length:%v ,max:%v", len(cache.cache), cache.cacheList.Len(), cache.maxItemSize)
}

var getTest = []struct {
	name       string
	keyToAdd   string
	keyToGet   string
	expectedOk bool
}{
	{"string_hits", "myKey", "myKey", true},
	{"string_miss", "myKey", "nonsense", false},
}

func Test_Get(t *testing.T) {
	c := NewMemoryCache(0)
	for _, tt := range getTest {
		c.Set(tt.keyToAdd, 1234)
		value, ok := c.Get(tt.keyToGet)
		t.Logf("%s: val:%v cache hit = %v;want %v ", tt.name, value, ok, tt.expectedOk)
	}
}

func Test_Del(t *testing.T) {
	c := NewMemoryCache(0)
	key := "myKey"
	c.Set(key, 1234)
	if value, ok := c.Get(key); !ok {
		t.Error("not found")
		return
	} else if value != 1234 {
		t.Error("value is err")
		return
	}
	c.Del(key)
	if _, ok := c.Get(key); ok {
		t.Error("delete err")
		return
	}
	t.Logf("gets: %v hits: %v \n", c.gets.Get(), c.hits.Get())
	t.Log("ok.")
}

func Test_Status(t *testing.T) {

	keys := []string{"1", "2", "3", "4", "5", "a", "b", "c", "d", "e"}
	var gets, hits, max, cnt int64
	max = 5
	c := NewMemoryCache(int(max))
	for _, key := range keys {
		c.Set(key, 1234)
		cnt++
	}

	newkeys := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	for _, key := range newkeys {
		if _, ok := c.Get(key); ok {
			hits++
		}
		gets++
	}

	t.Logf("gets:%v hits:%v max:%v cnt:%v ", gets, hits, max, cnt)
	t.Logf("status: %v ", c.Status())
}

func Test_LRU(t *testing.T) {

	keys := []string{"1", "2", "3", "4", "2", "1", "3", "5", "6", "5", "6"}
	max := 3
	c := NewMemoryCache(max)

	for _, key := range keys {
		c.Set(key, 1234)
		log.Println("============== set key: ", key)
		c.ShowValue()
	}

}
