package gee

import (
	"fmt"
)

type node struct {
	pattern     string
	value       string
	collections []*node
	isWild      bool
}

func (n *node) String() string {
	return fmt.Sprintf("node{pattern=%s,value=%s,isWild=%t}", n.pattern, n.value, n.isWild)
}

func (n *node) insert(pattern string, values []string) {

	_match := func(value string, tNode *node) *node {
		for _, collection := range tNode.collections {
			if collection.value == value || collection.isWild {
				return collection
			}
		}
		return nil
	}
	var tNode = n
	for _, value := range values {
		if tNode = _match(value, tNode); tNode == nil {
			ok := value[0] == ':' || value[0] == '*'
			tNode = &node{value: value, isWild: ok}
			n.collections = append(n.collections, tNode)
		}
	}
	tNode.pattern = pattern
}

func show(tNode *node) {
	if tNode == nil {
		return
	}
	fmt.Println(tNode.String())
	for _, n := range tNode.collections {
		show(n)
	}
}
