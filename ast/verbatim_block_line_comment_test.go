package ast

import (
	"testing"
)

func TestVerbatimBlockLineComment(t *testing.T) {
	nodes := map[string]testNode{
		`0x10f8e8e20 <col:39, col:54> Text=" OSAtomicAdd32}"`:
		testNode{&VerbatimBlockLineComment{
			Addr:       0x10f8e8e20,
			Pos:        NewPositionFromString("col:39, col:54"),
			Text:       " OSAtomicAdd32}",
			ChildNodes: []Node{},
		},
		0x10f8e8e20,
		NewPositionFromString("col:39, col:54"),
		[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
