package ast

import (
	"testing"
)

func TestUnavailableAttr(t *testing.T) {
	nodes := map[string]Node{
		`0x7faa18a445d8 <line:66:45> "" IR_None`:
		&UnavailableAttr{
			Addr:         0x7faa18a445d8,
			Pos:          NewPositionFromString("line:66:45"),
			Position2:    "",
			Content:      `"" IR_None`,
			ChildNodes:   []Node{},
		},
		`0x7faa18a289f8 <line:150:54, col:70> "use a (__bridge id) cast instead" IR_None`:
		&UnavailableAttr{
			Addr:         0x7faa18a289f8,
			Pos:          NewPositionFromString("line:150:54, col:70"),
			Position2:    "",
			Content:      `"use a (__bridge id) cast instead" IR_None`,
			ChildNodes:   []Node{},
		},
	}

	runNodeTests(t, nodes)
}
