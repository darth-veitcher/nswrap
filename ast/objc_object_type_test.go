package ast

import (
	"testing"
)

func TestObjCObjectType(t *testing.T) {
	nodes := map[string]testNode{
		`0x10c101ab0 'NSObject<OS_xpc_object>'`: testNode{&ObjCObjectType{
			Addr:       0x10c101ab0,
			Type:       `NSObject<OS_xpc_object>`,
			ChildNodes: []Node{},
		},
			0x10c101ab0,
			NewPositionFromString(""),
			[]Node{},
		},
		`0x7faa18805fc0 'id'`: testNode{&ObjCObjectType{
			Addr:       0x7faa18805fc0,
			Type:       `id`,
			ChildNodes: []Node{},
		},
			0x7faa18805fc0,
			NewPositionFromString(""),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
