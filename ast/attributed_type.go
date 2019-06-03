package ast

// AttributedType is attributed type
type AttributedType struct {
	Addr       Address
	Type       string
	Sugar      bool
	ChildNodes []Node
}

func parseAttributedType(line string) Node {
	groups := groupsFromRegex(
		`'(?P<type>.*)'
		(?P<sugar> sugar)?`,
		line,
	)
	if groups == nil {
		return &Unknown{}
	}

	return &AttributedType{
		Addr:       ParseAddress(groups["address"]),
		Type:       groups["type"],
		Sugar:      len(groups["sugar"])>0,
		ChildNodes: []Node{},
	}
}

// AddChild adds a new child node. Child nodes can then be accessed with the
// Children attribute.
func (n *AttributedType) AddChild(node Node) {
	n.ChildNodes = append(n.ChildNodes, node)
}

// Address returns the numeric address of the node. See the documentation for
// the Address type for more information.
func (n *AttributedType) Address() Address {
	return n.Addr
}

// Children returns the child nodes. If this node does not have any children or
// this node does not support children it will always return an empty slice.
func (n *AttributedType) Children() []Node {
	return n.ChildNodes
}

// Position returns the position in the original source code.
func (n *AttributedType) Position() Position {
	return Position{}
}
