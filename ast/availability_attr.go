package ast

// AvailabilityAttr is a type of attribute that is optionally attached to a variable
// or struct field definition.
type AvailabilityAttr struct {
	Addr          Address
	Pos           Position
	OS            string
	Version       string
	Unknown1      string
	Unknown2      string
	IsUnavailable bool
	Message1      string
	Message2      string
	IsInherited   bool
	ChildNodes    []Node
}

func parseAvailabilityAttr(line string) Node {
	groups := groupsFromRegex(
		`<(?P<position>.*)>
		(?P<inherited> Inherited)?
		 (?P<os>\w+)
		 (?P<version>[\d_.]+)
		 (?P<unknown1>[\d_.]+)
		 (?P<unknown2>[\d_.]+)
		(?P<unavalable> Unavailable)?
		 "(?P<message1>.*?)"
		(?P<message2> ".*?")?`,
		line,
	)
	if groups == nil {
		return &Unknown{}
	}

	return &AvailabilityAttr{
		Addr:          ParseAddress(groups["address"]),
		Pos:           NewPositionFromString(groups["position"]),
		OS:            groups["os"],
		Version:       groups["version"],
		Unknown1:      groups["unknown1"],
		Unknown2:      groups["unknown2"],
		IsUnavailable: len(groups["unavalable"]) > 0,
		Message1:      removeQuotes(groups["message1"]),
		Message2:      removeQuotes(groups["message2"]),
		IsInherited:   len(groups["inherited"]) > 0,
		ChildNodes:    []Node{},
	}
}

// AddChild adds a new child node. Child nodes can then be accessed with the
// Children attribute.
func (n *AvailabilityAttr) AddChild(node Node) {
	n.ChildNodes = append(n.ChildNodes, node)
}

// Address returns the numeric address of the node. See the documentation for
// the Address type for more information.
func (n *AvailabilityAttr) Address() Address {
	return n.Addr
}

// Children returns the child nodes. If this node does not have any children or
// this node does not support children it will always return an empty slice.
func (n *AvailabilityAttr) Children() []Node {
	return n.ChildNodes
}

// Position returns the position in the original source code.
func (n *AvailabilityAttr) Position() Position {
	return n.Pos
}
