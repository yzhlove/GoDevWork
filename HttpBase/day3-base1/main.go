package main

import (
	"fmt"
	"strings"
)

func main() {

	n := &node{}
	fmt.Println(n.String())
	str := "/hello/world/yzh"
	ss := strings.Split(str, "/")[1:]
	fmt.Println(ss, " - ", len(ss))
	n.insert(str, ss, 0)
	str = "/hello/world/wyq"
	ss = strings.Split(str, "/")[1:]
	fmt.Println(ss, " - ", len(ss))
	n.insert(str, ss, 0)
	fmt.Println(n.String())
	fmt.Println("====================")
	show(n)
	fmt.Println("====================")
	fmt.Println(n.search([]string{"hello", "world", "wyq"}, 0).String())
}

type node struct {
	pattern  string
	part     string
	children []*node
	isWild   bool
}

func (n *node) String() string {
	return fmt.Sprintf("node{pattern=%s,part=%s,isWild=%t,length=%d}", n.pattern, n.part, n.isWild, len(n.children))
}

func (n *node) insert(pattern string, parts []string, top int) {
	if len(parts) == top {
		n.pattern = pattern
		return
	}
	part := parts[top]
	fmt.Println("part ==> ", part)
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, top+1)
}

func show(n *node) {
	if n == nil {
		return
	}
	fmt.Println(n.String())
	for _, next := range n.children {
		show(next)
	}
}

func (n *node) search(parts []string, top int) *node {
	if len(parts) == top || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}
	part := parts[top]
	children := n.matchChildren(part)
	for _, child := range children {
		result := child.search(parts, top+1)
		if result != nil {
			return result
		}
	}
	return nil
}

func (n *node) matchChild(part string) *node {
	for _, c := range n.children {
		if c.part == part || c.isWild {
			return c
		}
	}
	return nil
}

func (n *node) matchChildren(part string) []*node {
	var nodes []*node
	for _, c := range n.children {
		if c.part == part || c.isWild {
			nodes = append(nodes, c)
		}
	}
	return nodes
}
