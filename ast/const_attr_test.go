package ast

import (
	"testing"
)

func TestConstAttr(t *testing.T) {
	nodes := map[string]testNode{
		`0x7fa3b88bbb38 <line:4:1, line:13:1>foo`:
		testNode{&ConstAttr{
			Addr:       0x7fa3b88bbb38,
			Pos:        NewPositionFromString("line:4:1, line:13:1"),
			Tags:       "foo",
			ChildNodes: []Node{},
		},
		0x7fa3b88bbb38,
		NewPositionFromString("line:4:1, line:13:1"),
		[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
