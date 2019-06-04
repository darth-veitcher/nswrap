package ast

import (
	"testing"
)

func TestBinaryOperator(t *testing.T) {
	nodes := map[string]testNode{
		`0x7fca2d8070e0 <col:11, col:23> 'unsigned char' '='`:
		testNode{&BinaryOperator{
			Addr:       0x7fca2d8070e0,
			Pos:        NewPositionFromString("col:11, col:23"),
			Type:       "unsigned char",
			Operator:   "=",
			ChildNodes: []Node{},
		},
		0x7fca2d8070e0,
		NewPositionFromString("col:11, col:23"),
		[]Node{},
		},
		`0x1ff95b8 <line:78:2, col:7> 'T_ENUM':'T_ENUM' '='`:
		testNode{&BinaryOperator{
			Addr:       0x1ff95b8,
			Pos:        NewPositionFromString("line:78:2, col:7"),
			Type:       "T_ENUM",
			Type2:      "",
			Operator:   "=",
			ChildNodes: []Node{},
		},
		0x1ff95b8,
		NewPositionFromString("line:78:2, col:7"),
		[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
