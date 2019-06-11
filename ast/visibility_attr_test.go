package ast

import (
	"testing"
)

func TestVisibilityAttr(t *testing.T) {
	nodes := map[string]testNode{
		`0x55c49d8dd1d8 <col:16, col:36> Default`: testNode{&VisibilityAttr{
			Addr:        0x55c49d8dd1d8,
			Pos:         NewPositionFromString("col:16, col:36"),
			ChildNodes:  []Node{},
			IsInherited: false,
			IsDefault:   true,
		},
			0x55c49d8dd1d8,
			NewPositionFromString("col:16, col:36"),
			[]Node{},
		},
		`0x7f8e7b00bb80 </cmark/src/cmark.h:497:16, col:36> Inherited Default`: testNode{&VisibilityAttr{
			Addr:        0x7f8e7b00bb80,
			Pos:         NewPositionFromString("/cmark/src/cmark.h:497:16, col:36"),
			ChildNodes:  []Node{},
			IsInherited: true,
			IsDefault:   true,
		},
			0x7f8e7b00bb80,
			NewPositionFromString("/cmark/src/cmark.h:497:16, col:36"),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
