package ast

import (
	"testing"
)

func TestBreakStmt(t *testing.T) {
	nodes := map[string]testNode{
		`0x7fca2d8070e0 <col:11, col:23>`:
		testNode{&BreakStmt{
			Addr:       0x7fca2d8070e0,
			Pos:        NewPositionFromString("col:11, col:23"),
			ChildNodes: []Node{},
		},
		0x7fca2d8070e0,
		NewPositionFromString("col:11, col:23"),
		[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
