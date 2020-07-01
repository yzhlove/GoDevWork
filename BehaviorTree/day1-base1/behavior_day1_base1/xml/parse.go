package xml

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"strings"
)

type Node struct {
	Name     string
	Height   int
	Value    strings.Builder
	Params   string
	Parent   *Node
	Children []*Node
}

func ParseXMLDoc(dec *xml.Decoder) (root *Node) {

	var node *Node
	var height = 1
	for {
		token, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic("xml decode error")
		}
		switch value := token.(type) {
		case xml.StartElement:
			if node == nil {
				node = &Node{}
				root = node
			} else {
				tmp := &Node{Parent: node, Height: height}
				node.Children = append(node.Children, tmp)
				node = tmp
				height++
			}
			node.Name = value.Name.Local
			for _, attr := range value.Attr {
				if local := attr.Name.Local; strings.HasPrefix(local, "para") {
					node.Params = attr.Value
				}
			}
		case xml.EndElement:
			if node != nil {
				node = node.Parent
			}
		case xml.CharData:
			if node != nil {
				node.Value.WriteString(strings.TrimSpace(string(value)))
			}
		}
	}
	PrintTree(root, 0)
	return
}

func PrintTree(node *Node, depth int) {
	if depth != 0 {
		fmt.Print("|")
		for i := 0; i < depth; i++ {
			fmt.Print("---")
		}
	}
	fmt.Println(node.Name, node.Height, node.Value.String(), node.Params, " parent = ", node.Parent)
	for _, n := range node.Children {
		PrintTree(n, depth+1)
	}
}
