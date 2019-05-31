package ast

import (
	"testing"
)

func TestBlockPointerType(t *testing.T) {
	nodes := map[string]Node{
		`0x7fa3b88bbb30 'void (^)(void)'`:
		&BlockPointerType{
			Addr:       0x7fa3b88bbb30,
			Type:       `void (^)(void)`,
			ChildNodes: []Node{},
		},
		`0x7fa3b88bbb30 'NSComparisonResult (^)(id _Nonnull, id _Nonnull)'`:
		&BlockPointerType{
			Addr:       0x7fa3b88bbb30,
			Type:       `NSComparisonResult (^)(id _Nonnull, id _Nonnull)`,
			ChildNodes: []Node{},
		},
	}

	runNodeTests(t, nodes)
}
