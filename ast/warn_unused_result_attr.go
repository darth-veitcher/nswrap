package ast

// WarnUnusedResultAttr is a type of attribute that is optionally attached to a variable
// or struct field definition.
type WarnUnusedResultAttr struct {
	Addr       Address
	Pos        Position
	ChildNodes []Node
}

func parseWarnUnusedResultAttr(line string) Node {
	groups := groupsFromRegex(`<(?P<position>.*)>( warn_unused_result)?`, line)
	if groups == nil {
		return &Unknown{}
	}

	return &WarnUnusedResultAttr{
		Addr:       ParseAddress(groups["address"]),
		Pos:        NewPositionFromString(groups["position"]),
		ChildNodes: []Node{},
	}
}

// AddChild adds a new child node. Child nodes can then be accessed with the
// Children attribute.
func (n *WarnUnusedResultAttr) AddChild(node Node) {
	n.ChildNodes = append(n.ChildNodes, node)
}

// Address returns the numeric address of the node. See the documentation for
// the Address type for more information.
func (n *WarnUnusedResultAttr) Address() Address {
	return n.Addr
}

// Children returns the child nodes. If this node does not have any children or
// this node does not support children it will always return an empty slice.
func (n *WarnUnusedResultAttr) Children() []Node {
	return n.ChildNodes
}

// Position returns the position in the original source code.
func (n *WarnUnusedResultAttr) Position() Position {
	return n.Pos
}
