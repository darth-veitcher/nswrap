package ast

// ObjCReturnsInnerPointerAttr
type ObjCReturnsInnerPointerAttr struct {
	Addr       Address
	Pos        Position
	Content    string
	ChildNodes []Node
}

func parseObjCReturnsInnerPointerAttr(line string) Node {
	groups := groupsFromRegex(
		"<(?P<position>.*)>(?P<content>.*)?",
		line,
	)
	if groups == nil {
		return &Unknown{}
	}

	return &ObjCReturnsInnerPointerAttr{
		Addr:       ParseAddress(groups["address"]),
		Pos:        NewPositionFromString(groups["position"]),
		Content:    groups["content"],
		ChildNodes: []Node{},
	}
}

// AddChild adds a new child node. Child nodes can then be accessed with the
// Children attribute.
func (n *ObjCReturnsInnerPointerAttr) AddChild(node Node) {
	n.ChildNodes = append(n.ChildNodes, node)
}

// Address returns the numeric address of the node. See the documentation for
// the Address type for more information.
func (n *ObjCReturnsInnerPointerAttr) Address() Address {
	return n.Addr
}

// Children returns the child nodes. If this node does not have any children or
// this node does not support children it will always return an empty slice.
func (n *ObjCReturnsInnerPointerAttr) Children() []Node {
	return n.ChildNodes
}

// Position returns the position in the original source code.
func (n *ObjCReturnsInnerPointerAttr) Position() Position {
	return n.Pos
}
