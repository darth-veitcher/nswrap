package ast

// ObjCObjectType is an Objective-C Object type
type ObjCObjectType struct {
	Addr       Address
	Type       string
	ChildNodes []Node
}

func parseObjCObjectType(line string) Node {
	groups := groupsFromRegex(
		"'(?P<type>.*)'",
		line,
	)
	if groups == nil {
		return &Unknown{}
	}

	return &ObjCObjectType{
		Addr:       ParseAddress(groups["address"]),
		Type:       groups["type"],
		ChildNodes: []Node{},
	}
}

// AddChild adds a new child node. Child nodes can then be accessed with the
// Children attribute.
func (n *ObjCObjectType) AddChild(node Node) {
	n.ChildNodes = append(n.ChildNodes, node)
}

// Address returns the numeric address of the node. See the documentation for
// the Address type for more information.
func (n *ObjCObjectType) Address() Address {
	return n.Addr
}

// Children returns the child nodes. If this node does not have any children or
// this node does not support children it will always return an empty slice.
func (n *ObjCObjectType) Children() []Node {
	return n.ChildNodes
}

// Position returns the position in the original source code.
func (n *ObjCObjectType) Position() Position {
	return Position{}
}
