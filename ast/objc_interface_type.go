package ast

// ObjCInterfaceType is an Objective-C interface type
type ObjCInterfaceType struct {
	Addr       Address
	Type       string
	ChildNodes []Node
}

func parseObjCInterfaceType(line string) *ObjCInterfaceType {
	groups := groupsFromRegex(
		"'(?P<type>.*)'",
		line,
	)

	return &ObjCInterfaceType{
		Addr:       ParseAddress(groups["address"]),
		Type:       groups["type"],
		ChildNodes: []Node{},
	}
}

// AddChild adds a new child node. Child nodes can then be accessed with the
// Children attribute.
func (n *ObjCInterfaceType) AddChild(node Node) {
	n.ChildNodes = append(n.ChildNodes, node)
}

// Address returns the numeric address of the node. See the documentation for
// the Address type for more information.
func (n *ObjCInterfaceType) Address() Address {
	return n.Addr
}

// Children returns the child nodes. If this node does not have any children or
// this node does not support children it will always return an empty slice.
func (n *ObjCInterfaceType) Children() []Node {
	return n.ChildNodes
}

// Position returns the position in the original source code.
func (n *ObjCInterfaceType) Position() Position {
	return Position{}
}
