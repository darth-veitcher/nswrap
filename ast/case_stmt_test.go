package ast

import (
	"testing"
)

func TestCaseStmt(t *testing.T) {
	nodes := map[string]testNode{
		`0x7fc8b5094688 <line:11:5, line:12:21>`:
		testNode{&CaseStmt{
			Addr:       0x7fc8b5094688,
			Pos:        NewPositionFromString("line:11:5, line:12:21"),
			ChildNodes: []Node{},
		},
		0x7fc8b5094688,
		NewPositionFromString("line:11:5, line:12:21"),
		[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
