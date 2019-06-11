package ast

import (
	"testing"
)

func TestTranslationUnitDecl(t *testing.T) {
	nodes := map[string]testNode{
		`0x7fe78a815ed0 <<invalid sloc>> <invalid sloc>`: testNode{&TranslationUnitDecl{
			Addr:       0x7fe78a815ed0,
			ChildNodes: []Node{},
		},
			0x7fe78a815ed0,
			NewPositionFromString(""),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
