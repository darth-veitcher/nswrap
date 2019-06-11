package ast

import (
	"testing"
)

func TestPureAttr(t *testing.T) {
	nodes := map[string]testNode{
		`0x7fe9eb899198 <col:1> Implicit`: testNode{&PureAttr{
			Addr:       0x7fe9eb899198,
			Pos:        NewPositionFromString("col:1"),
			Implicit:   true,
			Inherited:  false,
			ChildNodes: []Node{},
		},
			0x7fe9eb899198,
			NewPositionFromString("col:1"),
			[]Node{},
		},
		`0x7fe8d60992a0 <col:1> Inherited Implicit`: testNode{&PureAttr{
			Addr:       0x7fe8d60992a0,
			Pos:        NewPositionFromString("col:1"),
			Implicit:   true,
			Inherited:  true,
			ChildNodes: []Node{},
		},
			0x7fe8d60992a0,
			NewPositionFromString("col:1"),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
