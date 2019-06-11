package ast

import (
	"testing"
)

func TestCallExpr(t *testing.T) {
	nodes := map[string]testNode{
		`0x7f9bf3033240 <col:11, col:25> 'int'`: testNode{&CallExpr{
			Addr:       0x7f9bf3033240,
			Pos:        NewPositionFromString("col:11, col:25"),
			Type:       "int",
			ChildNodes: []Node{},
		},
			0x7f9bf3033240,
			NewPositionFromString("col:11, col:25"),
			[]Node{},
		},
		`0x7f9bf3035c20 <line:7:4, col:64> 'int'`: testNode{&CallExpr{
			Addr:       0x7f9bf3035c20,
			Pos:        NewPositionFromString("line:7:4, col:64"),
			Type:       "int",
			ChildNodes: []Node{},
		},
			0x7f9bf3035c20,
			NewPositionFromString("line:7:4, col:64"),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
