package geeseven

import (
	"fmt"
	"strings"
)

type trieTreeNode struct {
	pattern       string
	part          string
	trieTreeNodes []*trieTreeNode
	status        bool
}

func create(part string) *trieTreeNode {
	status := strings.HasPrefix(part, ":") || strings.HasPrefix(part, "*")
	return &trieTreeNode{part: part, status: status}
}

func (tree *trieTreeNode) getNode(part string) *trieTreeNode {
	for _, node := range tree.trieTreeNodes {
		if node.part == part || node.status {
			return node
		}
	}
	return nil
}

func (tree *trieTreeNode) getNodes(part string) []*trieTreeNode {
	var treeNodes []*trieTreeNode
	for _, node := range tree.trieTreeNodes {
		if node.part == part || node.status {
			treeNodes = append(treeNodes, node)
		}
	}
	return treeNodes
}

func (tree *trieTreeNode) insert(pattern string, parts []string, top int) {
	if len(parts) == top {
		tree.pattern = pattern
		return
	}
	part := parts[top]
	root := tree.getNode(part)
	if root == nil {
		root = create(part)
		tree.trieTreeNodes = append(tree.trieTreeNodes, root)
	}
	root.insert(pattern, parts, top+1)
}

func (tree *trieTreeNode) treeInsert(pattern string, parts []string) {
	tree.insert(pattern, parts, 0)
}

func (tree *trieTreeNode) search(parts []string, top int) *trieTreeNode {
	if len(parts) == top || strings.HasPrefix(tree.part, "*") {
		if tree.pattern == "" {
			return nil
		}
		return tree
	}
	for _, root := range tree.getNodes(parts[top]) {
		if node := root.search(parts, top+1); node != nil {
			return node
		}
	}
	return nil
}

func (tree *trieTreeNode) treeSearch(parts []string) *trieTreeNode {
	return tree.search(parts, 0)
}

func (tree *trieTreeNode) String() string {
	return fmt.Sprintf("treeTrieNode:{pattern:%s,part:%s,status:%t}", tree.pattern, tree.part, tree.status)
}
