package ast

import (
	"testing"
)

func TestPointerType(t *testing.T) {
	nodes := map[string]testNode{
		`0x7fa3b88bbb30 'struct _opaque_pthread_t *'`: testNode{&PointerType{
			Addr:       0x7fa3b88bbb30,
			Type:       "struct _opaque_pthread_t *",
			ChildNodes: []Node{},
		},
			0x7fa3b88bbb30,
			NewPositionFromString(""),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
