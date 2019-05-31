package ast

import (
	"testing"
)

func TestObjCInterfaceDecl(t *testing.T) {
	nodes := map[string]Node{
		`0x7fdef0862430 <line:120:1, col:16> col:16 NSObject`:
		&ObjCInterfaceDecl{
			Addr:         0x7fdef0862430,
			Pos:          NewPositionFromString("line:120:1, col:16"),
			Position2:    "",
			Name:         "NSObject",
			Implicit:     false,
			ChildNodes:   []Node{},
		},
		`0x7fdef0862430 prev 0x7fca43341430 <line:120:1, col:16> col:16 NSObject`:
		&ObjCInterfaceDecl{
			Addr:         0x7fdef0862430,
			Pos:          NewPositionFromString("line:120:1, col:16"),
			Position2:    "",
			Name:         "NSObject",
			Implicit:     false,
			ChildNodes:   []Node{},
		},
		`0x7fdef0862430 <<invalid sloc>> <invalid sloc> implicit Protocol`:
		&ObjCInterfaceDecl{
			Addr:         0x7fdef0862430,
			Pos:          Position{},
			Position2:    "",
			Name:         "Protocol",
			Implicit:     true,
			ChildNodes:   []Node{},
		},
	}

	runNodeTests(t, nodes)
}
