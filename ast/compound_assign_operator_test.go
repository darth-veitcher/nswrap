package ast

import (
	"testing"
)

func TestCompoundAssignOperator(t *testing.T) {
	nodes := map[string]testNode{
		`0x2dc5758 <line:5:2, col:7> 'int' '+=' ComputeLHSTy='int' ComputeResultTy='int'`:
		testNode{&CompoundAssignOperator{
			Addr:                  0x2dc5758,
			Pos:                   NewPositionFromString("line:5:2, col:7"),
			Type:                  "int",
			Opcode:                "+=",
			ComputationLHSType:    "int",
			ComputationResultType: "int",
			ChildNodes:            []Node{},
		},
		0x2dc5758,
		NewPositionFromString("line:5:2, col:7"),
		[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
