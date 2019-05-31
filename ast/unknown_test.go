package ast

import (
	"testing"
)

func TestUnknown(t *testing.T) {
	nodes := map[string]Node{
		`0x7faa18a445d8 <line:66:45> "asdf" aoeu`:
		&Unknown{
			Addr:         0x7faa18a445d8,
			Name:         "Unknown",
			Pos:          NewPositionFromString("line:66:45"),
			Position2:    "",
			Content:      ` "asdf" aoeu`,
			ChildNodes:   []Node{},
		},
	}

	runNodeTests(t, nodes)
}
