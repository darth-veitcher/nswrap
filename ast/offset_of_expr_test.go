package ast

import (
	"testing"
)

func TestOffsetOfExpr(t *testing.T) {
	nodes := map[string]testNode{
		`0x7fa855aab838 <col:63, col:95> 'unsigned long'`:
		testNode{&OffsetOfExpr{
			Addr:       0x7fa855aab838,
			Pos:        NewPositionFromString("col:63, col:95"),
			Type:       "unsigned long",
			ChildNodes: []Node{},
		},
		0x7fa855aab838,
		NewPositionFromString("col:63, col:95"),
		[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
