package ast

import (
	"testing"
)

func TestRecordType(t *testing.T) {
	nodes := map[string]testNode{
		`0x7fd3ab84dda0 'struct _opaque_pthread_condattr_t'`:
		testNode{&RecordType{
			Addr:       0x7fd3ab84dda0,
			Type:       "struct _opaque_pthread_condattr_t",
			ChildNodes: []Node{},
		},
		0x7fd3ab84dda0,
		NewPositionFromString(""),
		[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
