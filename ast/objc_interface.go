package ast

// ObjCInterface is an Objective-C interface
type ObjCInterface struct {
	Addr       Address
	Name       string
	Super      bool
	ChildNodes []Node
}

func parseObjCInterface(line string, super bool) Node {
	groups := groupsFromRegex(
		"'(?P<name>.*)'",
		line,
	)
	if groups == nil {
		return &Unknown{}
	}

	return &ObjCInterface{
		Addr:       ParseAddress(groups["address"]),
		Name:       groups["name"],
		Super:      super,
		ChildNodes: []Node{},
	}
}

// AddChild adds a new child node. Child nodes can then be accessed with the
// Children attribute.
func (n *ObjCInterface) AddChild(node Node) {
	n.ChildNodes = append(n.ChildNodes, node)
}

// Address returns the numeric address of the node. See the documentation for
// the Address type for more information.
func (n *ObjCInterface) Address() Address {
	return n.Addr
}

// Children returns the child nodes. If this node does not have any children or
// this node does not support children it will always return an empty slice.
func (n *ObjCInterface) Children() []Node {
	return n.ChildNodes
}

// Position returns the position in the original source code.
func (n *ObjCInterface) Position() Position {
	return Position{}
}
