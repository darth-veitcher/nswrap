package ast

import (
	"testing"
)

func TestLabelStmt(t *testing.T) {
	nodes := map[string]testNode{
		`0x7fe3ba82edb8 <line:18906:1, line:18907:22> 'end_getDigits'`: testNode{&LabelStmt{
			Addr:       0x7fe3ba82edb8,
			Pos:        NewPositionFromString("line:18906:1, line:18907:22"),
			Name:       "end_getDigits",
			ChildNodes: []Node{},
		},
			0x7fe3ba82edb8,
			NewPositionFromString("line:18906:1, line:18907:22"),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
