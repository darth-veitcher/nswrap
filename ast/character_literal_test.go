package ast

import (
	"testing"
)

func TestCharacterLiteral(t *testing.T) {
	nodes := map[string]testNode{
		`0x7f980b858308 <col:62> 'int' 10`: testNode{&CharacterLiteral{
			Addr:       0x7f980b858308,
			Pos:        NewPositionFromString("col:62"),
			Type:       "int",
			Value:      10,
			ChildNodes: []Node{},
		},
			0x7f980b858308,
			NewPositionFromString("col:62"),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
