package ast

// ObjCBridgeAttr
type ObjCBridgeAttr struct {
	Addr       Address
	Pos        Position
	Content    string
	ChildNodes []Node
}

func parseObjCBridgeAttr(line string) Node {
	groups := groupsFromRegex(
		"<(?P<position>.*)>(?P<content>.*)",
		line,
	)
        if groups == nil {
                return &Unknown{}
        }

	return &ObjCBridgeAttr{
		Addr:       ParseAddress(groups["address"]),
		Pos:        NewPositionFromString(groups["position"]),
		Content:    groups["content"],
		ChildNodes: []Node{},
	}
}

// AddChild adds a new child node. Child nodes can then be accessed with the
// Children attribute.
func (n *ObjCBridgeAttr) AddChild(node Node) {
	n.ChildNodes = append(n.ChildNodes, node)
}

// Address returns the numeric address of the node. See the documentation for
// the Address type for more information.
func (n *ObjCBridgeAttr) Address() Address {
	return n.Addr
}

// Children returns the child nodes. If this node does not have any children or
// this node does not support children it will always return an empty slice.
func (n *ObjCBridgeAttr) Children() []Node {
	return n.ChildNodes
}

// Position returns the position in the original source code.
func (n *ObjCBridgeAttr) Position() Position {
	return n.Pos
}
