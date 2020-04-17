package main

import "fmt"

/////////////////////////////////////
//trie tree 树实现
/////////////////////////////////////

func main() {

	tree := NewTrieTree()
	tree.InsertWorld("我和我的祖国")
	tree.InsertWorld("我爱我的祖国")
	tree.InsertWorld("我和我的家乡")
	fmt.Println(tree.Contains("我"))
	fmt.Println(tree.Contains("我爱"))
	fmt.Println(tree.Contains("我和"))
	fmt.Println(tree.Contains("我和我的家乡"))
	fmt.Println(tree.IsPrefix("我"))
	fmt.Println(tree.IsPrefix("我爱"))
	fmt.Println(tree.IsPrefix("我和"))
	fmt.Println(tree.GetSize())
}

type trieNode struct {
	isWorld bool
	next    map[rune]*trieNode
}

type trieTree struct {
	size int
	root *trieNode
}

func newNode() *trieNode {
	return &trieNode{isWorld: false, next: make(map[rune]*trieNode)}
}

func NewTrieTree() *trieTree {
	return &trieTree{
		size: 0,
		root: &trieNode{isWorld: false, next: make(map[rune]*trieNode)},
	}
}

func (t *trieTree) GetSize() int {
	return t.size
}

func (t *trieTree) InsertWorld(word string) {
	if len(word) == 0 {
		return
	}
	root := t.root
	for _, w := range []rune(word) {
		if _, ok := root.next[w]; !ok {
			root.next[w] = newNode()
		}
		root = root.next[w]
	}
	if !root.isWorld {
		root.isWorld = true
		t.size++
	}
}

func (t *trieTree) Contains(word string) bool {
	if len(word) == 0 {
		return false
	}
	root := t.root
	for _, w := range []rune(word) {
		if node, ok := root.next[w]; ok {
			root = node
		} else {
			return false
		}
	}
	return root.isWorld
}

func (t *trieTree) IsPrefix(word string) bool {
	if len(word) == 0 {
		return false
	}
	root := t.root
	for _, w := range []rune(word) {
		if node, ok := root.next[w]; ok {
			root = node
		} else {
			return false
		}
	}
	return true
}
