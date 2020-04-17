package gee

import (
	"fmt"
	"strings"
)

type trieNode struct {
	path        string
	value       string
	collections []*trieNode
	isWild      bool
}

func (n *trieNode) String() string {
	return fmt.Sprintf("trieNode{path=%s,value=%s,iswild=%t}", n.path, n.value, n.isWild)
}

func (n *trieNode) Insert(path string, values []string) {
	n.insert(path, values, 0)
}

func (n *trieNode) insert(path string, values []string, top int) {
	if len(values) == top {
		n.path = path
		return
	}
	value := values[top]
	node := n.matchNode(value)
	if node == nil {
		fmt.Println("+++++++++++++++++++++++++")
		ok := strings.HasPrefix(value, ":") || strings.HasPrefix(value, "*")
		node = &trieNode{value: value, isWild: ok}
		fmt.Println("insert node => ", node.String())
		n.collections = append(n.collections, node)
	}
	node.insert(path, values, top+1)
}

func Show(n *trieNode) {
	fmt.Println(n.String())
	for _, node := range n.collections {
		Show(node)
	}
}

func (n *trieNode) Search(values []string) *trieNode {
	return n.search(values, 0)
}

func (n *trieNode) search(values []string, top int) *trieNode {
	if len(values) == top || strings.HasPrefix(n.value, "*") {
		if n.path == "" {
			return nil
		}
		return n
	}
	value := values[top]
	fmt.Println("search value => ", value, " >>>>>>>>>>>>>>>>>>")
	Show(n)
	nodes := n.matchNodes(value)
	fmt.Println("nodes => ", nodes)
	for _, node := range nodes {
		if result := node.search(values, top+1); result != nil {
			return result
		}
	}
	return nil
}

func (n *trieNode) travel() []*trieNode {
	nodes := []*trieNode{n}
	var values []*trieNode
	for i := 0; i < len(nodes); i++ {
		if len(nodes[i].path) > 0 {
			values = append(values, nodes[i])
		}
		for _, node := range nodes[i].collections {
			nodes = append(nodes, node)
		}
	}
	return nodes
}

func (n *trieNode) matchNode(value string) *trieNode {
	for _, node := range n.collections {
		if node.value == value || node.isWild {
			return node
		}
	}
	return nil
}

func (n *trieNode) matchNodes(value string) []*trieNode {
	var nodes []*trieNode
	for _, node := range n.collections {
		if node.value == value || node.isWild {
			nodes = append(nodes, node)
		}
	}
	return nodes
}
