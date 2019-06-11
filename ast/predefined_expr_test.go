package ast

import (
	"testing"
)

func TestPredefinedExpr(t *testing.T) {
	nodes := map[string]testNode{
		`0x33d6e08 <col:30> 'const char [25]' lvalue __PRETTY_FUNCTION__`: testNode{&PredefinedExpr{
			Addr:       0x33d6e08,
			Pos:        NewPositionFromString("col:30"),
			Type:       "const char [25]",
			Lvalue:     true,
			Name:       "__PRETTY_FUNCTION__",
			ChildNodes: []Node{},
		},
			0x33d6e08,
			NewPositionFromString("col:30"),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
