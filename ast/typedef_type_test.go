package ast

import (
	"testing"
)

func TestTypedefType(t *testing.T) {
	nodes := map[string]testNode{
		`0x7f887a0dc760 '__uint16_t' sugar`:
		testNode{&TypedefType{
			Addr:       0x7f887a0dc760,
			Type:       "__uint16_t",
			Tags:       "sugar",
			ChildNodes: []Node{},
		},
		0x7f887a0dc760,
		NewPositionFromString(""),
		[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
