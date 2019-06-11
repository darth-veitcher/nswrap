package ast

import (
	"testing"
)

func TestCompoundLiteralExpr(t *testing.T) {
	nodes := map[string]testNode{
		`0x5575acce81f0 <col:21, col:40> 'struct node':'struct node' lvalue`: testNode{&CompoundLiteralExpr{
			Addr:       0x5575acce81f0,
			Pos:        NewPositionFromString("col:21, col:40"),
			Type1:      "struct node",
			Type2:      "struct node",
			ChildNodes: []Node{},
		},
			0x5575acce81f0,
			NewPositionFromString("col:21, col:40"),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
