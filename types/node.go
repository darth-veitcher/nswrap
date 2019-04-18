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
	ret.WriteString(fmt.Sprintf("%s<%s> %p '%s'\n",prefix, n.Kind, n, n.Content))
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

