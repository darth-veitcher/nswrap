package ast

import (
	"testing"
)

func TestImplicitCastExpr(t *testing.T) {
	nodes := map[string]testNode{
		`0x7f9f5b0a1288 <col:8> 'FILE *' <LValueToRValue>`:
		testNode{&ImplicitCastExpr{
			Addr:       0x7f9f5b0a1288,
			Pos:        NewPositionFromString("col:8"),
			Type:       "FILE *",
			Kind:       "LValueToRValue",
			ChildNodes: []Node{},
		},
		0x7f9f5b0a1288,
		NewPositionFromString("col:8"),
		[]Node{},
		},
		`0x7f9f5b0a7828 <col:11> 'int (*)(int, FILE *)' <FunctionToPointerDecay>`:
		testNode{&ImplicitCastExpr{
			Addr:       0x7f9f5b0a7828,
			Pos:        NewPositionFromString("col:11"),
			Type:       "int (*)(int, FILE *)",
			Kind:       "FunctionToPointerDecay",
			ChildNodes: []Node{},
		},
		0x7f9f5b0a7828,
		NewPositionFromString("col:11"),
		[]Node{},
		},
		`0x21267c8 <col:8> 'enum week1':'enum week2' <IntegralCast>`:
		testNode{&ImplicitCastExpr{
			Addr:       0x21267c8,
			Pos:        NewPositionFromString("col:8"),
			Type:       "enum week1",
			Type2:      "enum week2",
			Kind:       "IntegralCast",
			ChildNodes: []Node{},
		},
		0x21267c8,
		NewPositionFromString("col:8"),
		[]Node{},
		},
		`0x26fd2d8 <col:20, col:32> 'extCoord':'extCoord' <LValueToRValue>`:
		testNode{&ImplicitCastExpr{
			Addr:       0x26fd2d8,
			Pos:        NewPositionFromString("col:20, col:32"),
			Type:       "extCoord",
			Type2:      "extCoord",
			Kind:       "LValueToRValue",
			ChildNodes: []Node{},
		},
		0x26fd2d8,
		NewPositionFromString("col:20, col:32"),
		[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
