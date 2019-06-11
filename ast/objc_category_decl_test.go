package ast

import (
	"testing"
)

func TestObjCCategoryDecl(t *testing.T) {
	nodes := map[string]testNode{
		`0x7fdef0862430 <line:120:1, col:16> col:16 NSObject`: testNode{&ObjCCategoryDecl{
			Addr:       0x7fdef0862430,
			Pos:        NewPositionFromString("line:120:1, col:16"),
			Position2:  "",
			Name:       "NSObject",
			ChildNodes: []Node{},
		},
			0x7fdef0862430,
			NewPositionFromString("line:120:1, col:16"),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
