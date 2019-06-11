package ast

import (
	"testing"
)

func TestForStmt(t *testing.T) {
	nodes := map[string]testNode{
		`0x7f961e018848 <line:9:4, line:10:70>`: testNode{&ForStmt{
			Addr:       0x7f961e018848,
			Pos:        NewPositionFromString("line:9:4, line:10:70"),
			ChildNodes: []Node{},
		},
			0x7f961e018848,
			NewPositionFromString("line:9:4, line:10:70"),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
