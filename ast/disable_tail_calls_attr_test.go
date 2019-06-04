package ast

import (
	"testing"
)

func TestDisableTailCallsAttr(t *testing.T) {
	nodes := map[string]testNode{
		`0x7fc8fa094558 <col:107> `:
		testNode{&DisableTailCallsAttr{
			Addr:       0x7fc8fa094558,
			Pos:        NewPositionFromString("col:107"),
			ChildNodes: []Node{},
		},
		0x7fc8fa094558,
		NewPositionFromString("col:107"),
		[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
