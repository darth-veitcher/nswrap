package ast

import (
	"testing"
)

func TestVariadic(t *testing.T) {
	runNodeTest(t,
		Parse(`...`),
		&Variadic{
			ChildNodes: []Node{},
		},
		1,
	)
}
