package ast

import (
	"testing"
)

func TestBlockCommandComment(t *testing.T) {
	nodes := map[string]testNode{
		`0x1069fae60 <col:4, line:163:57> Name="abstract"`:
		testNode{&BlockCommandComment{
			Addr:       0x1069fae60,
			Pos:        NewPositionFromString("col:4, line:163:57"),
			Name:       "abstract",
			ChildNodes: []Node{},
		},
		0x1069fae60,
		NewPositionFromString("col:4, line:163:57"),
		[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
