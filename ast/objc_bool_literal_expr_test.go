package ast

import (
	"testing"
)

func TestObjCBoolLiteralExpr(t *testing.T) {
	nodes := map[string]testNode{
		`0x7fcd0f9e7fe8 <col:13> 'BOOL':'signed char' __objc_yes`: testNode{&ObjCBoolLiteralExpr{
			Addr:       0x7fcd0f9e7fe8,
			Pos:        NewPositionFromString("col:13"),
			Type:       "BOOL",
			Type2:      ":'signed char'",
			Attr:       " __objc_yes",
			ChildNodes: []Node{},
		},
			0x7fcd0f9e7fe8,
			NewPositionFromString("col:13"),
			[]Node{},
		},
		`0x7fcd0f9ed000 <col:13> 'BOOL':'signed char' __objc_no`: testNode{&ObjCBoolLiteralExpr{
			Addr:       0x7fcd0f9ed000,
			Pos:        NewPositionFromString("col:13"),
			Type:       "BOOL",
			Type2:      ":'signed char'",
			Attr:       " __objc_no",
			ChildNodes: []Node{},
		},
			0x7fcd0f9ed000,
			NewPositionFromString("col:13"),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
