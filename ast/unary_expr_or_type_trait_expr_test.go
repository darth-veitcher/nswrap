package ast

import (
	"testing"
)

func TestUnaryExprOrTypeTraitExpr(t *testing.T) {
	nodes := map[string]testNode{
		`0x7fccd70adf50 <col:29, col:40> 'unsigned long' sizeof 'char'`: testNode{&UnaryExprOrTypeTraitExpr{
			Addr:       0x7fccd70adf50,
			Pos:        NewPositionFromString("col:29, col:40"),
			Type1:      "unsigned long",
			Function:   "sizeof",
			Type2:      "char",
			Type3:      "",
			ChildNodes: []Node{},
		},
			0x7fccd70adf50,
			NewPositionFromString("col:29, col:40"),
			[]Node{},
		},
		`0x7fae1a800190 <col:36, col:44> 'unsigned long' sizeof`: testNode{&UnaryExprOrTypeTraitExpr{
			Addr:       0x7fae1a800190,
			Pos:        NewPositionFromString("col:36, col:44"),
			Type1:      "unsigned long",
			Function:   "sizeof",
			Type2:      "",
			Type3:      "",
			ChildNodes: []Node{},
		},
			0x7fae1a800190,
			NewPositionFromString("col:36, col:44"),
			[]Node{},
		},
		`0x557e575e70b8 <col:432, col:452> 'unsigned long' sizeof 'union MyUnion':'union MyUnion'`: testNode{&UnaryExprOrTypeTraitExpr{
			Addr:       0x557e575e70b8,
			Pos:        NewPositionFromString("col:432, col:452"),
			Type1:      "unsigned long",
			Function:   "sizeof",
			Type2:      "union MyUnion",
			Type3:      "union MyUnion",
			ChildNodes: []Node{},
		},
			0x557e575e70b8,
			NewPositionFromString("col:432, col:452"),
			[]Node{},
		},
		`0x3f142d8 <col:30, col:45> 'unsigned long' sizeof 'extCoord':'extCoord'`: testNode{&UnaryExprOrTypeTraitExpr{
			Addr:       0x3f142d8,
			Pos:        NewPositionFromString("col:30, col:45"),
			Type1:      "unsigned long",
			Function:   "sizeof",
			Type2:      "extCoord",
			Type3:      "extCoord",
			ChildNodes: []Node{},
		},
			0x3f142d8,
			NewPositionFromString("col:30, col:45"),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
