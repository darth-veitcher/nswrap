package ast

import (
	"testing"
)

func TestEmptyDecl(t *testing.T) {
	nodes := map[string]testNode{
		`0x480bec8 <col:13> col:13`:
		testNode{&EmptyDecl{
			Addr:       0x480bec8,
			Pos:        NewPositionFromString("col:13"),
			Position2:  NewPositionFromString("col:13"),
			ChildNodes: []Node{},
		},
		0x480bec8,
		NewPositionFromString("col:13"),
		[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
