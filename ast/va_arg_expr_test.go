package ast

import (
	"testing"
)

func TestVAArgExpr(t *testing.T) {
	nodes := map[string]testNode{
		`0x7ff7d314bca8 <col:6, col:31> 'int *'`:
		testNode{&VAArgExpr{
			Addr:       0x7ff7d314bca8,
			Pos:        NewPositionFromString("col:6, col:31"),
			Type:       "int *",
			ChildNodes: []Node{},
		},
		0x7ff7d314bca8,
		NewPositionFromString("col:6, col:31"),
		[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
