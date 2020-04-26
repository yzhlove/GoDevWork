package main

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	LFU := New(10)
	LFU.Set("1", 1)
	LFU.Set("2", 2)
	LFU.ShowBuckets()
	fmt.Println("=========================================")
	LFU.ShowValues()
}

type Entry struct {
	key       string
	value     interface{}
	bucketTop *list.Element //指向Bucket
}

func (e *Entry) String() string {
	return fmt.Sprintf("Entry:{key:%v value:%v Top:%v}", e.key, e.value, e.bucketTop.Value.(*Bucket).index)
}

type Bucket struct {
	entries map[*Entry]struct{} //存放计数次数相同的entry
	index   int                 //计数次数
}

func (b *Bucket) String() string {
	sb := strings.Builder{}
	sb.WriteString("=========(" + strconv.Itoa(b.index) + ")=========\n")
	for e := range b.entries {
		sb.WriteString("\t" + e.String() + "\n")
	}
	return sb.String()
}

type Cache struct {
	values   map[string]*Entry
	buckets  *list.List
	capacity int
	size     int
	count    int //由于size>cap，需要delete
}

func (c *Cache) ShowBuckets() {
	fmt.Println("ShowBuckets Length ==> ", c.buckets.Len())
	for e := c.buckets.Front(); e != nil; e = e.Next() {
		bucket := e.Value.(*Bucket)
		fmt.Printf("=================[%d]=================\n", bucket.index)
		for entry := range bucket.entries {
			fmt.Printf("\t\t %v \n", entry.String())
		}
	}
}

func (c *Cache) ShowValues() {
	for key, e := range c.values {
		fmt.Printf("key:%v Entry:%v \n", key, e.String())
	}
}

func New(capacity int) *Cache {
	return &Cache{
		values:   make(map[string]*Entry),
		buckets:  list.New(),
		capacity: capacity,
		count:    1, //default delete to 10
	}
}

func (c *Cache) Set(key string, value interface{}) {
	if entry, ok := c.values[key]; ok {
		entry.value = value
	} else {
		e := &Entry{key: key, value: value}
		c.values[key] = e
		c.size++
		if c.isMax() {
			c.Evict(c.count)
		}
		c.increment(e)
	}
}

func (c *Cache) Get(key string) interface{} {
	if e, ok := c.values[key]; ok {
		c.increment(e)
		return e.value
	}
	return nil
}

func (c *Cache) increment(entry *Entry) {

	top := entry.bucketTop

	var nextIndex int
	var nextBucket *list.Element

	if top == nil { //add entry
		nextIndex = 1
		nextBucket = c.buckets.Front()
	} else { // add bucket index
		nextIndex = top.Value.(*Bucket).index + 1
		nextBucket = top.Next()
	}

	if nextBucket == nil || nextBucket.Value.(*Bucket).index != nextIndex {
		newBucket := &Bucket{
			entries: make(map[*Entry]struct{}),
			index:   nextIndex,
		}
		if top == nil { //font 节点没有
			nextBucket = c.buckets.PushFront(newBucket)
		} else {
			//在bucketTop节点后插入newBucket
			nextBucket = c.buckets.InsertAfter(newBucket, top)
		}
	}

	entry.bucketTop = nextBucket
	nextBucket.Value.(*Bucket).entries[entry] = struct{}{}
	if top != nil { //add entry
		//需要从旧的bucket里面删除entry
		c.Del(top, entry)
	}
}

func (c *Cache) Del(oldBucket *list.Element, entry *Entry) {
	bucket := oldBucket.Value.(*Bucket)
	delete(bucket.entries, entry)
	if len(bucket.entries) == 0 {
		c.buckets.Remove(oldBucket)
	}
}

func (c *Cache) Evict(count int) {
	for i := 0; i < count; {
		if bucket := c.buckets.Front(); bucket != nil {
			for entry := range bucket.Value.(*Bucket).entries {
				if i < count {
					delete(c.values, entry.key)
					c.Del(bucket, entry)
					c.size--
					i++
				}
			}
		}
	}
}

func (c *Cache) isMax() bool {
	return c.size > c.capacity
}
