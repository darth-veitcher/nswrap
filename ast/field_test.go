package ast

import (
	"testing"
)

func TestField(t *testing.T) {
	nodes := map[string]testNode{
		`0x44159a0 '' 'union sigcontext::(anonymous at /usr/include/x86_64-linux-gnu/bits/sigcontext.h:165:17)'`:
		testNode{&Field{
			Addr:       0x44159a0,
			String1:    "",
			String2:    "union sigcontext::(anonymous at /usr/include/x86_64-linux-gnu/bits/sigcontext.h:165:17)",
			ChildNodes: []Node{},
		},
		0x44159a0,
		NewPositionFromString(""),
		[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
