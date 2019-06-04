package ast

import (
	"testing"
)

func TestEnumDecl(t *testing.T) {
	nodes := map[string]testNode{
		`0x22a6c80 <line:180:1, line:186:1> __codecvt_result`:
		testNode{&EnumDecl{
			Addr:       0x22a6c80,
			Pos:        NewPositionFromString("line:180:1, line:186:1"),
			Position2:  "",
			Name:       "__codecvt_result",
			ChildNodes: []Node{},
		},
		0x22a6c80,
		NewPositionFromString("line:180:1, line:186:1"),
		[]Node{},
		},
		`0x32fb5a0 <enum.c:3:1, col:45> col:6 week`:
		testNode{&EnumDecl{
			Addr:       0x32fb5a0,
			Pos:        NewPositionFromString("enum.c:3:1, col:45"),
			Position2:  " col:6",
			Name:       "week",
			ChildNodes: []Node{},
		},
		0x32fb5a0,
		NewPositionFromString("enum.c:3:1, col:45"),
		[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
