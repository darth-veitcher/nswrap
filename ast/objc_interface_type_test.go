package ast

import (
	"testing"
)

func TestObjCInterfaceType(t *testing.T) {
	nodes := map[string]testNode{
		`0x7fdef0862430 'NSObject'`: testNode{&ObjCInterfaceType{
			Addr:       0x7fdef0862430,
			Type:       "NSObject",
			ChildNodes: []Node{},
		},
			0x7fdef0862430,
			NewPositionFromString(""),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
