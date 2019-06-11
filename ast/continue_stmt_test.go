package ast

import (
	"testing"
)

func TestContinueStmt(t *testing.T) {
	nodes := map[string]testNode{
		`0x1e044e0 <col:20>`: testNode{&ContinueStmt{
			Addr:       0x1e044e0,
			Pos:        NewPositionFromString("col:20"),
			ChildNodes: []Node{},
		},
			0x1e044e0,
			NewPositionFromString("col:20"),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
