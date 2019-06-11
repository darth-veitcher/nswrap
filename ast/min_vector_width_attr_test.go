package ast

import (
	"testing"
)

func TestMinVectorWidthAttr(t *testing.T) {
	nodes := map[string]testNode{
		`0x7fc0a69091d1 <line:11:7, line:18:7> content`: testNode{&MinVectorWidthAttr{
			Addr:       0x7fc0a69091d1,
			Pos:        NewPositionFromString("line:11:7, line:18:7"),
			Content:    " content",
			ChildNodes: []Node{},
		},
			0x7fc0a69091d1,
			NewPositionFromString("line:11:7, line:18:7"),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
