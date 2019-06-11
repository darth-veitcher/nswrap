package ast

import (
	"testing"
)

func TestAllocSizeAttr(t *testing.T) {
	nodes := map[string]testNode{
		`0x7f8e390a5d38 <col:100, col:114> 1 2`: testNode{&AllocSizeAttr{
			Addr:       0x7f8e390a5d38,
			Pos:        NewPositionFromString("col:100, col:114"),
			A:          "1",
			B:          "2",
			ChildNodes: []Node{},
		},
			0x7f8e390a5d38,
			NewPositionFromString("col:100, col:114"),
			[]Node{},
		},
		`0x7fbd1a167f48 </usr/include/stdlib.h:342:37> Inherited 1 0`: testNode{&AllocSizeAttr{
			Addr:       0x7fbd1a167f48,
			Pos:        NewPositionFromString("/usr/include/stdlib.h:342:37"),
			Inherited:  true,
			A:          "1",
			B:          "0",
			ChildNodes: []Node{},
		},
			0x7fbd1a167f48,
			NewPositionFromString("/usr/include/stdlib.h:342:37"),
			[]Node{},
		},
		`0x7fbd1a167f48 </usr/include/stdlib.h:342:37> Inherited 1`: testNode{&AllocSizeAttr{
			Addr:       0x7fbd1a167f48,
			Pos:        NewPositionFromString("/usr/include/stdlib.h:342:37"),
			Inherited:  true,
			A:          "1",
			B:          "",
			ChildNodes: []Node{},
		},
			0x7fbd1a167f48,
			NewPositionFromString("/usr/include/stdlib.h:342:37"),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
