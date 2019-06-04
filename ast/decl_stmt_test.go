package ast

import (
	"testing"
)

func TestDeclStmt(t *testing.T) {
	nodes := map[string]testNode{
		`0x7fb791846e80 <line:11:4, col:31>`:
		testNode{&DeclStmt{
			Addr:       0x7fb791846e80,
			Pos:        NewPositionFromString("line:11:4, col:31"),
			ChildNodes: []Node{},
		},
		0x7fb791846e80,
		NewPositionFromString("line:11:4, col:31"),
		[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
