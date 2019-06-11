package ast

import (
	"testing"
)

func TestTransparentUnionAttr(t *testing.T) {
	nodes := map[string]testNode{
		`0x304f700 <col:35>`: testNode{&TransparentUnionAttr{
			Addr:       0x304f700,
			Pos:        NewPositionFromString("col:35"),
			ChildNodes: []Node{},
		},
			0x304f700,
			NewPositionFromString("col:35"),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
