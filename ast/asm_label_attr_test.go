package ast

import (
	"testing"
)

func TestAsmLabelAttr(t *testing.T) {
	nodes := map[string]testNode{
		`0x7ff26d8224e8 </usr/include/sys/cdefs.h:569:36> "_fopen"`:
		testNode{&AsmLabelAttr{
			Addr:         0x7ff26d8224e8,
			Pos:          NewPositionFromString("/usr/include/sys/cdefs.h:569:36"),
			Inherited:    false,
			FunctionName: "_fopen",
			ChildNodes:   []Node{},
		},
		0x7ff26d8224e8,
		NewPositionFromString("/usr/include/sys/cdefs.h:569:36"),
		[]Node{},
		},
		`0x7fd55a169318 </usr/include/stdio.h:325:47> Inherited "_popen"`:
		testNode{&AsmLabelAttr{
			Addr:         0x7fd55a169318,
			Pos:          NewPositionFromString("/usr/include/stdio.h:325:47"),
			Inherited:    true,
			FunctionName: "_popen",
			ChildNodes:   []Node{},
		},
		0x7fd55a169318,
		NewPositionFromString("/usr/include/stdio.h:325:47"),
		[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
