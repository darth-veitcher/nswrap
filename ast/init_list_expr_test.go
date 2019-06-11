package ast

import (
	"testing"
)

func TestInitListExpr(t *testing.T) {
	nodes := map[string]testNode{
		`0x7fbdd1906c20 <col:52, line:17160:1> 'const unsigned char [256]'`: testNode{&InitListExpr{
			Addr:       0x7fbdd1906c20,
			Pos:        NewPositionFromString("col:52, line:17160:1"),
			Type1:      "const unsigned char [256]",
			ChildNodes: []Node{},
		},
			0x7fbdd1906c20,
			NewPositionFromString("col:52, line:17160:1"),
			[]Node{},
		},
		`0x32017f0 <col:24, col:41> 'struct node [2]'`: testNode{&InitListExpr{
			Addr:       0x32017f0,
			Pos:        NewPositionFromString("col:24, col:41"),
			Type1:      "struct node [2]",
			ChildNodes: []Node{},
		},
			0x32017f0,
			NewPositionFromString("col:24, col:41"),
			[]Node{},
		},
		`0x3201840 <col:25, col:31> 'struct node':'struct node'`: testNode{&InitListExpr{
			Addr:       0x3201840,
			Pos:        NewPositionFromString("col:25, col:31"),
			Type1:      "struct node",
			Type2:      "struct node",
			ChildNodes: []Node{},
		},
			0x3201840,
			NewPositionFromString("col:25, col:31"),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
