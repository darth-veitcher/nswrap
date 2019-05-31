package ast

import (
	"testing"
)

func TestObjCProtocol(t *testing.T) {
	nodes := map[string]Node{
		`0x10c26d630 'NSColorPickingDefault'`:
		&ObjCProtocol{
			Addr:         0x10c26d630,
			Name:         "NSColorPickingDefault",
			ChildNodes:   []Node{},
		},
	}

	runNodeTests(t, nodes)
}
