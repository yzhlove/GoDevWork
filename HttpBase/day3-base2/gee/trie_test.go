package gee

import (
	"fmt"
	"strings"
	"testing"
)

func TestNode_String(t *testing.T) {

	fmt.Println("Hello Test")
	tNode := &node{}
	str := "/hello/world/yzh"
	ss := strings.Split(str, "/")[1:]
	tNode.insert(str, ss)
	show(tNode)

}
