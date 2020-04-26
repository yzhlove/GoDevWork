package main

import (
	"container/list"
	"fmt"
)

//////////////////////////////////////////////
// LFU
//////////////////////////////////////////////

func main() {

	cacheLFU := New(20)
	cacheLFU.Set("1", 1)
	cacheLFU.Set("2", 2)
	cacheLFU.Set("3", 3)
	cacheLFU.Get("3")
	cacheLFU.Get("3")
	cacheLFU.Get("2")
	cacheLFU.Get("1")
	for k, e := range cacheLFU.bykey {
		fmt.Printf("Map -> key:%v Value:%v \n", k, e)
	}

	fmt.Println("========================")

	for e := cacheLFU.freqs.Front(); e != nil; e = e.Next() {
		bucket := e.Value.(*FrequencyItem)
		fmt.Println(bucket.freq, " - ", len(bucket.entries))
	}

}

type CacheItem struct {
	key             string
	value           interface{}
	frequencyParent *list.Element
}

type FrequencyItem struct {
	entries map[*CacheItem]byte
	freq    int
}

type Cache struct {
	bykey    map[string]*CacheItem
	freqs    *list.List
	capacity int
	size     int
}

func New(max int) *Cache {
	return &Cache{
		bykey:    make(map[string]*CacheItem),
		freqs:    list.New(),
		size:     0,
		capacity: max,
	}
}

func (cache *Cache) Set(key string, value interface{}) {
	if item, ok := cache.bykey[key]; ok {
		item.value = value
	} else {
		item = &CacheItem{
			key:   key,
			value: value,
		}
		cache.bykey[key] = item
		cache.size++
		if cache.atCapacity() {
			cache.Evict(10)
		}
		cache.increment(item)
	}
}

func (cache *Cache) Get(key string) interface{} {
	if e, ok := cache.bykey[key]; ok {
		cache.increment(e)
		return e.value
	}
	return nil
}

func (cache *Cache) increment(item *CacheItem) {
	currentFrequency := item.frequencyParent
	var nextFrequencyAmount int
	var nextFrequency *list.Element

	if currentFrequency == nil {
		nextFrequencyAmount = 1
		nextFrequency = cache.freqs.Front()
	} else {
		nextFrequencyAmount = currentFrequency.Value.(*FrequencyItem).freq + 1
		nextFrequency = currentFrequency.Next()
	}

	if nextFrequency == nil || nextFrequency.Value.(*FrequencyItem).freq != nextFrequencyAmount {
		newFrequencyItem := &FrequencyItem{
			freq:    nextFrequencyAmount,
			entries: make(map[*CacheItem]byte),
		}
		if currentFrequency == nil {
			nextFrequency = cache.freqs.PushFront(newFrequencyItem)
		} else {
			nextFrequency = cache.freqs.InsertAfter(newFrequencyItem, currentFrequency)
		}
	}
	item.frequencyParent = nextFrequency
	nextFrequency.Value.(*FrequencyItem).entries[item] = 1
	if currentFrequency != nil {
		cache.Remove(currentFrequency, item)
	}
}

func (cache *Cache) Remove(listItem *list.Element, item *CacheItem) {
	frequencyItem := listItem.Value.(*FrequencyItem)
	delete(frequencyItem.entries, item)
	if len(frequencyItem.entries) == 0 {
		cache.freqs.Remove(listItem)
	}
}

func (cache *Cache) Evict(count int) {
	for i := 0; i < count; {
		if item := cache.freqs.Front(); item != nil {
			for entity := range item.Value.(*FrequencyItem).entries {
				if i < count {
					delete(cache.bykey, entity.key)
					cache.Remove(item, entity)
					cache.size--
					i++
				}
			}
		}
	}
}

func (cache *Cache) atCapacity() bool {
	return cache.size >= cache.capacity
}
