package ast

import (
	"testing"
)

func TestDefaultStmt(t *testing.T) {
	nodes := map[string]testNode{
		`0x7f951308bfb0 <line:17:5, line:18:34>`: testNode{&DefaultStmt{
			Addr:       0x7f951308bfb0,
			Pos:        NewPositionFromString("line:17:5, line:18:34"),
			ChildNodes: []Node{},
		},
			0x7f951308bfb0,
			NewPositionFromString("line:17:5, line:18:34"),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
