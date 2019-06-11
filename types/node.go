package types

import (
	"fmt"
	"os"
	"strings"
)

// Type definition and basic functions for Nodes

type Node struct {
	Kind, Content string
	Children      []*Node
}

//NewNode returns a new node of kind k with an optional content string as its
//second parameter.
func NewNode(k string, cs ...string) *Node {
	c := ""
	if len(cs) > 0 {
		c = cs[0]
	}
	ret := &Node{Kind: k, Content: c, Children: []*Node{}}
	dbg("NewNode(%p) %s\n", ret, ret.Kind)
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
			return "(*Node)String(): Recursion too deep"
		}
	}
	prefix := strings.Repeat("-", level)
	ret.WriteString(fmt.Sprintf("%s<%s> '%s'\n", prefix, n.Kind, n.Content))
	for _, c := range n.Children {
		ret.WriteString(c.String(level + 1))
	}
	return ret.String()
}

func (n *Node) AddChild(c *Node) *Node {
	erp := func(s string) {
		dbg("(%p)AddChild(%p): %s\n", n, c, s)
		os.Exit(-1)
	}
	if n == nil {
		erp("Called on nil node")
	}
	if c == nil {
		erp("Child is nil")
	}
	if n == c {
		erp("Node cannot be its own child")
	}

	// Skip literals
	if c.Kind == "Lit" {
		return n
	}

	for _, d := range n.Children {
		if c == d {
			return n
		}
	}
	dbg("(%p)AddChild(%p)\n", n, c)
	n.Children = append(n.Children, c)
	return n
}

//returns true if anything gets renamed
func (n *Node) renameTypedefs(a, b string) (ret bool) {
	ret = false
	if n == nil {
		return
	}
	for i, c := range n.Children {
		if c.Kind == "TypedefName" && c.Content == a {
			ret = true
			n.Children[i] = NewNode("TypedefName", b)
			n.Children[i].Children = c.Children
		}
		if len(c.Children) > 0 {
			ret2 := c.renameTypedefs(a, b)
			ret = ret || ret2
		}
	}
	return
}

func (n *Node) CTypeSimplified() string {
	ignore := map[string]bool{
		"NullableAnnotation": true,
		"KindQualifier":      true,
		"TypeQualifier":      true,
		"GenericList":        true,
	}
	return n._CType(ignore)
}

func (n *Node) CType() string {
	return n._CType(map[string]bool{})
}

func (n *Node) _CType(ignore map[string]bool) string {
	if n == nil || ignore[n.Kind] {
		return ""
	}
	var ret strings.Builder
	childStrings := func(n *Node) []string {
		ret := []string{}
		if n == nil {
			return ret
		}
		for _, c := range n.Children {
			if x := c._CType(ignore); x != "" {
				ret = append(ret, x)
			}
		}
		return ret
	}
	switch n.Kind {
	case "Parenthesized":
		ret.WriteString("(" + strings.Join(childStrings(n), " ") + ")")
	case "Function":
		ret.WriteString("(" + strings.Join(childStrings(n), ", ") + ")")
	case "GenericList":
		ret.WriteString("<" + strings.Join(childStrings(n), ", ") + ">")
	case "Array":
		ret.WriteString("[" + strings.Join(childStrings(n), " ") + "]")
	default:
		ret.WriteString(n.Content)
		cc := strings.Join(childStrings(n), " ")
		if n.Content != "" && cc != "" {
			ret.WriteString(" ")
		}
		ret.WriteString(cc)
	}
	s := ret.String()
	s = strings.ReplaceAll(s, " *", "*")
	s = strings.ReplaceAll(s, " [", "[")
	s = strings.ReplaceAll(s, ") (", ")(")
	return s
}

func (n *Node) Qualifiers() string {
	if n == nil {
		return ""
	}
	ret := []string{}
	for _, c := range n.Children {
		switch c.Kind {
		case "TypeQualifier":
			ret = append(ret, c.Content)
		}
	}
	return strings.Join(ret, " ")
}

func (n *Node) Annotations() string {
	if n == nil {
		return ""
	}
	ret := []string{}
	for _, c := range n.Children {
		switch c.Kind {
		case "NullableAnnotation":
			ret = append(ret, c.Content)
		}
	}
	return strings.Join(ret, " ")
}
