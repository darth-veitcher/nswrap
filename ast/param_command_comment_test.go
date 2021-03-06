package ast

import (
	"testing"
)

func TestParamCommandComment(t *testing.T) {
	nodes := map[string]testNode{
		`0x104bca8d0 <col:4, line:59:45> [in] implicitly Param="__attr" ParamIndex=0`: testNode{&ParamCommandComment{
			Addr:       0x104bca8d0,
			Pos:        NewPositionFromString("col:4, line:59:45"),
			Other:      "[in] implicitly Param=\"__attr\" ParamIndex=0",
			ChildNodes: []Node{},
		},
			0x104bca8d0,
			NewPositionFromString("col:4, line:59:45"),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
