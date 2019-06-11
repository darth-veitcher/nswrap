package ast

import (
	"testing"
)

func TestConditionalOperator(t *testing.T) {
	nodes := map[string]testNode{
		`0x7fc6ae0bc678 <col:6, col:89> 'void'`: testNode{&ConditionalOperator{
			Addr:       0x7fc6ae0bc678,
			Pos:        NewPositionFromString("col:6, col:89"),
			Type:       "void",
			ChildNodes: []Node{},
		},
			0x7fc6ae0bc678,
			NewPositionFromString("col:6, col:89"),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
