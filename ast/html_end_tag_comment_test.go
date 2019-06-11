package ast

import (
	"testing"
)

func TestHTMLEndTagComment(t *testing.T) {
	nodes := map[string]testNode{
		`0x4259670 <col:27, col:30> Name="i"`: testNode{&HTMLEndTagComment{
			Addr:       0x4259670,
			Pos:        NewPositionFromString("col:27, col:30"),
			Name:       "i",
			ChildNodes: []Node{},
		},
			0x4259670,
			NewPositionFromString("col:27, col:30"),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
