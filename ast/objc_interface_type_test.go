package ast

import (
	"testing"
)

func TestObjCInterfaceType(t *testing.T) {
	nodes := map[string]Node{
		`0x7fdef0862430 'NSObject'`:
		&ObjCInterfaceType{
			Addr:         0x7fdef0862430,
			Type:         "NSObject",
			ChildNodes:   []Node{},
		},
	}

	runNodeTests(t, nodes)
}
