package ast

// NSReturnsRetainedAttr
type NSReturnsRetainedAttr struct {
	Addr       Address
	Pos        Position
	ChildNodes []Node
}

func parseNSReturnsRetainedAttr(line string) Node {
	groups := groupsFromRegex(
		"<(?P<position>.*)>",
		line,
	)
	if groups == nil {
		return &Unknown{}
	}

	return &NSReturnsRetainedAttr{
		Addr:       ParseAddress(groups["address"]),
		Pos:        NewPositionFromString(groups["position"]),
		ChildNodes: []Node{},
	}
}

// AddChild adds a new child node. Child nodes can then be accessed with the
// Children attribute.
func (n *NSReturnsRetainedAttr) AddChild(node Node) {
	n.ChildNodes = append(n.ChildNodes, node)
}

// Address returns the numeric address of the node. See the documentation for
// the Address type for more information.
func (n *NSReturnsRetainedAttr) Address() Address {
	return n.Addr
}

// Children returns the child nodes. If this node does not have any children or
// this node does not support children it will always return an empty slice.
func (n *NSReturnsRetainedAttr) Children() []Node {
	return n.ChildNodes
}

// Position returns the position in the original source code.
func (n *NSReturnsRetainedAttr) Position() Position {
	return n.Pos
}
