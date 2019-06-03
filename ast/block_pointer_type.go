package ast

// BlockPointerType is block pointer type
type BlockPointerType struct {
	Addr       Address
	Type       string
	ChildNodes []Node
}

func parseBlockPointerType(line string) Node {
	groups := groupsFromRegex(
		"'(?P<type>.*)'",
		line,
	)
	if groups == nil {
		return &Unknown{}
	}

	return &BlockPointerType{
		Addr:       ParseAddress(groups["address"]),
		Type:       groups["type"],
		ChildNodes: []Node{},
	}
}

// AddChild adds a new child node. Child nodes can then be accessed with the
// Children attribute.
func (n *BlockPointerType) AddChild(node Node) {
	n.ChildNodes = append(n.ChildNodes, node)
}

// Address returns the numeric address of the node. See the documentation for
// the Address type for more information.
func (n *BlockPointerType) Address() Address {
	return n.Addr
}

// Children returns the child nodes. If this node does not have any children or
// this node does not support children it will always return an empty slice.
func (n *BlockPointerType) Children() []Node {
	return n.ChildNodes
}

// Position returns the position in the original source code.
func (n *BlockPointerType) Position() Position {
	return Position{}
}
