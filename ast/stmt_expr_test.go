package ast

import (
	"testing"
)

func TestStmtExpr(t *testing.T) {
	nodes := map[string]testNode{
		`0x7ff4f9100d28 <col:11, col:18> 'int'`:
		testNode{&StmtExpr{
			Addr:       0x7ff4f9100d28,
			Pos:        NewPositionFromString("col:11, col:18"),
			Type:       "int",
			ChildNodes: []Node{},
		},
		0x7ff4f9100d28,
		NewPositionFromString("col:11, col:18"),
		[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
