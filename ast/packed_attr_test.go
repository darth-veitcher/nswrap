package ast

import (
	"testing"
)

func TestPackedAttr(t *testing.T) {
	nodes := map[string]testNode{
		`0x7fae33b1ed40 <line:551:18>`: testNode{&PackedAttr{
			Addr:       0x7fae33b1ed40,
			Pos:        NewPositionFromString("line:551:18"),
			ChildNodes: []Node{},
		},
			0x7fae33b1ed40,
			NewPositionFromString("line:551:18"),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
