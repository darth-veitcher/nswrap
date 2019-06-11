package ast

import (
	"testing"
)

func TestElaboratedType(t *testing.T) {
	nodes := map[string]testNode{
		`0x7f873686c120 'union __mbstate_t' sugar`: testNode{&ElaboratedType{
			Addr:       0x7f873686c120,
			Type:       "union __mbstate_t",
			Tags:       "sugar",
			ChildNodes: []Node{},
		},
			0x7f873686c120,
			NewPositionFromString(""),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
