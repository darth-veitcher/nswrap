package ast

import (
	"testing"
)

func TestObjCInterface(t *testing.T) {
	runNodeTest(t,
		Parse(`ObjCInterface 0x7f84d10dc1d0 'NSObject'`),
		&ObjCInterface{
			Addr:       0x7f84d10dc1d0,
			Name:       `NSObject`,
			Super:      false,
			ChildNodes: []Node{},
		},
		1,
	)
	runNodeTest(t,
		Parse(`super ObjCInterface 0x7f84d10dc1d0 'NSObject'`),
		&ObjCInterface{
			Addr:       0x7f84d10dc1d0,
			Name:       `NSObject`,
			Super:      true,
			ChildNodes: []Node{},
		},
		2,
	)
}
