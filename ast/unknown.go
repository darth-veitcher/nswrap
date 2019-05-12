package ast

import (
	//"strings"
)

// Unknown is node represents an unknown node.
type Unknown struct {
	Name         string
	Addr         Address
	Pos          Position
	Position2    string
	Content      string
	ChildNodes   []Node
}

func parseUnknown(name, line string) *Unknown {
	groups := groupsFromRegex(
                `(?:prev (?P<prev>0x[0-9a-f]+) )?
		<(?P<position>.*<scratch space>.*?|.*<built-in>.*?|.*<invalid sloc>|.*?)>
		(?P<position2> <invalid sloc>| col:\d+| line:\d+:\d+)?
		(?P<content>.*)`,
		line,
	)
/*
		(?P<position2> <invalid sloc>| col:\d+| line:\d+:\d+)?
*/

	return &Unknown{
		Name:         name,
		Addr:         ParseAddress(groups["address"]),
		Pos:          NewPositionFromString(groups["position"]),
		//Position2:    strings.TrimSpace(groups["position2"]),
		Content:      groups["content"],
		ChildNodes:   []Node{},
	}
}

// AddChild adds a new child node. Child nodes can then be accessed with the
// Children attribute.
func (n *Unknown) AddChild(node Node) {
	n.ChildNodes = append(n.ChildNodes, node)
}

// Address returns the numeric address of the node. See the documentation for
// the Address type for more information.
func (n *Unknown) Address() Address {
	return n.Addr
}

// Children returns the child nodes. If this node does not have any children or
// this node does not support children it will always return an empty slice.
func (n *Unknown) Children() []Node {
	return n.ChildNodes
}

// Position returns the position in the original source code.
func (n *Unknown) Position() Position {
	return n.Pos
}
