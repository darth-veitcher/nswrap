package ast

import (
	"strings"
)

// ObjCTypeParamDecl is node represents a parameter of variable declaration.
type ObjCTypeParamDecl struct {
	Addr         Address
	Pos          Position
	Position2    string
	Name         string
	Type         string
	Type2        string
	IsReferenced bool
	IsCovariant  bool
	IsBounded    bool
	ChildNodes   []Node
}

func parseObjCTypeParamDecl(line string) Node {
	groups := groupsFromRegex(
		`<(?P<position>.*)>
		(?P<position2> [^ ]+:[\d:]+)?
		(?P<referenced> referenced)?
		(?P<name> \w+)?
		(?P<covariant> covariant)?
		(?P<bounded> bounded)?
		 '(?P<type>.*?)'
		(?P<type2>:'.*?')?
		`,
		line,
	)
	if groups == nil {
		return &Unknown{}
	}

	/*type2 := groups["type2"]
	if type2 != "" {
		type2 = type2[2 : len(type2)-1]
	}*/

	if strings.Index(groups["position"], "<invalid sloc>") > -1 {
		groups["position"] = "<invalid sloc>"
		groups["position2"] = "<invalid sloc>"
	}

	return &ObjCTypeParamDecl{
		Addr:         ParseAddress(groups["address"]),
	        Pos:          NewPositionFromString(groups["position"]),
		//Position2:    strings.TrimSpace(groups["position2"]),
		Name:         strings.TrimSpace(groups["name"]),
		Type:         groups["type"],
		//Type2:        type2,
		IsReferenced: len(groups["referenced"]) > 0,
		IsCovariant:  len(groups["covariant"]) > 0,
		IsBounded  :  len(groups["bounded"]) > 0,
		ChildNodes:   []Node{},
	}
}

// AddChild adds a new child node. Child nodes can then be accessed with the
// Children attribute.
func (n *ObjCTypeParamDecl) AddChild(node Node) {
	n.ChildNodes = append(n.ChildNodes, node)
}

// Address returns the numeric address of the node. See the documentation for
// the Address type for more information.
func (n *ObjCTypeParamDecl) Address() Address {
	return n.Addr
}

// Children returns the child nodes. If this node does not have any children or
// this node does not support children it will always return an empty slice.
func (n *ObjCTypeParamDecl) Children() []Node {
	return n.ChildNodes
}

// Position returns the position in the original source code.
func (n *ObjCTypeParamDecl) Position() Position {
	return n.Pos
}
