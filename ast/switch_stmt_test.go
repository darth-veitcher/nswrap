package ast

import (
	"testing"
)

func TestSwitchStmt(t *testing.T) {
	nodes := map[string]testNode{
		`0x7fbca3894638 <line:9:5, line:20:5>`:
		testNode{&SwitchStmt{
			Addr:       0x7fbca3894638,
			Pos:        NewPositionFromString("line:9:5, line:20:5"),
			ChildNodes: []Node{},
		},
		0x7fbca3894638,
		NewPositionFromString("line:9:5, line:20:5"),
		[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
