package bt

import (
	"behavior_day1_base1/xml"
	emmm "encoding/xml"
	"errors"
	"os"
	"strings"
)

func CreateTree(xnode *xml.Node) (root NodeInterface) {
	if xnode != nil {
		root = CreateNode(xnode)
		BuildTree(root, xnode)
	} else {
		panic("xnode is nil")
	}
	return
}

func BuildTree(root NodeInterface, xnode *xml.Node) {
	for _, child := range xnode.Children {
		node := CreateNode(child)
		root.AddChild(node)
		node.AddParent(root)
		BuildTree(node, child)
	}
}

func CreateNode(xnode *xml.Node) (node NodeInterface) {
	var err error
	switch xnode.Name {
	case "root":
		node = NewRoot(xnode)
	case "sequence":
		node = NewSequence(xnode)
	case "selector":
		node = NewSelector(xnode)
	case "parallel":
		node = NewParallel(xnode)
	case "skill":
		node, err = NewSkillAction(xnode)
	case "escape":
		node = NewEscapeAction(xnode)
	case "eat":
		node, err = NewEatAction(xnode)
	case "condhp":
		node, err = NewCondHp(xnode)
	}
	if err != nil {
		panic(err.Error())
	}
	return
}

func Source(path string) error {

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	var name string
	if stat, err := os.Stat(path); err != nil {
		return errors.New("path err:" + err.Error())
	} else {
		if i := strings.LastIndexAny(stat.Name(), "."); i == -1 {
			return errors.New("error file")
		} else {
			name = stat.Name()[:i]
		}
	}
	GlobalTreeMap[name] = CreateTree(xml.ParseXMLDoc(emmm.NewDecoder(f)))
	return nil
}
