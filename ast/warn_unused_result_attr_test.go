package ast

import (
	"testing"
)

func TestWarnUnusedResultAttr(t *testing.T) {
	nodes := map[string]testNode{
		`0x7fa1d704d420 <col:60> warn_unused_result`:
		testNode{&WarnUnusedResultAttr{
			Addr:       0x7fa1d704d420,
			Pos:        NewPositionFromString("col:60"),
			ChildNodes: []Node{},
		},
		0x7fa1d704d420,
		NewPositionFromString("col:60"),
		[]Node{},
		},
		`0x1fac810 <line:481:52>`:
		testNode{&WarnUnusedResultAttr{
			Addr:       0x1fac810,
			Pos:        NewPositionFromString("line:481:52"),
			ChildNodes: []Node{},
		},
		0x1fac810,
		NewPositionFromString("line:481:52"),
		[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
