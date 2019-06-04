package ast

import (
	"testing"
)

func TestEnumType(t *testing.T) {
	nodes := map[string]testNode{
		`0x7f980b858309 'foo'`:
		testNode{&EnumType{
			Addr:       0x7f980b858309,
			Name:       "foo",
			ChildNodes: []Node{},
		},
		0x7f980b858309,
		NewPositionFromString(""),
		[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
