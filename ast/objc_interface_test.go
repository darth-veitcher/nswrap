package ast

import (
	"testing"
)

func TestObjCInterface(t *testing.T) {
	i := 1
	runNodeTest(t,
		Parse(`ObjCInterface 0x7f84d10dc1d0 'NSObject'`),
		testNode{&ObjCInterface{
			Addr:       0x7f84d10dc1d0,
			Name:       `NSObject`,
			Super:      false,
			ChildNodes: []Node{},
		},
		0x7f84d10dc1d0,
		NewPositionFromString(""),
		[]Node{},
		},
		&i,
	)
	runNodeTest(t,
		Parse(`super ObjCInterface 0x7f84d10dc1d0 'NSObject'`),
		testNode{&ObjCInterface{
			Addr:       0x7f84d10dc1d0,
			Name:       `NSObject`,
			Super:      true,
			ChildNodes: []Node{},
		},
		0x7f84d10dc1d0,
		NewPositionFromString(""),
		[]Node{},
		},
		&i,
	)
}
