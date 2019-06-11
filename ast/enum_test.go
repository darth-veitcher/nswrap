package ast

import (
	"testing"
)

func TestEnum(t *testing.T) {
	nodes := map[string]testNode{
		`0x7f980b858308 'foo'`: testNode{&Enum{
			Addr:       0x7f980b858308,
			Name:       "foo",
			ChildNodes: []Node{},
		},
			0x7f980b858308,
			NewPositionFromString(""),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
