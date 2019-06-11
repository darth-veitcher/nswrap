package ast

import (
	"testing"
)

func TestReturnStmt(t *testing.T) {
	nodes := map[string]testNode{
		`0x7fbb7a8325e0 <line:13:4, col:11>`: testNode{&ReturnStmt{
			Addr:       0x7fbb7a8325e0,
			Pos:        NewPositionFromString("line:13:4, col:11"),
			ChildNodes: []Node{},
		},
			0x7fbb7a8325e0,
			NewPositionFromString("line:13:4, col:11"),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
