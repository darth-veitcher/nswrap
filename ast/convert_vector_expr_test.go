package ast

import (
	"testing"
)

func TestConvertVectorExpr(t *testing.T) {
	nodes := map[string]testNode{
		`0x7fdef0862430 <line:120:1, col:16>`: testNode{&ConvertVectorExpr{
			Addr:       0x7fdef0862430,
			Pos:        NewPositionFromString("line:120:1, col:16"),
			Type:       "",
			Type2:      "",
			ChildNodes: []Node{},
		},
			0x7fdef0862430,
			NewPositionFromString("line:120:1, col:16"),
			[]Node{},
		},
		`0x113368318 <line:1354:20, line:1355:70> '__v2df':'__attribute__((__vector_size__(2 * sizeof(double)))) double'`: testNode{&ConvertVectorExpr{
			Addr:       0x113368318,
			Pos:        NewPositionFromString("line:1354:20, line:1355:70"),
			Type:       `__v2df`,
			Type2:      `:'__attribute__((__vector_size__(2 * sizeof(double)))) double'`,
			ChildNodes: []Node{},
		},
			0x113368318,
			NewPositionFromString("line:1354:20, line:1355:70"),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
