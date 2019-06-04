package ast

import (
	"testing"
)

func TestObjCInterfaceDecl(t *testing.T) {
	nodes := map[string]testNode{
		`0x7fdef0862430 <line:120:1, col:16> col:16 NSObject`:
		testNode{&ObjCInterfaceDecl{
			Addr:         0x7fdef0862430,
			Pos:          NewPositionFromString("line:120:1, col:16"),
			Position2:    "",
			Name:         "NSObject",
			Implicit:     false,
			ChildNodes:   []Node{},
		},
		0x7fdef0862430,
		NewPositionFromString("line:120:1, col:16"),
		[]Node{},
		},
		`0x7fdef0862430 prev 0x7fca43341430 <line:120:1, col:16> col:16 NSObject`:
		testNode{&ObjCInterfaceDecl{
			Addr:         0x7fdef0862430,
			Pos:          NewPositionFromString("line:120:1, col:16"),
			Position2:    "",
			Name:         "NSObject",
			Implicit:     false,
			ChildNodes:   []Node{},
		},
		0x7fdef0862430,
		NewPositionFromString("line:120:1, col:16"),
		[]Node{},
		},
		`0x7fdef0862430 <<invalid sloc>> <invalid sloc> implicit Protocol`:
		testNode{&ObjCInterfaceDecl{
			Addr:         0x7fdef0862430,
			Pos:          Position{},
			Position2:    "",
			Name:         "Protocol",
			Implicit:     true,
			ChildNodes:   []Node{},
		},
		0x7fdef0862430,
		Position{},
		[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
