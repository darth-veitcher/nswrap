package ast

// NoDebugAttr
type NoDebugAttr struct {
	Addr       Address
	Pos        Position
	ChildNodes []Node
}

func parseNoDebugAttr(line string) Node {
	groups := groupsFromRegex(
		"<(?P<position>.*)>",
		line,
	)
	if groups == nil {
		return &Unknown{}
	}

	return &NoDebugAttr{
		Addr:       ParseAddress(groups["address"]),
		Pos:        NewPositionFromString(groups["position"]),
		ChildNodes: []Node{},
	}
}

// AddChild adds a new child node. Child nodes can then be accessed with the
// Children attribute.
func (n *NoDebugAttr) AddChild(node Node) {
	n.ChildNodes = append(n.ChildNodes, node)
}

// Address returns the numeric address of the node. See the documentation for
// the Address type for more information.
func (n *NoDebugAttr) Address() Address {
	return n.Addr
}

// Children returns the child nodes. If this node does not have any children or
// this node does not support children it will always return an empty slice.
func (n *NoDebugAttr) Children() []Node {
	return n.ChildNodes
}

// Position returns the position in the original source code.
func (n *NoDebugAttr) Position() Position {
	return n.Pos
}
