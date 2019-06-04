package ast

import (
	"testing"
)

func TestQualType(t *testing.T) {
	nodes := map[string]testNode{
		`0x7fa3b88bbb31 'struct _opaque_pthread_t *' foo`:
		testNode{&QualType{
			Addr:       0x7fa3b88bbb31,
			Type:       "struct _opaque_pthread_t *",
			Kind:       "foo",
			ChildNodes: []Node{},
		},
		0x7fa3b88bbb31,
		NewPositionFromString(""),
		[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
