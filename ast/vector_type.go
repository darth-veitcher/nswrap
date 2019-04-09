package ast

import (
	"github.com/elliotchance/c2go/util"
)

// VectorType is vector type
type VectorType struct {
	Addr       Address
	Type       string
	Length     int
	ChildNodes []Node
}

func parseVectorType(line string) *VectorType {
	groups := groupsFromRegex(
		`'(?P<type>.*)'
		 (?P<length>[\d]+)`,
		line,
	)

	return &VectorType{
		Addr:       ParseAddress(groups["address"]),
		Type:       groups["type"],
		Length:     util.Atoi(groups["length"]),
		ChildNodes: []Node{},
	}
}

// AddChild adds a new child node. Child nodes can then be accessed with the
// Children attribute.
func (n *VectorType) AddChild(node Node) {
	n.ChildNodes = append(n.ChildNodes, node)
}

// Address returns the numeric address of the node. See the documentation for
// the Address type for more information.
func (n *VectorType) Address() Address {
	return n.Addr
}

// Children returns the child nodes. If this node does not have any children or
// this node does not support children it will always return an empty slice.
func (n *VectorType) Children() []Node {
	return n.ChildNodes
}

// Position returns the position in the original source code.
func (n *VectorType) Position() Position {
	return Position{}
}
