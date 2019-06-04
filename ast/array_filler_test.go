package ast

import (
	"testing"
)

func TestArrayFiller(t *testing.T) {
	i := 0
	runNodeTest(t, Parse(`array filler`),
		testNode{ &ArrayFiller{ ChildNodes: []Node{} },
		0,NewPositionFromString(""),[]Node{}},
		&i)
}
