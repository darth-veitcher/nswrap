package ast

import (
	"testing"
)

func TestParenType(t *testing.T) {
	nodes := map[string]testNode{
		`0x7faf820a4c60 'void (int)' sugar`:
		testNode{&ParenType{
			Addr:       0x7faf820a4c60,
			Type:       "void (int)",
			Sugar:      true,
			ChildNodes: []Node{},
		},
		0x7faf820a4c60,
		NewPositionFromString(""),
		[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
