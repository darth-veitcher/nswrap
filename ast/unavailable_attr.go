package ast

import (
	"strings"
)

// UnavailableAttr is node represents an unavailable attribute.
type UnavailableAttr struct {
	Addr         Address
	Pos          Position
	Position2    string
	Content      string
	ChildNodes   []Node
}

func parseUnavailableAttr(line string) Node {
	groups := groupsFromRegex(
		`(?:prev (?P<prev>0x[0-9a-f]+) )?
		<(?P<position><invalid sloc>|.*?)>
		(?P<position2> <invalid sloc>| col:\d+| line:\d+:\d+)?
		(?P<content>.*)`,
		line,
	)
	if groups == nil {
		return &Unknown{}
	}

	return &UnavailableAttr{
		Addr:         ParseAddress(groups["address"]),
		Pos:          NewPositionFromString(groups["position"]),
		//Position2:    strings.TrimSpace(groups["position2"]),
		Content:      strings.TrimSpace(groups["content"]),
		ChildNodes:   []Node{},
	}
}

// AddChild adds a new child node. Child nodes can then be accessed with the
// Children attribute.
func (n *UnavailableAttr) AddChild(node Node) {
	n.ChildNodes = append(n.ChildNodes, node)
}

// Address returns the numeric address of the node. See the documentation for
// the Address type for more information.
func (n *UnavailableAttr) Address() Address {
	return n.Addr
}

// Children returns the child nodes. If this node does not have any children or
// this node does not support children it will always return an empty slice.
func (n *UnavailableAttr) Children() []Node {
	return n.ChildNodes
}

// Position returns the position in the original source code.
func (n *UnavailableAttr) Position() Position {
	return n.Pos
}
