package ast

import (
	"testing"
)

func TestVerbatimLineComment(t *testing.T) {
	nodes := map[string]testNode{
		`0x108af4dd0 <col:4, col:28> Text=" qos_class_self"`:
		testNode{&VerbatimLineComment{
			Addr:       0x108af4dd0,
			Pos:        NewPositionFromString("col:4, col:28"),
			Text:       " qos_class_self",
			ChildNodes: []Node{},
		},
		0x108af4dd0,
		NewPositionFromString("col:4, col:28"),
		[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
