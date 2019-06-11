package ast

import (
	"testing"
)

func TestWeakAttr(t *testing.T) {
	nodes := map[string]testNode{
		`0x56069ece5110 <line:736:22>`: testNode{&WeakAttr{
			Addr:       0x56069ece5110,
			Pos:        NewPositionFromString("line:736:22"),
			ChildNodes: []Node{},
		},
			0x56069ece5110,
			NewPositionFromString("line:736:22"),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
