package ast

import (
	"testing"
)

func TestNoInlineAttr(t *testing.T) {
	nodes := map[string]testNode{
		`0x7fc02a8a6730 <line:24619:23>`:
		testNode{&NoInlineAttr{
			Addr:       0x7fc02a8a6730,
			Pos:        NewPositionFromString("line:24619:23"),
			ChildNodes: []Node{},
		},
		0x7fc02a8a6730,
		NewPositionFromString("line:24619:23"),
		[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
