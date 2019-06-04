package ast

import (
	"testing"
)

func TestAlignedAttr(t *testing.T) {
	nodes := map[string]testNode{
		`0x7f8a1d8ccfd0 <col:47, col:57> aligned`:
		testNode{&AlignedAttr{
			Addr:       0x7f8a1d8ccfd0,
			Pos:        NewPositionFromString("col:47, col:57"),
			IsAligned:  true,
			ChildNodes: []Node{},
			},
			0x7f8a1d8ccfd0,
			NewPositionFromString("col:47, col:57"),
			[]Node{},
		},
		`0x2c8ba10 <col:42>`: testNode{ &AlignedAttr{
			Addr:       0x2c8ba10,
			Pos:        NewPositionFromString("col:42"),
			IsAligned:  false,
			ChildNodes: []Node{},
			},
			0x2c8ba10,
			NewPositionFromString("col:42"),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}

