package ast

import (
	"testing"
)

func TestFullComment(t *testing.T) {
	nodes := map[string]testNode{
		`0x3860920 <line:10176:4, line:10180:45>`:
		testNode{&FullComment{
			Addr:       0x3860920,
			Pos:        NewPositionFromString("line:10176:4, line:10180:45"),
			ChildNodes: []Node{},
		},
		0x3860920,
		NewPositionFromString("line:10176:4, line:10180:45"),
		[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
