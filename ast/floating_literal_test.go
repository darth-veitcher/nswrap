package ast

import (
	"testing"
)

func TestFloatingLiteral(t *testing.T) {
	nodes := map[string]testNode{
		`0x7febe106f5e8 <col:24> 'double' 1.230000e+00`: testNode{&FloatingLiteral{
			Addr:       0x7febe106f5e8,
			Pos:        NewPositionFromString("col:24"),
			Type:       "double",
			Value:      1.23,
			ChildNodes: []Node{},
		},
			0x7febe106f5e8,
			NewPositionFromString("col:24"),
			[]Node{},
		},
		`0x21c65b8 <col:41> 'double' 2.718282e+00`: testNode{&FloatingLiteral{
			Addr:       0x21c65b8,
			Pos:        NewPositionFromString("col:41"),
			Type:       "double",
			Value:      2.718282e+00,
			ChildNodes: []Node{},
		},
			0x21c65b8,
			NewPositionFromString("col:41"),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
