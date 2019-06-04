package ast

import (
	"testing"
)

func TestObjCIvarDecl(t *testing.T) {
	nodes := map[string]testNode{
		`0x7fcd103e8270 </usr/include/objc/NSObject.h:56:5, col:11> col:11 isa 'Class':'Class' protected`:
		testNode{&ObjCIvarDecl{
			Addr:         0x7fcd103e8270,
			Pos:          NewPositionFromString("/usr/include/objc/NSObject.h:56:5, col:11"),
			Position2:    "col:11",
			Name:         "isa",
			Type:         "Class",
			Type2:        "",
			Attr:         "protected",
			ChildNodes:   []Node{},
		},
		0x7fcd103e8270,
		NewPositionFromString("/usr/include/objc/NSObject.h:56:5, col:11"),
		[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
