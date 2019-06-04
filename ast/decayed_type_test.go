package ast

import (
	"testing"
)

func TestDecayedType(t *testing.T) {
	nodes := map[string]testNode{
		`0x7f1234567890 'struct __va_list_tag *' sugar`:
		testNode{&DecayedType{
			Addr:       0x7f1234567890,
			Type:       "struct __va_list_tag *",
			ChildNodes: []Node{},
		},
		0x7f1234567890,
		NewPositionFromString(""),
		[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
