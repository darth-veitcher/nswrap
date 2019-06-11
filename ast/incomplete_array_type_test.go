package ast

import (
	"testing"
)

func TestIncompleteArrayType(t *testing.T) {
	nodes := map[string]testNode{
		`0x7fcb7d005c20 'int []' `: testNode{&IncompleteArrayType{
			Addr:       0x7fcb7d005c20,
			Type:       "int []",
			ChildNodes: []Node{},
		},
			0x7fcb7d005c20,
			NewPositionFromString(""),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
