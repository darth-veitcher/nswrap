package ast

import (
	"testing"
)

func TestShuffleVectorExpr(t *testing.T) {
	nodes := map[string]testNode{
		`0x7fdef0862430 <line:120:1, col:16>`:
		testNode{&ShuffleVectorExpr{
			Addr:         0x7fdef0862430,
			Pos:          NewPositionFromString("line:120:1, col:16"),
			Type:         "",
			Type2:        "",
			ChildNodes:   []Node{},
		},
		0x7fdef0862430,
		NewPositionFromString("line:120:1, col:16"),
		[]Node{},
		},
		`0x113368318 <line:1354:20, line:1355:70> '__v4sf':'__attribute__((__vector_size__(4 * sizeof(float)))) float'`:
		testNode{&ShuffleVectorExpr{
			Addr:         0x113368318,
			Pos:          NewPositionFromString("line:1354:20, line:1355:70"),
			Type:         `__v4sf`,
			Type2:        `:'__attribute__((__vector_size__(4 * sizeof(float)))) float'`,
			ChildNodes:   []Node{},
		},
		0x113368318,
		NewPositionFromString("line:1354:20, line:1355:70"),
		[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
