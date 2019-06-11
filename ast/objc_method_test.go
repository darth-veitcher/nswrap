package ast

import (
	"testing"
)

func TestObjCMethod(t *testing.T) {
	i := 1
	runNodeTest(t,
		Parse(`ObjCMethod 0x7f84d10dc1d0 'isValid'`),
		testNode{&ObjCMethod{
			Addr:       0x7f84d10dc1d0,
			Name:       `isValid`,
			ChildNodes: []Node{},
		},
			0x7f84d10dc1d0,
			NewPositionFromString(""),
			[]Node{},
		},
		&i,
	)
	runNodeTest(t,
		Parse(`getter ObjCMethod 0x7f84d10dc1d0 'isValid'`),
		testNode{&ObjCMethod{
			Addr:       0x7f84d10dc1d0,
			Name:       `isValid`,
			ChildNodes: []Node{},
		},
			0x7f84d10dc1d0,
			NewPositionFromString(""),
			[]Node{},
		},
		&i,
	)
}
