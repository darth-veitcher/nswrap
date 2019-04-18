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

func Parse(s string) *Node {
	_, n2 := TypeName(s,NewNode("AST"))
	return n2
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
	cs := n.Children[:]

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

//ArrayOf, when called on an array node returns a node describing the type
//of the elements of the array. Otherwise returns nil when called on
//non-array types.
func (n *Node) ArrayOf() *Node {
	dbg("ArrayOf()\n")
	return n.stripAbstract("Array")
}

