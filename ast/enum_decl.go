package ast

import (
	"strings"
)

// EnumDecl is node represents a enum declaration.
type EnumDecl struct {
	Addr       Address
	Pos        Position
	Position2  string
	Name       string
	Type       string
	Type2      string
	ChildNodes []Node
}

func parseEnumDecl(line string) Node {
	groups := groupsFromRegex(
		`(?:prev (?P<prev>0x[0-9a-f]+) )?
		<(?P<position>.*)>
		(?P<position2> .+:\d+)?
		(?P<name> \w+)?
		(?P<type> '.*?')?
		(?P<type2>:'.*')?`,
		line,
	)
	if groups == nil {
		return &Unknown{}
	}

	/*type2 := groups["type2"]
	if type2 != "" {
		type2 = type2[2 : len(type2)-1]
	}*/

	return &EnumDecl{
		Addr:      ParseAddress(groups["address"]),
		Pos:       NewPositionFromString(groups["position"]),
		Position2: groups["position2"],
		Name:      strings.TrimSpace(groups["name"]),
		Type:      removeQuotes(groups["type"]),
		//Type2:      type2,
		ChildNodes: []Node{},
	}
}

// AddChild adds a new child node. Child nodes can then be accessed with the
// Children attribute.
func (n *EnumDecl) AddChild(node Node) {
	n.ChildNodes = append(n.ChildNodes, node)
}

// Address returns the numeric address of the node. See the documentation for
// the Address type for more information.
func (n *EnumDecl) Address() Address {
	return n.Addr
}

// Children returns the child nodes. If this node does not have any children or
// this node does not support children it will always return an empty slice.
func (n *EnumDecl) Children() []Node {
	return n.ChildNodes
}

// Position returns the position in the original source code.
func (n *EnumDecl) Position() Position {
	return n.Pos
}
