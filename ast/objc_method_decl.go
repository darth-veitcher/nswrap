package ast

import (
	"strings"
)

// ObjCMethodDecl is node represents an Objective-C method declaration
type ObjCMethodDecl struct {
	Addr        Address
	Pos         Position
	Position2   string
	Implicit    bool
	ClassMethod bool
	Name        string
	Parameters  []string
	Type        string
	Type2       string
	Attr        string
	ChildNodes  []Node
}

func parseObjCMethodDecl(line string) Node {
	groups := groupsFromRegex(
		`(?:prev (?P<prev>0x[0-9a-f]+) )?
		<(?P<position>.*<scratch space>.*?|.*<built-in>.*?|.*<invalid sloc>|.*?)>
		(?P<position2> <invalid sloc>| col:\d+| line:\d+:\d+)?
		(?P<implicit> implicit)?
		(?P<methodtype> \+| \-)
		(?P<names>.*?)
		'(?P<type>[^']*?)'
		(:'(?P<type2>.*?)')?
		(?P<attr> .*)?`,
		line,
	)
	if groups == nil {
		return &Unknown{}
	}

	names := strings.TrimSpace(groups["names"])
	parts := strings.Split(strings.TrimSpace(groups["names"]), ":")
	params := []string{}
	if names[len(names)-1] == ':' {
		params = parts[:len(parts)-1]
	}

	return &ObjCMethodDecl{
		Addr: ParseAddress(groups["address"]),
		Pos:  NewPositionFromString(groups["position"]),
		//Position2:    strings.TrimSpace(groups["position2"]),
		Implicit:    len(groups["implicit"]) > 0,
		ClassMethod: groups["methodtype"] == " +",
		Name:        parts[0],
		Parameters:  params,
		Type:        strings.TrimSpace(groups["type"]),
		//Type2:        strings.TrimSpace(groups["type2"]),
		Attr:       strings.TrimSpace(groups["attr"]),
		ChildNodes: []Node{},
	}
}

// AddChild adds a new child node. Child nodes can then be accessed with the
// Children attribute.
func (n *ObjCMethodDecl) AddChild(node Node) {
	n.ChildNodes = append(n.ChildNodes, node)
}

// Address returns the numeric address of the node. See the documentation for
// the Address type for more information.
func (n *ObjCMethodDecl) Address() Address {
	return n.Addr
}

// Children returns the child nodes. If this node does not have any children or
// this node does not support children it will always return an empty slice.
func (n *ObjCMethodDecl) Children() []Node {
	return n.ChildNodes
}

// Position returns the position in the original source code.
func (n *ObjCMethodDecl) Position() Position {
	return n.Pos
}
