package ast

import (
	"testing"
)

func TestObjCObjectPointerType(t *testing.T) {
	nodes := map[string]testNode{
		`0x10c101ab0 'NSObject<OS_xpc_object> *'`: testNode{&ObjCObjectPointerType{
			Addr:       0x10c101ab0,
			Type:       `NSObject<OS_xpc_object> *`,
			ChildNodes: []Node{},
		},
			0x10c101ab0,
			NewPositionFromString(""),
			[]Node{},
		},
		`0x7faa18805fc0 'id'`: testNode{&ObjCObjectPointerType{
			Addr:       0x7faa18805fc0,
			Type:       `id`,
			ChildNodes: []Node{},
		},
			0x7faa18805fc0,
			NewPositionFromString(""),
			[]Node{},
		},
		`0x7fca45a08a60 'NSAppleEventDescriptor *'`: testNode{&ObjCObjectPointerType{
			Addr:       0x7fca45a08a60,
			Type:       `NSAppleEventDescriptor *`,
			ChildNodes: []Node{},
		},
			0x7fca45a08a60,
			NewPositionFromString(""),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
