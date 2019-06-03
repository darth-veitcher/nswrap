package ast

// ObjCBoolLiteralExpr
type ObjCBoolLiteralExpr struct {
	Addr         Address
	Pos          Position
	Type         string
	Type2        string
	Attr         string
	ChildNodes   []Node
}

func parseObjCBoolLiteralExpr(line string) Node {
	groups := groupsFromRegex(
		`<(?P<position><invalid sloc>|.*?)>
		(?P<type> '.*?')?
		(?P<type2>:'.*?')?
		(?P<attr>.*)`,
		line,
	)
        if groups == nil {
                return &Unknown{}
        }

	return &ObjCBoolLiteralExpr{
		Addr:         ParseAddress(groups["address"]),
		Pos:          NewPositionFromString(groups["position"]),
		Type:         removeQuotes(groups["type"]),
		Type2:        groups["type2"],
		Attr:         groups["attr"],
		ChildNodes:   []Node{},
	}
}

// AddChild adds a new child node. Child nodes can then be accessed with the
// Children attribute.
func (n *ObjCBoolLiteralExpr) AddChild(node Node) {
	n.ChildNodes = append(n.ChildNodes, node)
}

// Address returns the numeric address of the node. See the documentation for
// the Address type for more information.
func (n *ObjCBoolLiteralExpr) Address() Address {
	return n.Addr
}

// Children returns the child nodes. If this node does not have any children or
// this node does not support children it will always return an empty slice.
func (n *ObjCBoolLiteralExpr) Children() []Node {
	return n.ChildNodes
}

// Position returns the position in the original source code.
func (n *ObjCBoolLiteralExpr) Position() Position {
	return n.Pos
}
