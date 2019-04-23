package types

import (
	"fmt"
	"os"
	"strings"
)

// Type definition and basic functions for Nodes

type Node struct {
	Kind, Content string
	Children []*Node
}

func NewNode(k string,cs ...string) *Node {
	c := ""
	if len(cs) > 0 {
		c = cs[0]
	}
	ret := &Node{ Kind: k, Content: c, Children: []*Node{} }
	dbg("NewNode(%p) %s\n",ret,ret.Kind)
	return ret
}

func (n *Node) String(ls ...int) string {
	if n == nil {
		return ""
	}
	var ret strings.Builder
	level := 0
	if len(ls) > 0 {
		level = ls[0]
		if level > 100 {
			fmt.Println("(*Node)String(): Recursion too deep")
			os.Exit(-1)
		}
	}
	prefix := strings.Repeat("-",level)
	ret.WriteString(fmt.Sprintf("%s<%s> '%s'\n",prefix, n.Kind, n.Content))
	for _,c := range n.Children {
		ret.WriteString(c.String(level+1))
	}
	return ret.String()
}

func (n *Node) AddChild(c *Node) *Node {
	erp := func(s string) {
		dbg("(%p)AddChild(%p): %s\n",n,c,s)
		os.Exit(-1)
	}
	if n == nil { erp("Called on nil node") }
	if c == nil { erp("Child is nil") }
	if n == c { erp("Node cannot be its own child") }

	// Skip literals
	if c.Kind == "Lit" { return n }

	// Do we already have this child? (FIXME: Not needed?)
	for _,d := range n.Children {
		if c == d {
			return n
		}
	}
	dbg("(%p)AddChild(%p)\n",n,c)
	n.Children = append(n.Children,c)
	return n
}

func (n *Node) CtypeSimplified() string {
	ignore := map[string]bool{
		"NullableAnnotation": true,
		"TypeQualifier": true,
		"GenericList": true,
	}
	return n._Ctype(ignore)
}

func (n *Node) Ctype() string {
	return n._Ctype(map[string]bool{})
}

func (n *Node) _Ctype(ignore map[string]bool) string {
	if n == nil || ignore[n.Kind] {
		return ""
	}
	var ret strings.Builder
	childStrings := func(n *Node) []string {
		if n == nil { return []string{} }
		ret := []string{}
		for _,c := range n.Children {
			if x := c._Ctype(ignore); x != "" {
				ret = append(ret, x)
			}
		}
		return ret
	}
	switch n.Kind {
	case "Parenthesized":
		ret.WriteString("(" + strings.Join(childStrings(n)," ") + ")")
	case "Function":
		ret.WriteString("(" + strings.Join(childStrings(n),", ") + ")")
	case "GenericList":
		ret.WriteString("<" + strings.Join(childStrings(n),", ") + ">")
	case "Array":
		ret.WriteString("[" + strings.Join(childStrings(n)," ") + "]")
	default:
		ret.WriteString(n.Content)
		cc := strings.Join(childStrings(n)," ")
		if n.Content != "" && cc != "" {
			ret.WriteString(" ")
		}
		ret.WriteString(cc)
	}
	s := ret.String()
	s = strings.ReplaceAll(s," *","*")
	s = strings.ReplaceAll(s," [","[")
	s = strings.ReplaceAll(s,") (",")(")
	return s
}

