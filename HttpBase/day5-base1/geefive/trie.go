package geefive

import (
	"fmt"
	"strings"
)

type trieNode struct {
	path  string
	value string
	nodes []*trieNode
	isEnd bool
}

func (n *trieNode) String() string {
	return fmt.Sprintf("trieNode:{path:%s,value:%s,isEnd:%t}", n.path, n.value, n.isEnd)
}

func (n *trieNode) InsertNode(path string, values []string) {
	n.insert(path, values, 0)
}

func (n *trieNode) insert(path string, values []string, topN int) {
	if len(values) == topN {
		n.path = path
		return
	}
	parent := values[topN]
	node := n.matchNode(parent)
	if node == nil {
		isEnd := strings.HasPrefix(parent, ":") || strings.HasPrefix(parent, "*")
		node = &trieNode{value: parent, isEnd: isEnd}
		n.nodes = append(n.nodes, node)
	}
	node.insert(path, values, topN+1)
}

func (n *trieNode) SearchNode(values []string) *trieNode {
	return n.search(values, 0)
}

func (n *trieNode) search(values []string, topN int) *trieNode {
	if len(values) == topN || strings.HasPrefix(n.value, "*") {
		if n.path == "" {
			return nil
		}
		return n
	}
	for _, node := range n.matchNodes(values[topN]) {
		if tNode := node.search(values, topN+1); tNode != nil {
			return tNode
		}
	}
	return nil
}

func (n *trieNode) Travel() []*trieNode {
	nodes := []*trieNode{n}
	var resultNodes []*trieNode
	for i := 0; i < len(nodes); i++ {
		if nodes[i].path != "" {
			resultNodes = append(resultNodes, nodes[i])
		}
		for _, node := range nodes[i].nodes {
			nodes = append(nodes, node)
		}
	}
	return resultNodes
}

func (n *trieNode) matchNode(parent string) *trieNode {
	for _, node := range n.nodes {
		if node.value == parent || node.isEnd {
			return node
		}
	}
	return nil
}

func (n *trieNode) matchNodes(parent string) []*trieNode {
	var nodes []*trieNode
	for _, node := range n.nodes {
		if node.value == parent || node.isEnd {
			nodes = append(nodes, node)
		}
	}
	return nodes
}
