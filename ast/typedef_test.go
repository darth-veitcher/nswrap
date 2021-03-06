package ast

import (
	"testing"
)

func TestTypedef(t *testing.T) {
	nodes := map[string]testNode{
		`0x7f84d10dc1d0 '__darwin_ssize_t'`: testNode{&Typedef{
			Addr:       0x7f84d10dc1d0,
			Type:       "__darwin_ssize_t",
			ChildNodes: []Node{},
		},
			0x7f84d10dc1d0,
			NewPositionFromString(""),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
