package ast

import (
	"testing"
)

func TestRestrictAttr(t *testing.T) {
	nodes := map[string]testNode{
		`0x7f980b858305 <line:11:7, line:18:7> foo`: testNode{&RestrictAttr{
			Addr:       0x7f980b858305,
			Pos:        NewPositionFromString("line:11:7, line:18:7"),
			Name:       "foo",
			ChildNodes: []Node{},
		},
			0x7f980b858305,
			NewPositionFromString("line:11:7, line:18:7"),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
