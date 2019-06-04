package ast

import (
	"testing"
)

func TestTextComment(t *testing.T) {
	nodes := map[string]testNode{
		`0x3085bc0 <line:9950:2, col:29> Text="* CUSTOM AUXILIARY FUNCTIONS"`:
		testNode{&TextComment{
			Addr:       0x3085bc0,
			Pos:        NewPositionFromString("line:9950:2, col:29"),
			Text:       "* CUSTOM AUXILIARY FUNCTIONS",
			ChildNodes: []Node{},
		},
		0x3085bc0,
		NewPositionFromString("line:9950:2, col:29"),
		[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
