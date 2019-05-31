package ast

import (
	"testing"
)

func TestObjCMethod(t *testing.T) {
	runNodeTest(t,
		Parse(`ObjCMethod 0x7f84d10dc1d0 'isValid'`),
		&ObjCMethod{
			Addr:       0x7f84d10dc1d0,
			Name:       `isValid`,
			ChildNodes: []Node{},
		},
		1,
	)
	runNodeTest(t,
		Parse(`getter ObjCMethod 0x7f84d10dc1d0 'isValid'`),
		&ObjCMethod{
			Addr:       0x7f84d10dc1d0,
			Name:       `isValid`,
			ChildNodes: []Node{},
		},
		2,
	)
}
