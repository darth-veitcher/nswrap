package ast

import (
	"testing"
)

func TestBlockPointerType(t *testing.T) {
	nodes := map[string]testNode{
		`0x7fa3b88bbb30 'void (^)(void)'`:
		testNode{&BlockPointerType{
			Addr:       0x7fa3b88bbb30,
			Type:       `void (^)(void)`,
			ChildNodes: []Node{},
		},
		0x7fa3b88bbb30,
		NewPositionFromString(""),
		[]Node{},
		},
		`0x7fa3b88bbb30 'NSComparisonResult (^)(id _Nonnull, id _Nonnull)'`:
		testNode{&BlockPointerType{
			Addr:       0x7fa3b88bbb30,
			Type:       `NSComparisonResult (^)(id _Nonnull, id _Nonnull)`,
			ChildNodes: []Node{},
		},
		0x7fa3b88bbb30,
		NewPositionFromString(""),
		[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
