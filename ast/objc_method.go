package ast

// ObjCMethod is an Objective-C method
type ObjCMethod struct {
	Addr       Address
	Name       string
	ChildNodes []Node
}

func parseObjCMethod(line string) *ObjCMethod {
	groups := groupsFromRegex(
		"'(?P<name>.*)'",
		line,
	)

	return &ObjCMethod{
		Addr:       ParseAddress(groups["address"]),
		Name:       groups["name"],
		ChildNodes: []Node{},
	}
}

// AddChild adds a new child node. Child nodes can then be accessed with the
// Children attribute.
func (n *ObjCMethod) AddChild(node Node) {
	n.ChildNodes = append(n.ChildNodes, node)
}

// Address returns the numeric address of the node. See the documentation for
// the Address type for more information.
func (n *ObjCMethod) Address() Address {
	return n.Addr
}

// Children returns the child nodes. If this node does not have any children or
// this node does not support children it will always return an empty slice.
func (n *ObjCMethod) Children() []Node {
	return n.ChildNodes
}

// Position returns the position in the original source code.
func (n *ObjCMethod) Position() Position {
	return Position{}
}
