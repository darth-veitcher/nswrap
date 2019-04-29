package ast

// Variadic node indicates an ellipsis (...) in the Clang AST.
type Variadic struct {
	ChildNodes []Node
}

func parseVariadic(line string) *Variadic {
	return &Variadic{
		ChildNodes: []Node{},
	}
}

// AddChild adds a new child node. Child nodes can then be accessed with the
// Children attribute.
func (n *Variadic) AddChild(node Node) {
	n.ChildNodes = append(n.ChildNodes, node)
}

// Address returns the numeric address of the node. See the documentation for
// the Address type for more information.
func (n *Variadic) Address() Address {
	return Address(0)
}

// Children returns the child nodes. If this node does not have any children or
// this node does not support children it will always return an empty slice.
func (n *Variadic) Children() []Node {
	return n.ChildNodes
}

// Position returns the position in the original source code.
func (n *Variadic) Position() Position {
	return Position{}
}
