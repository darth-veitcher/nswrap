package types

import (
	"fmt"
)

var (
	Debug bool = false
)

func dbg(f string, xs ...interface{}) {
	if Debug {
		fmt.Printf(f,xs...)
	}
}

func (n *Node) HasFunc() bool {
	if n == nil {
		return false
	}
	for _,c := range n.Children {
		if c.Kind == "Function" || c.HasFunc() {
			return true
		}
	}
	return false
}

func Parse(s string) (*Node, error) {
	s2, n := TypeName(s,NewNode("AST"))
	//fmt.Printf("%p Parsed %s\n",n,s)
	if s2 != "" {
		return n,fmt.Errorf("Parse failed or incomplete. Remainder: %s",s2)
	}
	return n, nil
}

//Evaluate a node to determine if it is a pointer or array
func (n *Node) isAbstract(k string) bool {
	if n.stripAbstract(k) == nil {
		return false
	}
	return true
}

//Strip one level of pointer or array indirection from a node
func (n *Node) stripAbstract(k string) *Node {
	if n == nil {
		return nil
	}
	i := len(n.Children) - 1
	if i < 1 {
		return nil
	}
	ret := NewNode(n.Kind)
	cs := append([]*Node{},n.Children...)

	dbg("stripAbstract(): i = %d\n",i)
	//Scan backwords skipping TypeQualifier and NullableAnnotation tags
	for ;i > 0 &&
		(cs[i].Kind == "TypeQualifier" ||
		 cs[i].Kind == "NullableAnnotation") ; i-- { }

	if cs[i].Kind == k {
		dbg("stripAbstract(): last node is %s\n",k)
		ret.Children = cs[:i]
		return ret
	}
	if i > 1 && cs[i-1].Kind == "Parenthesized" {
		j := len(cs[i-1].Children) - 1
		for ;j > 0 &&
			(cs[i-1].Children[j].Kind == "TypeQualifier" ||
			 cs[i-1].Children[j].Kind == "NullableAnnotation");
		     j-- { }
		if cs[i-1].Children[j].Kind != k {
			return nil
		}
		if j == 0 { // strip Parenthesized tag
			cs[i-1] = cs[i]
			ret.Children = cs[:i]
			return ret
		}
		// strip last child from Parenthesized tag
		cs[i-1].Children = cs[i-1].Children[:j]
		ret.Children = cs
		return ret
	}
	return nil
}

//PointsTo, when called on a pointer node returns a node describing the type
//pointed to. Otherwise returns nil when called on non-pointer types.
func (n *Node) PointsTo() *Node {
	dbg("PointsTo()\n")
	return n.stripAbstract("Pointer")
}

//IsPointer returns true if the node is a pointer
func (n *Node) IsPointer() bool {
	if pt := n.PointsTo(); pt != nil {
		return true
	}
	return n.IsInstancetype() || n.IsId()
}

//ArrayOf, when called on an array node returns a node describing the type
//of the elements of the array. Otherwise returns nil when called on
//non-array types.
func (n *Node) ArrayOf() *Node {
	dbg("ArrayOf()\n")
	return n.stripAbstract("Array")
}

//IsArray returns true if the node is an array
func (n *Node) IsArray() bool {
	return n.ArrayOf() != nil
}

func (n *Node) IsStruct() bool {
	if n == nil || len(n.Children) < 1 {
		return false
	}
	return n.Children[0].Kind == "Struct"
}

func (n *Node) IsFunction() bool {
	if n == nil || len(n.Children) < 1 {
		return false
	}
	if pt := n.PointsTo(); pt != nil {
		return false
	}
	return n.Children[len(n.Children)-1].Kind == "Function"
}

func (n *Node) ReturnType() *Node {
	if !n.IsFunction() {
		return nil
	}
	ret := NewNode(n.Kind)
	ret.Children = n.Children[:len(n.Children)-1]
	return ret
}

func (n *Node) IsId() bool {
	if n == nil || len(n.Children) < 1 {
		return false
	}
	return !n.IsFunction() &&
		n.Children[0].Kind == "TypedefName" &&
		n.Children[0].Content == "id"
}

func (n *Node) IsInstancetype() bool {
	if n == nil || len(n.Children) < 1 {
		return false
	}
	return n.Children[0].Kind == "TypedefName" &&
		n.Children[0].Content == "instancetype"
}

//BaseType strips off all layers of pointer indirection
func (n *Node) BaseType() *Node {
	if n == nil {
		return nil
	}
	if n2 := n.PointsTo(); n2 == nil {
		return n
	} else {
		return n2.BaseType()
	}
}
