package ast

import (
	"testing"
)

func TestObjCPropertyDecl(t *testing.T) {
	nodes := map[string]testNode{
		`0x10a9be000 <line:157:1, col:46> col:46 blueColor 'NSColor * _Nonnull':'NSColor *' readonly atomic strong class`: testNode{&ObjCPropertyDecl{
			Addr:       0x10a9be000,
			Pos:        NewPositionFromString("line:157:1, col:46"),
			Position2:  "col:46",
			Name:       "blueColor",
			Type:       `NSColor * _Nonnull`,
			Type2:      "",
			Attr:       "readonly atomic strong class",
			ChildNodes: []Node{},
		},
			0x10a9be000,
			NewPositionFromString("line:157:1, col:46"),
			[]Node{},
		},
		`0x7fca44e4a180 <line:50:1, col:61> col:61 undoRegistrationEnabled 'BOOL':'signed char' readonly atomic getter<col:29>`: testNode{&ObjCPropertyDecl{
			Addr:       0x7fca44e4a180,
			Pos:        NewPositionFromString("line:50:1, col:61"),
			Position2:  "col:61",
			Name:       "undoRegistrationEnabled",
			Type:       "BOOL",
			Type2:      "",
			Attr:       "readonly atomic getter<col:29>",
			ChildNodes: []Node{},
		},
			0x7fca44e4a180,
			NewPositionFromString("line:50:1, col:61"),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
