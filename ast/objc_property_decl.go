package ast

import (
	"strings"
)

// ObjCPropertyDecl is node represents an Objective-C property declaration
type ObjCPropertyDecl struct {
	Addr         Address
	Pos          Position
	Position2    string
	Name         string
	Type         string
	Type2        string
        Attr         string
	ChildNodes   []Node
}

func parseObjCPropertyDecl(line string) Node {
	groups := groupsFromRegex(
                `(?:prev (?P<prev>0x[0-9a-f]+) )?
		<(?P<position>.*<scratch space>.*?|.*<built-in>.*?|.*<invalid sloc>|.*?)>
		(?P<position2> <invalid sloc>| col:\d+| line:\d+:\d+)?
		(?P<name>.*?)
		'(?P<type>[^']*?)'
		(:'(?P<type2>.*?)')?
		(?P<attr> .*)?`,
		line,
	)
	if groups == nil {
		return &Unknown{}
	}

	return &ObjCPropertyDecl{
		Addr:         ParseAddress(groups["address"]),
		Pos:          NewPositionFromString(groups["position"]),
		Position2:    strings.TrimSpace(groups["position2"]),
		Name:         strings.TrimSpace(groups["name"]),
		Type:         groups["type"],
		//Type2:        strings.TrimSpace(groups["type2"]),
		Attr:         strings.TrimSpace(groups["attr"]),
		ChildNodes:   []Node{},
	}
}

// AddChild adds a new child node. Child nodes can then be accessed with the
// Children attribute.
func (n *ObjCPropertyDecl) AddChild(node Node) {
	n.ChildNodes = append(n.ChildNodes, node)
}

// Address returns the numeric address of the node. See the documentation for
// the Address type for more information.
func (n *ObjCPropertyDecl) Address() Address {
	return n.Addr
}

// Children returns the child nodes. If this node does not have any children or
// this node does not support children it will always return an empty slice.
func (n *ObjCPropertyDecl) Children() []Node {
	return n.ChildNodes
}

// Position returns the position in the original source code.
func (n *ObjCPropertyDecl) Position() Position {
	return n.Pos
}
