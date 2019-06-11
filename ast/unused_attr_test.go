package ast

import (
	"testing"
)

func TestUnusedAttr(t *testing.T) {
	nodes := map[string]testNode{
		`0x7fe3e01416d0 <col:47> unused`: testNode{&UnusedAttr{
			Addr:       0x7fe3e01416d0,
			Pos:        NewPositionFromString("col:47"),
			ChildNodes: []Node{},
			IsUnused:   true,
		},
			0x7fe3e01416d0,
			NewPositionFromString("col:47"),
			[]Node{},
		},
		`0x7fe3e01416d0 <col:47>`: testNode{&UnusedAttr{
			Addr:       0x7fe3e01416d0,
			Pos:        NewPositionFromString("col:47"),
			ChildNodes: []Node{},
			IsUnused:   false,
		},
			0x7fe3e01416d0,
			NewPositionFromString("col:47"),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
