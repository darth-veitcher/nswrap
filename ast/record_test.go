package ast

import (
	"testing"
)

func TestRecord(t *testing.T) {
	nodes := map[string]testNode{
		`0x7fd3ab857950 '__sFILE'`: testNode{&Record{
			Addr:       0x7fd3ab857950,
			Type:       "__sFILE",
			ChildNodes: []Node{},
		},
			0x7fd3ab857950,
			NewPositionFromString(""),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
