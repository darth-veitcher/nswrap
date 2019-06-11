package ast

import (
	"testing"
)

func TestGotoStmt(t *testing.T) {
	nodes := map[string]testNode{
		`0x7fb9cc1994d8 <line:18893:9, col:14> 'end_getDigits' 0x7fb9cc199490`: testNode{&GotoStmt{
			Addr:       0x7fb9cc1994d8,
			Pos:        NewPositionFromString("line:18893:9, col:14"),
			Name:       "end_getDigits",
			Position2:  "0x7fb9cc199490",
			ChildNodes: []Node{},
		},
			0x7fb9cc1994d8,
			NewPositionFromString("line:18893:9, col:14"),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
