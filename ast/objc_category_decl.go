package ast

import (
	//	"fmt"
	"strings"
)

// ObjCCategoryDecl is node represents a category declaration.
type ObjCCategoryDecl struct {
	Addr       Address
	Pos        Position
	Position2  string
	Name       string
	ChildNodes []Node
}

func parseObjCCategoryDecl(line string) Node {
	groups := groupsFromRegex(
		`(?:prev (?P<prev>0x[0-9a-f]+) )?
		<(?P<position><invalid sloc>|.*)>
		(?P<position2> <invalid sloc>| col:\d+| line:\d+:\d+)
		(?P<name> \w+)?`,
		line,
	)
	if groups == nil {
		return &Unknown{}
	}

	return &ObjCCategoryDecl{
		Addr: ParseAddress(groups["address"]),
		Pos:  NewPositionFromString(groups["position"]),
		//Position2:    strings.TrimSpace(groups["position2"]),
		Name:       strings.TrimSpace(groups["name"]),
		ChildNodes: []Node{},
	}
}

// AddChild adds a new child node. Child nodes can then be accessed with the
// Children attribute.
func (n *ObjCCategoryDecl) AddChild(node Node) {
	n.ChildNodes = append(n.ChildNodes, node)
}

// Address returns the numeric address of the node. See the documentation for
// the Address type for more information.
func (n *ObjCCategoryDecl) Address() Address {
	return n.Addr
}

// Children returns the child nodes. If this node does not have any children or
// this node does not support children it will always return an empty slice.
func (n *ObjCCategoryDecl) Children() []Node {
	return n.ChildNodes
}

// Position returns the position in the original source code.
func (n *ObjCCategoryDecl) Position() Position {
	return n.Pos
}
