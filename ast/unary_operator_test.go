package ast

import (
	"testing"
)

func TestUnaryOperator(t *testing.T) {
	nodes := map[string]testNode{
		`0x7fe0260f50d8 <col:6, col:12> 'int' prefix '--'`: testNode{&UnaryOperator{
			Addr:       0x7fe0260f50d8,
			Pos:        NewPositionFromString("col:6, col:12"),
			Type:       "int",
			Type2:      "",
			IsLvalue:   false,
			IsPrefix:   true,
			Operator:   "--",
			ChildNodes: []Node{},
		},
			0x7fe0260f50d8,
			NewPositionFromString("col:6, col:12"),
			[]Node{},
		},
		`0x7fe0260fb468 <col:11, col:18> 'unsigned char' lvalue prefix '*'`: testNode{&UnaryOperator{
			Addr:       0x7fe0260fb468,
			Pos:        NewPositionFromString("col:11, col:18"),
			Type:       "unsigned char",
			Type2:      "",
			IsLvalue:   true,
			IsPrefix:   true,
			Operator:   "*",
			ChildNodes: []Node{},
		},
			0x7fe0260fb468,
			NewPositionFromString("col:11, col:18"),
			[]Node{},
		},
		`0x7fe0260fb448 <col:12, col:18> 'unsigned char *' postfix '++'`: testNode{&UnaryOperator{
			Addr:       0x7fe0260fb448,
			Pos:        NewPositionFromString("col:12, col:18"),
			Type:       "unsigned char *",
			Type2:      "",
			IsLvalue:   false,
			IsPrefix:   false,
			Operator:   "++",
			ChildNodes: []Node{},
		},
			0x7fe0260fb448,
			NewPositionFromString("col:12, col:18"),
			[]Node{},
		},
		`0x26fd2b8 <col:20, col:32> 'extCoord':'extCoord' lvalue prefix '*'`: testNode{&UnaryOperator{
			Addr:       0x26fd2b8,
			Pos:        NewPositionFromString("col:20, col:32"),
			Type:       "extCoord",
			Type2:      "extCoord",
			IsLvalue:   true,
			IsPrefix:   true,
			Operator:   "*",
			ChildNodes: []Node{},
		},
			0x26fd2b8,
			NewPositionFromString("col:20, col:32"),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
