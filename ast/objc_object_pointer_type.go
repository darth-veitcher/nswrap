package ast

// ObjCObjectPointerType is an Objective-C Object Pointer type
type ObjCObjectPointerType struct {
	Addr       Address
	Type       string
	ChildNodes []Node
}

func parseObjCObjectPointerType(line string) *ObjCObjectPointerType {
	groups := groupsFromRegex(
		"'(?P<type>.*)'",
		line,
	)

	return &ObjCObjectPointerType{
		Addr:       ParseAddress(groups["address"]),
		Type:       groups["type"],
		ChildNodes: []Node{},
	}
}

// AddChild adds a new child node. Child nodes can then be accessed with the
// Children attribute.
func (n *ObjCObjectPointerType) AddChild(node Node) {
	n.ChildNodes = append(n.ChildNodes, node)
}

// Address returns the numeric address of the node. See the documentation for
// the Address type for more information.
func (n *ObjCObjectPointerType) Address() Address {
	return n.Addr
}

// Children returns the child nodes. If this node does not have any children or
// this node does not support children it will always return an empty slice.
func (n *ObjCObjectPointerType) Children() []Node {
	return n.ChildNodes
}

// Position returns the position in the original source code.
func (n *ObjCObjectPointerType) Position() Position {
	return Position{}
}
