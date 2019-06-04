package ast

import (
	"testing"
)

func TestObjCProtocol(t *testing.T) {
	nodes := map[string]testNode{
		`0x10c26d630 'NSColorPickingDefault'`:
		testNode{&ObjCProtocol{
			Addr:         0x10c26d630,
			Name:         "NSColorPickingDefault",
			ChildNodes:   []Node{},
		},
		0x10c26d630,
		NewPositionFromString(""),
		[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
