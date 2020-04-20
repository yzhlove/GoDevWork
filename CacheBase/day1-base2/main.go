package main

import "fmt"

////////////////////////////////
// LRU
////////////////////////////////

type Node struct {
	Key   int
	Value int
	prev  *Node
	next  *Node
}

type LRUCache struct {
	limit   int
	HashMap map[int]*Node
	head    *Node
	end     *Node
}

func Construct(capacity int) *LRUCache {
	return &LRUCache{
		limit:   capacity,
		HashMap: make(map[int]*Node, capacity),
	}
}

func (l *LRUCache) Get(key int) int {
	if v, ok := l.HashMap[key]; ok {
		l.refresh(v)
		return v.Value
	} else {
		return -1
	}
}

func (l *LRUCache) Put(key, value int) {
	if v, ok := l.HashMap[key]; !ok {
		if len(l.HashMap) >= l.limit {
			oldKey := l.remove(l.head)
			delete(l.HashMap, oldKey)
		}
		node := &Node{Key: key, Value: value}
		l.add(node)
		l.HashMap[key] = node
		fmt.Printf("--> node:{key:%d value:%d }\n", node.Key, node.Value)
	} else {
		v.Value = value
		l.refresh(v)
	}
}

func (l *LRUCache) refresh(node *Node) {
	if node == l.end {
		return
	}
	l.remove(node)
	l.add(node)
}

func (l *LRUCache) remove(node *Node) int {
	if node == l.head {
		l.head = l.head.next
		l.head.prev = nil
	} else if node == l.end {
		l.end = l.end.prev
		l.end.next = nil
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}
	return node.Key
}

func (l *LRUCache) add(node *Node) {
	if l.end != nil {
		l.end.next = node
		node.prev = l.end
		node.next = nil
	}
	l.end = node
	if l.head == nil {
		l.head = node
	}
}

func (l *LRUCache) GetCache() {
	for node := l.head; node != nil; node = node.next {
		fmt.Printf("node:{key:%d value:%d }\n", node.Key, node.Value)
	}
}

func (l *LRUCache) GetHashMap() {
	for key, value := range l.HashMap {
		fmt.Printf("hashmap:{key:%d value:%d}\n", key, value)
	}
}

func main() {

	cache := Construct(3)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	cache.GetCache()
	//cache.GetHashMap()
	cache.Put(4, 4)
	fmt.Println("=====================================")
	cache.GetCache()
	cache.Put(2, 2)
	fmt.Println("=====================================")
	cache.GetCache()
	cache.Put(1, 1)
	fmt.Println("=====================================")
	cache.GetCache()
	cache.Put(4, 4)
	fmt.Println("=====================================")
	cache.GetCache()
	cache.Put(2, 2)
	fmt.Println("=====================================")
	cache.GetCache()
}
