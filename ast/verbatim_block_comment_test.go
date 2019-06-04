package ast

import (
	"testing"
)

func TestVerbatimBlockComment(t *testing.T) {
	nodes := map[string]testNode{
		`0x107781dd0 <col:34, col:39> Name="link" CloseName=""`:
		testNode{&VerbatimBlockComment{
			Addr:       0x107781dd0,
			Pos:        NewPositionFromString("col:34, col:39"),
			Name:       "link",
			CloseName:  "",
			ChildNodes: []Node{},
		},
		0x107781dd0,
		NewPositionFromString("col:34, col:39"),
		[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
