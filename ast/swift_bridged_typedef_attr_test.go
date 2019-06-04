package ast

import (
	"testing"
)

func TestSwiftBridgedTypedefAttr(t *testing.T) {
	nodes := map[string]testNode{
		`0x7fc0a69091d1 <line:11:7, line:18:7>`:
		testNode{&SwiftBridgedTypedefAttr{
			Addr:       0x7fc0a69091d1,
			Pos:        NewPositionFromString("line:11:7, line:18:7"),
			ChildNodes: []Node{},
		},
		0x7fc0a69091d1,
		NewPositionFromString("line:11:7, line:18:7"),
		[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
