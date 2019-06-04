package ast

import (
	"testing"
)

func TestVariadic(t *testing.T) {
	i := 1
	runNodeTest(t,
		Parse(`...`),
		testNode{&Variadic{
			ChildNodes: []Node{},
		},
		0,
		NewPositionFromString(""),
		[]Node{},
		},
		&i,
	)
}
