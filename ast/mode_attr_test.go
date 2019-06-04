package ast

import (
	"testing"
)

func Test(t *testing.T) {
	nodes := map[string]testNode{
		`0x7f980b858309 <line:11:7, line:18:7> foo`:
		testNode{&ModeAttr{
			Addr:       0x7f980b858309,
			Pos:        NewPositionFromString("line:11:7, line:18:7"),
			Name:       "foo",
			ChildNodes: []Node{},
		},
		0x7f980b858309,
		NewPositionFromString("line:11:7, line:18:7"),
		[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
